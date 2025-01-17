import React, { useEffect, useState } from "react"
import Timer from "./Timer"
import axios from "axios"

const StatusList = () => {
    const [statuses, setStatuses] = useState([])
    const [loading, setLoading] = useState(true)

    const backendUrl = process.env.REACT_APP_BACKEND_URL
    const timeLimit = 30000

    const fetchStatuses = async () => {
        try {
            const response = await axios.get(`${backendUrl}/status`)
            setStatuses(response.data)
        } catch (error) {
            console.error("Error fetching statuses:", error)
        }
        setLoading(false)
    }

    useEffect(() => {
        fetchStatuses()
        const intervalId = setInterval(fetchStatuses, timeLimit)
        return () => clearInterval(intervalId)
        // eslint-disable-next-line
    }, [])

    return (
        <div>
            {loading ? (
                <p>Loading...</p>
            ) : (
                <>
                    <Timer customTime={timeLimit} />
                    <div style={{ display: "flex", justifyContent: "center", marginTop: "20px" }}>
                        <table style={{ borderCollapse: "collapse", border: "3px solid grey" }}>
                            <thead>
                                <tr>
                                    <th style={{ padding: "8px", border: "2px solid #ddd" }}>Endpoint</th>
                                    <th style={{ padding: "8px", border: "2px solid #ddd" }}>Current Status</th>
                                </tr>
                            </thead>
                            <tbody>
                                {statuses.map((status) => (
                                    <tr key={status?.url}>
                                        <td style={{ padding: "8px", border: "2px solid #ddd" }}>
                                            {status?.url}
                                        </td>
                                        <td
                                            style={{
                                                padding: "8px",
                                                border: "2px solid #ddd",
                                                color: status?.status === "live" ? "green" : "red",
                                            }}
                                        >
                                            {status?.status?.toUpperCase()}
                                        </td>
                                    </tr>
                                ))}
                            </tbody>
                        </table>
                    </div>
                </>
            )
            }
        </div >
    )
}

export default StatusList