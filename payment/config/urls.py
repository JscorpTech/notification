from django.contrib import admin
from django.urls import path, include

urlpatterns = [
    path("payment/api/", include("api.urls")),
    path("payment/admin/", admin.site.urls),
]
