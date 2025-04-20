from config.env import env

APPS = [
    "cacheops",
    "drf_spectacular",
    "rest_framework",
    "corsheaders",
    "django_filters",
    "django_redis",
    "rest_framework_simplejwt",
    "django_core",
    "core.apps.accounts.apps.AccountsConfig",
]

if env.str("PROJECT_ENV") == "debug":
    APPS += [
        "silk",
    ]
