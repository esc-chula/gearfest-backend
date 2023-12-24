# Gearfest backend

Backend interface for the GearFestival website.

## Prerequisites

- [Go](https://go.dev) 1.21.5 or above
- [Docker](https://docs.docker.com/get-docker/)
- [Supabase CLI](https://github.com/supabase/cli)

## Installation

1. Clone this repo.
2. Copy `config.local.yaml` in `config` and paste it in the same directory with `.local` removed from its name.
3. Run `go mod download` to download all the dependencies.

## Running

1. Run `supabase start` to start supabase.

   > Note: You have to install [Docker](https://docs.docker.com/get-docker/) and [Supabase CLI](https://github.com/supabase/cli) before running this command.

2. Run `go run ./src/` to start server.
3. Server should be running on `localhost:8080` and `localhost:54323/project/default/editor` for database editor.

Ensure to run `supabase stop` to close supabase after finishing your code.

## Contributing

1. Create a new branch

   ```bash
   git checkout -b <branch-name> origin/dev
   ```

1. Make your changes
1. Stage and commit your changes

   ```bash
   git add .

   git commit -m "<commit-message>"
   ```

   > Note: Don't forget to use the [conventional commit format](#conventional-commit-format) for your commit message.

1. Push your changes

   ```bash
   git push origin <branch-name>
   ```

1. Create a pull request to the dev branch in GitHub
1. Wait for the code to be reviewed and merged
1. Repeat

   > Note: Don't forget to always pull the latest changes from the dev branch before creating a new branch.
   >
   > ```bash
   > git pull origin dev
   > ```

### Conventional Commit Format

In short, the commit message should look like this:

```bash
git commit -m "feat: <what-you-did>"

# or

git commit -m "fix: <what-you-fixed>"

# or

git commit -m "refactor: <what-you-refactored>"
```

The commit message should start with one of the following types:

- feat: A new feature
- fix: A bug fix
- refactor: A code change that neither fixes a bug nor adds a feature
- style: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
- docs: Documentation only changes
- chore: Changes to the build process or auxiliary tools and libraries

For more information, please read the [conventional commit format](https://www.conventionalcommits.org/en/v1.0.0/) documentation.
