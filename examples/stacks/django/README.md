# Django Example

[![Built with Codex](https://www.khulnasoft/img/codex/shield_moon.svg)](https://www.khulnasoft/codex/docs/contributor-quickstart/)

[![Open In Codex.sh](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.sh/open/templates/django)

## How to Use

1. Install [Codex](https://www.khulnasoft/codex/docs/installing_codex/)
1. Run `codex shell` to install your packages and run the init_hook. This will activate your virtual environment and install Django.
1. Initialize PostgreSQL with `codex run initdb`.
1. Start the Postgres service by running `codex services up postgresql`. You can start it in the background using `codex services up -b postgresql`.
1. In the root directory, run `codex run create_db` to create the database and run your Django migrations
1. In the root directory, run `codex run server` to start the server. You can access the Django example at `localhost:8000`

## How to Create this Example from Scratch

### Setting up the Project

1. Install [Codex](https://www.khulnasoft/codex/docs/installing_codex/).
1. Run `codex create --template django` to create a new Codex project in your directory.
1. Install Python and PostgreSQL with `codex install`. This will also install the Codex plugins for pip (which sets up your .venv directory) and PostgreSQL.
1. Copy the requirements.txt and `todo_project` directory into the root folder of your project
1. Start a codex shell with `codex shell`. This will activate your virtual environment and install your requirements using the commands below.

   ```bash
   . $VENV_DIR/bin/activate
   pip install -r requirements.txt
   ```

   These lines are already added to your `init_hook` to automatically activate your venv.

### Setting up the Database

The Django example uses a database. To set up the database, we will first create a new PostgreSQL database cluster, create the `todo_db` and user, and run the Django migrations.

1. Initialize your Postgres database cluster with `codex run initdb`.

1. Start the Postgres service by running `codex services start postgres`

1. In your `codex shell`, create the empty `todo_db` database and user with the following commands.

   ```bash
   createdb todo_db
   psql todo_db -c "CREATE USER todo_user WITH PASSWORD 'secretpassword';"
   ```

   You can add this as a codex script in your `codex.json` file, so you can replicate the setup on other machines.

1. Run the Django migrations to create the tables in your database.

   ```bash
   python todo_project/manage.py makemigrations
   python todo_project/manage.py migrate
   ```

Your database is now ready to use. You can add these commands as a script in your `codex.json` if you want to automate them for future use. See `create_db` in the projects `codex.json` for an example.

### Running the Server

You can now start your Django server by running the following command.

   ```bash
   python todo_project/manage.py runserver
   ```

This should start the development server.

### Related Docs

* [Using Python with Codex](https://www.khulnasoft/codex/docs/codex_examples/languages/python/)
* [Using PostgreSQL with Codex](https://www.khulnasoft/codex/docs/codex_examples/stacks/django/)
