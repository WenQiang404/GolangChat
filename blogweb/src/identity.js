import React, { createContext, useState } from 'react';

const UserContext = createContext();

export const UserProvider = ({ children }) => {
    const [identity, setIdentity] = useState(null);

    const login = (userIdentity) => {
        setIdentity(userIdentity);
    };

    const logout = () => {
        setIdentity(null);
    };

    return (
        <UserContext.Provider value={{ identity, login, logout }}>
            {children}
        </UserContext.Provider>
    );
};

export default UserContext;