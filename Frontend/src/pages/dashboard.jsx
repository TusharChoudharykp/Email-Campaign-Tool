import { useEffect, useState } from "react";

import DashboardLayout from "../components/DashboardLayout";
import API from "../api/axios";

function Dashboard() {

    const [stats, setStats] = useState({
        total_contacts: 0,
        total_campaigns: 0,
        sent_emails: 0,
        failed_emails: 0,
    });

    useEffect(() => {
        fetchDashboardStats();
    }, []);

    const fetchDashboardStats = async () => {
        try {

            const response = await API.get("/dashboard/stats");

            setStats(response.data.data);

        } catch (err) {
            console.log(err);
        }
    };

    return (
        <DashboardLayout>

            <h1 className="text-3xl font-bold mb-8">
                Dashboard Overview
            </h1>

            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">

                <div className="bg-white p-6 rounded-2xl shadow">

                    <h2 className="text-gray-500">
                        Total Contacts
                    </h2>

                    <p className="text-3xl font-bold mt-3">
                        {stats.total_contacts}
                    </p>

                </div>

                <div className="bg-white p-6 rounded-2xl shadow">

                    <h2 className="text-gray-500">
                        Campaigns
                    </h2>

                    <p className="text-3xl font-bold mt-3">
                        {stats.total_campaigns}
                    </p>

                </div>

                <div className="bg-white p-6 rounded-2xl shadow">

                    <h2 className="text-gray-500">
                        Emails Sent
                    </h2>

                    <p className="text-3xl font-bold mt-3">
                        {stats.sent_emails}
                    </p>

                </div>

                <div className="bg-white p-6 rounded-2xl shadow">

                    <h2 className="text-gray-500">
                        Failed Emails
                    </h2>

                    <p className="text-3xl font-bold mt-3 text-red-500">
                        {stats.failed_emails}
                    </p>

                </div>

            </div>

        </DashboardLayout>
    );
}

export default Dashboard;