# Summary

## Goal

Opinionated project: `sqlite` > `postgres` > `anything else`.
Every project has packages and every package has its own sqlite database.


Project structure: 
```bash
<project_dir>
├─ main.go
├─ pkg/
│  ├─ pkg_1/
│  ├─ pkg_2/
│  ├─ ...
│  ├─ pkg_n/
```

Package structure:
```bash
pkg
├─ db/
│  ├─ 0.0.0
│  ├─ 0.0.1
│  ├─ ...
│  ├─ x.y.z
├─ domain
│  ├─ entity_1
│  ├─ entity_2
│  ├─ ...
│  ├─ entity_n
├─ repository
│  ├─ repository.go
│  ├─ entity_i
│  │  ├─ repository.go
│  │  ├─ crud.go
├─ service
│  ├─ service.go
│  ├─ entity_i
│  │  ├─ service.go
│  │  ├─ use_case_1.go
│  │  ├─ use_case_2.go
│  │  ├─ ...
│  │  ├─ use_case_n.go
├─ setup
```

Example for user management
```sql
-- pkg/iam/db/0.0.0/001_create_users_table.sql
create table users (
   id           uuid    primary key not null,
   emailaddress varchar unique      default '',
   name         varchar             default '',
   enabled      boolean             default true,
   created_at   integer             default (strftime('%s', 'now'))
);
```

```golang
// pkg/iam/domain/user.go
package iam_domain

type User struct {
   Id        *uuid.UUID // since this can be optional when creating a user
   Name      string
   Email     string
   Password  string
   Enabled   bool
   CreatedAt time.Time
}
```


```golang
// pkg/iam/repository/repository.go

type IamRepository struct {
   UserRepository iam_user_repository.IamUserRepositoryInterface
   // repositories for other entities here
}

func NewIamRepository(db DatabaseInterface) *IamRepository {
   return &IamRepository{
      UserRepository: iam_user_repository.NewIamUserRepository(db),
      // repositories for other entities here
   }
}
```

```golang
// pkg/iam/repository/user/repository.go
package iam_user_repository

type IamUserRepositoryInterface interface {
   Create(user *iam_domain.User) (*uuid.UUID, error)
   Read(id *uuid.UUID) (*iam_domain.User, error)
   Update(user *iam_domain.User) error
   Delete(id *uuid.UUID) error
}

type IamUserRepository struct {
   db DatabaseInterface
}

func NewIamUserRepository(db DatabaseInterface) IamUserRepositoryInterface {
   return &IamUserRepository{
      db: db,
   }
}
```

Example: C of CRUD
```golang
// pkg/iam/repository/user/create.go
package iam_user_repository

func (repo *IamUserRepository) Create(user *iam_domain.User) (*uuid.UUID, error) {
   id, _ := uuid.NewV7()
   result, err := db.Exec("insert into users (id, name, emailaddress, enabled) values (?, ?, ?, ?)", id.String(), user.Name, user.Email, false)
   if err != nil {
      return nil, err
   }

   return &id, nil
}
```

## Development

The library itself is split into five packages:

- `version/` - the `Version` type (major.minor.patch) and comparisons (`EarlierThan`, `LaterThan`, `Equals`). Solves: knowing whether a package needs migrating and how far.
- `migrator/` - `From(...).To(...).WithLocalFilesAt(...)` / `WithPublicRepository(...)`, reads a package's versioned `.sql` files and applies the ones between the installed and target version. Solves: turning a set of versioned SQL files, local or downloaded, into applied migrations and open database connections.
- `registry/` - `For`/`Reset`, caches a built `PackageBuilder` per name+dataDir+target so repeated calls don't re-migrate or reopen a database. Solves: giving a process one long-lived, migrated instance per package, and a way to close and rebuild it (used by tests, and by live db resets).
- `db/` - a small `DatabaseInterface` wrapping sqlite. Solves: a swappable, mockable database handle.
- `internal/` - private helper for downloading and extracting a repo zip, used only by `migrator`. Not part of the public API.
