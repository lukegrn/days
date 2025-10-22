# Days - A photo project by Luke Green

Days is a purpose built photo website for me to use for a year in photos project
I will be doing in 2026. If you want to use it for the same, you'll just need to
make a few changes in the `templates` folder. Namely, update
`templates/base.html` to use your name, `templates/about.html` to whatever
content you would like to have there, and `templates/index.html` to have your
name there.

## Customization and deployment

Days expects three environment variables to be present:

- `DAYS_DB_FILE` - A sqlite database file to store data in. Defaults to
  `/tmp/days_db.sqlite`, do not use the default in production as it will be
  erased on a reboot.
- `DAYS_PORT` - The port to run the server on, defaults to `:8080`.
- `DAYS_PW` - A password hash for uploading files. Authentication is rudementary
  and requires this password on every upload. `cmd/serve` will panic without
  this being provided. You can generate yours with
  `make build && ./bin/hashgen <your pw here>`. If you aren't okay with that
  level of authentication, don't use this project.

## Usage

- Upload files at https://yoururl.example/upload.
- You can only upload one file for each day, and they will be displayed on the
  index page in chronological order.
- Clicking on an image preview will take you to the full size version of that
  image, with left and right arrows for navigation.
