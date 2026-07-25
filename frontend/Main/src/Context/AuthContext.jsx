import React from 'react';

const AuthContext = React.createContext(null);

export const AuthProvider = ({ children }) => {
  return (
    <AuthContext.Provider value={{ user: null }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  return React.useContext(AuthContext);
};