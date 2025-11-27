import { Navigate } from "react-router-dom";
import { authService } from "../../services/auth.service";

interface GuestRouteProps {
    children: React.ReactNode;
}

export const GuestRoute = ({ children }: GuestRouteProps) => {
    const isAuthenticated = authService.isAuthenticated();

    if (isAuthenticated) {
        return <Navigate to="/dashboard" replace />;
    }

    return <>{children}</>;
};
