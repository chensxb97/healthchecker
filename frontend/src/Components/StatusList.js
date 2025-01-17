import React, { useEffect, useState } from "react";
import Timer from "./Timer";
import axios from "axios";

const StatusList = () => {
    const [statuses, setStatuses] = useState([]);
    const [loading, setLoading] = useState(true);
    const backendUrl = process.env.REACT_APP_BACKEND_URL

    const fetchStatuses = async () => {
        try {
            const response = await axios.get(`${backendUrl}/status`);
            setStatuses(response.data);
            console.log(response.data)
        } catch (error) {
            console.error("Error fetching statuses:", error);
        }
        setLoading(false);
    };

    useEffect(() => {
        fetchStatuses();
        const intervalId = setInterval(fetchStatuses, 5000); // Refresh every 5 seconds

        return () => clearInterval(intervalId); // Cleanup interval on unmount
    }, []);

    return (
        <div>
            {loading ? (
                <p>Loading...</p>
            ) : (
                <>
                    <Timer />
                    <ul>
                        {statuses.map((status) => (
                            <li key={status.url}>
                                <span>{status.url}</span>{" "}
                                <span
                                    style={{
                                        color: status.status === "live" ? "green" : "red",
                                    }}
                                >
                                    {status.status.toUpperCase()}
                                </span>
                            </li>
                        ))}
                    </ul>
                </>
            )}
        </div>
    );
};

export default StatusList;