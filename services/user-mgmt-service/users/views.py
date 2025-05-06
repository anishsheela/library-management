from rest_framework import generics, permissions
from rest_framework.response import Response
from django.contrib.auth.models import User
from django.contrib.auth import authenticate
from rest_framework.permissions import IsAuthenticated
from .serializers import UserSerializer, RegisterSerializer
from rest_framework_simplejwt.tokens import RefreshToken
from django.http import JsonResponse
import json
import os
from django.http import JsonResponse
from cryptography.hazmat.primitives import serialization
from cryptography.hazmat.backends import default_backend
from rest_framework.generics import RetrieveAPIView
import json

# Register API
class RegisterAPI(generics.GenericAPIView):
    serializer_class = RegisterSerializer
    permission_classes = [permissions.AllowAny]

    def post(self, request):
        serializer = self.get_serializer(data=request.data)
        serializer.is_valid(raise_exception=True)
        user = serializer.save()

        refresh = RefreshToken.for_user(user)

        return Response({
            "user": UserSerializer(user).data,
            "refresh": str(refresh),
            "access": str(refresh.access_token),
        })

# Login API 
class LoginAPI(generics.GenericAPIView):
    permission_classes = [permissions.AllowAny]

    def post(self, request):
        username = request.data.get("username")
        password = request.data.get("password")
        user = authenticate(username=username, password=password)
        if user:
            refresh = RefreshToken.for_user(user)
            return Response({
                "refresh": str(refresh),
                "access": str(refresh.access_token),
                "user": UserSerializer(user).data
            })
        else:
            return Response({"error": "Invalid Credentials"}, status=400)

# Protected User List API
class UserListAPI(generics.ListAPIView):
    queryset = User.objects.all()
    serializer_class = UserSerializer
    permission_classes = [IsAuthenticated]  # JWT will handle authentication

class UserDetailAPI(RetrieveAPIView):
    queryset = User.objects.all()
    serializer_class = UserSerializer
    lookup_field = 'id'  # Change this if you want to use another field for lookup, e.g., username

def load_public_key():
    with open('keys/public_key.pem', 'rb') as key_file:
        public_key = serialization.load_pem_public_key(
            key_file.read(),
            backend=default_backend()
        )
    return public_key

def jwks_view(request):
    # Load public key and convert to JWKS format
    public_key = load_public_key()

    jwks = {
        "keys": [
            {
                "kty": "RSA",
                "kid": "my-key-id",  # A unique key identifier (can be any string)
                "use": "sig",  # This key is for signing
                "alg": "RS256",
                "n": public_key.public_numbers().n,
                "e": public_key.public_numbers().e
            }
        ]
    }

    return JsonResponse(jwks)