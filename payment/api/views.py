from rest_framework.views import APIView
from rest_framework.response import Response


class HealthView(APIView):
    def get(self, *args, **kwargs):
        return Response(data={"detail": "OK"})