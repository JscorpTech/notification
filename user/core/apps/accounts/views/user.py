from rest_framework.viewsets import GenericViewSet
from django.contrib.auth import get_user_model
from rest_framework.mixins import RetrieveModelMixin, ListModelMixin
from ..serializers import UserSerializer


class UserViewSet(RetrieveModelMixin, ListModelMixin, GenericViewSet):
    model = get_user_model()
    serializer_class = UserSerializer
