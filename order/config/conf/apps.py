from config.env import env

APPS = [
    "channels",
    "cacheops",
    "django_ckeditor_5",
    "drf_spectacular",
    "rest_framework",
    "corsheaders",
    "django_filters",
    "django_redis",
    "rest_framework_simplejwt",
    "django_core",
]

if env.str("PROJECT_ENV") == "debug":
    APPS += [
        "silk",
    ]
