import {
    LayoutDashboard,
    Users,
    Mail,
    Calendar,
    FileText,
    LogOut,
} from "lucide-react";

import { Link, useNavigate } from "react-router-dom";

function Sidebar() {
    const navigate = useNavigate();

    const logout = () => {
        localStorage.removeItem("token");
        navigate("/");
    };

    return (
        <div className="w-64 h-screen bg-black text-white fixed left-0 top-0 p-5">

            <h1 className="text-2xl font-bold mb-10">
                Email Tool
            </h1>

            <div className="space-y-4">

                <Link
                    to="/dashboard"
                    className="flex items-center gap-3 hover:bg-gray-800 p-3 rounded-lg"
                >
                    <LayoutDashboard size={20} />
                    Dashboard
                </Link>

                <Link
                    to="/contacts"
                    className="flex items-center gap-3 hover:bg-gray-800 p-3 rounded-lg"
                >
                    <Users size={20} />
                    Contacts
                </Link>

                <Link
                    to="/campaigns"
                    className="flex items-center gap-3 hover:bg-gray-800 p-3 rounded-lg"
                >
                    <Mail size={20} />
                    Campaigns
                </Link>

                <Link
                    to="/schedule"
                    className="flex items-center gap-3 hover:bg-gray-800 p-3 rounded-lg"
                >
                    <Calendar size={20} />
                    Schedule
                </Link>

                <Link
                    to="/logs"
                    className="flex items-center gap-3 hover:bg-gray-800 p-3 rounded-lg"
                >
                    <FileText size={20} />
                    Logs
                </Link>

                <button
                    onClick={logout}
                    className="flex items-center gap-3 hover:bg-red-600 p-3 rounded-lg w-full text-left"
                >
                    <LogOut size={20} />
                    Logout
                </button>

            </div>
        </div>
    );
}

export default Sidebar;