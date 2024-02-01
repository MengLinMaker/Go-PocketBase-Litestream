# PocketBase + Litestream on fly.io

This repo uses PocketBase as a Go framework. Backup is done via Litestream to an S3 compatible endpoint. The demo backend can be deployed to fly.io with a persistent volume.

## Usage

### Prerequisites

You'll need to have an S3-compatible store to connect to. Please see the [Litestream Guides](https://litestream.io/guides/) to get set up on your preferred object store.

## Local Development

Since this is using PocketBase as a Go framework, you can run this locally with `go run *.go serve --http "localhost:8080"` from the `src` directory.

Note: do not backup to Litestream during development.

## Deploying to production

Please create a `.env` variable with the following content:

```
LITESTREAM_ACCESS_KEY_ID=
LITESTREAM_SECRET_ACCESS_KEY=
S3_ENDPOINT=
S3_DATA_BUCKET=
S3_LOGS_BUCKET=
STAGE=PROD
```

STAGE=PROD will enable litestream backup. Any other varaibles will not.

Fly.io deployment configuration is specified in `fly.toml`.
