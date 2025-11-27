import { usePageTitle } from "../hooks/usePageTitle"

export const DashboardPage = () => {
    usePageTitle("Dashboard")

    return (
        <div>
            <h1>Dashboard</h1>
        </div>
    )
}