import React, { useEffect, useState } from "react";

const Timer = ({ customTime }) => {
    const timeLimit = 30 || customTime // default 30 seconds, unless specified in component prop
    const [countdown, setCountdown] = useState(timeLimit)
    const [lastRefreshed,] = useState(new Date())

    useEffect(() => {
        const interval = setInterval(() => {
            setCountdown((prev) => {
                if (prev === 1) {
                    return timeLimit
                }
                return prev - 1
            })
        }, 1000)

        return () => clearInterval(interval)
    }, [timeLimit]);

    return (
        <div style={{ fontSize: '16px' }}>
            <p>Time until next refresh: <span style={{ color: 'darkred' }}>{countdown}</span></p>
            <p>Last refreshed: <span style={{ color: 'darkgreen' }}>{lastRefreshed.toLocaleTimeString()}</span></p>
        </div>
    );
};

export default Timer;
