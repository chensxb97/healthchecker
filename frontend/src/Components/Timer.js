import React, { useEffect, useState } from "react";
const Timer = ({ customTime }) => {
    const timeLimit = 5 || customTime
    const [countdown, setCountdown] = useState(timeLimit)
    const [lastRefreshed, setLastRefreshed] = useState(new Date())

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
    }, []);

    return (
        <div>
            <p>Time until next refresh: {countdown}s</p>
            <p>Last refreshed: {lastRefreshed.toLocaleTimeString()}</p>
        </div>
    );
};

export default Timer;
