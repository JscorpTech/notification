#!/bin/bash

python3 manage.py collectstatus --no-input
python3 manage.py migrate --no-input
python3 manage.py runserver 0.0.0.0:8000