from rest_framework.permissions import BasePermission
from rest_framework.exceptions import AuthenticationFailed
from rest_framework.authentication import TokenAuthentication

class TokenRequiredPermission(BasePermission):
    """
    Custom permission to ensure that the request is authenticated
    with a valid token.
    """

    def has_permission(self, request, view):
        # Use the TokenAuthentication to authenticate the token
        authentication = TokenAuthentication()
        user_auth_tuple = authentication.authenticate(request)
        if user_auth_tuple is None:
            raise AuthenticationFailed("Invalid or missing token")
        return True
