# Gearfest backend

Backend interface for the GearFestival website.

## Prerequisites

- [Go](https://go.dev) 1.21.5 or above
- [Docker](https://docs.docker.com/get-docker/)
- [Supabase CLI](https://github.com/supabase/cli)

## Installation

1. Clone this repo
2. Copy `config.local.yaml` in `config` and paste it in the same directory with `.local` removed from its name.
3. Run `go mod download` to download all the dependencies.

## Running

1. Run `supabase start` to start supabase.
2. Run `go run ./src/` to start server.

Ensure to run `supabase stop` to close supabase after finishing your code.
