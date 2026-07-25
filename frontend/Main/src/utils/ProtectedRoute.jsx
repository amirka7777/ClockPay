import React from 'react';
import { Navigate } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';

const ProtectedRoute = ({ children }) => {
  const { user } = useAuth();

  // Если пользователя нет — отправляем на логин
  if (!user) {
    return <Navigate to="/login" replace />;
  }

  // Если пользователь есть — показываем страницу
  return children;
};

export default ProtectedRoute;