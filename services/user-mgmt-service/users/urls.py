from django.urls import path
from .views import RegisterAPI, LoginAPI, UserListAPI, UserDetailAPI, jwks_view
from rest_framework_simplejwt.views import TokenObtainPairView, TokenRefreshView

urlpatterns = [
    path('register/', RegisterAPI.as_view(), name='register'),
    path('login/', LoginAPI.as_view(), name='login'),  # or use TokenObtainPairView
    path('users/', UserListAPI.as_view(), name='user-list'),
    path('users/<int:id>/', UserDetailAPI.as_view(), name='user-detail'),
    path('api/token/', TokenObtainPairView.as_view(), name='token_obtain_pair'),
    path('api/token/refresh/', TokenRefreshView.as_view(), name='token_refresh'),
    path('.well-known/jwks.json', jwks_view),
]
