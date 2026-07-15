import React from 'react'
import { Navigate } from 'react-router-dom'

const ProtectedRoute = ({ children }) => {
  // Пока всегда показываем детей (потом добавим проверку)
  return children
}

export default ProtectedRoute