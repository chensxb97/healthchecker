import React, { useEffect, useState } from "react";

const Timer = ({ customTime }) => {
    const timeLimit = customTime / 1000 || 5 // Set default as 5 seconds
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
            <p>Time until next refresh: <span style={{ color: 'royalblue' }}>{countdown}</span></p>
            <p>Last refreshed: <span style={{ color: 'royalblue' }}>
                {lastRefreshed.toLocaleString('en-US', {
                    year: 'numeric', // Full year (e.g., 2025)
                    month: 'short',   // Full name of the month (e.g., January)
                    day: 'numeric',  // Numeric day of the month (e.g., 18)
                    hour: 'numeric', // Hour in 12-hour format
                    minute: '2-digit', // Minutes with leading zero if needed
                    hour12: true     // 12-hour format with AM/PM
                })}</span></p>
        </div>
    );
};

export default Timer;
