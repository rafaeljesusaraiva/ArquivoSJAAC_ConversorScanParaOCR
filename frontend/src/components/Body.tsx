import React from 'react';

interface BodyProps {
    className: string;
    children: React.ReactNode;
}

const bodyStyle = {
    height: "calc(100vh - 40px)",
}

const Body: React.FC<BodyProps> = ({ className, children }) => {
    return (
        <div style={bodyStyle} className={className}>
            {children}
        </div>
    );
};

export default Body;