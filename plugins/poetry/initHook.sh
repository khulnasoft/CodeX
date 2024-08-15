#!/bin/sh

poetry env use $(command -v python) --directory="${CODEX_PYPROJECT_DIR:-$CODEX_DEFAULT_PYPROJECT_DIR}" --no-interaction >&2
