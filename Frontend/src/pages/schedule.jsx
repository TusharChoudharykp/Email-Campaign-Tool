import { useEffect, useState } from "react";

import API from "../api/axios";
import DashboardLayout from "../components/DashboardLayout";

function Schedule() {

    const [campaigns, setCampaigns] = useState([]);

    const [schedules, setSchedules] = useState([]);

    const [form, setForm] = useState({
        campaign_id: "",
        segment: "",
        send_at: "",
    });

    useEffect(() => {
        fetchCampaigns();
        fetchSchedules();
    }, []);

    // =========================
    // FETCH CAMPAIGNS
    // =========================

    const fetchCampaigns = async () => {

        try {

            const response = await API.get("/campaigns");

            setCampaigns(response.data.data);

        } catch (err) {
            console.log(err);
        }
    };

    // =========================
    // FETCH SCHEDULES
    // =========================

    const fetchSchedules = async () => {

        try {

            const response = await API.get("/schedules");

            setSchedules(response.data.data);

        } catch (err) {
            console.log(err);
        }
    };

    // =========================
    // CREATE SCHEDULE
    // =========================

    const createSchedule = async (e) => {

        e.preventDefault();

        try {

            // convert datetime-local format
            const formattedDate = form.send_at.replace("T", " ") + ":00";

            await API.post("/schedules", {
                ...form,
                campaign_id: Number(form.campaign_id),
                send_at: formattedDate,
            });

            alert("Campaign scheduled successfully");

            setForm({
                campaign_id: "",
                segment: "",
                send_at: "",
            });

            fetchSchedules();

        } catch (err) {

            console.log("FULL ERROR:", err);

            console.log("BACKEND RESPONSE:", err.response);

            console.log("BACKEND DATA:", err.response?.data);

            alert(
                err.response?.data?.message ||
                "Failed to schedule campaign"
            );
        }
    };

    return (
        <DashboardLayout>

            <h1 className="text-3xl font-bold mb-8">
                Schedule Campaigns
            </h1>

            {/* =========================
                SCHEDULE FORM
            ========================== */}

            <form
                onSubmit={createSchedule}
                className="bg-white p-6 rounded-2xl shadow mb-8"
            >

                <h2 className="text-xl font-bold mb-5">
                    Schedule New Campaign
                </h2>

                <div className="grid grid-cols-1 md:grid-cols-3 gap-4">

                    {/* SELECT CAMPAIGN */}

                    <select
                        value={form.campaign_id}
                        onChange={(e) =>
                            setForm({
                                ...form,
                                campaign_id: e.target.value,
                            })
                        }
                        className="border p-3 rounded-lg"
                        required
                    >

                        <option value="">
                            Select Campaign
                        </option>

                        {campaigns.map((campaign) => (

                            <option
                                key={campaign.id}
                                value={campaign.id}
                            >
                                {campaign.name}
                            </option>

                        ))}

                    </select>

                    {/* SEGMENT */}

                    <input
                        type="text"
                        placeholder="Segment (optional)"
                        value={form.segment}
                        onChange={(e) =>
                            setForm({
                                ...form,
                                segment: e.target.value,
                            })
                        }
                        className="border p-3 rounded-lg"
                    />

                    {/* DATETIME */}

                    <input
                        type="datetime-local"
                        value={form.send_at}
                        onChange={(e) =>
                            setForm({
                                ...form,
                                send_at: e.target.value,
                            })
                        }
                        className="border p-3 rounded-lg"
                        required
                    />

                </div>

                <button
                    className="bg-black text-white px-6 py-3 rounded-lg mt-5 hover:bg-gray-800"
                >
                    Schedule Campaign
                </button>

            </form>

            {/* =========================
                SCHEDULE TABLE
            ========================== */}

            <div className="bg-white rounded-2xl shadow overflow-hidden">

                <table className="w-full">

                    <thead className="bg-black text-white">

                        <tr>

                            <th className="p-4 text-left">
                                Campaign ID
                            </th>

                            <th className="p-4 text-left">
                                Segment
                            </th>

                            <th className="p-4 text-left">
                                Send At
                            </th>

                            <th className="p-4 text-left">
                                Status
                            </th>

                        </tr>

                    </thead>

                    <tbody>

                        {schedules.map((schedule) => (

                            <tr
                                key={schedule.id}
                                className="border-b"
                            >

                                <td className="p-4">
                                    {schedule.campaign_id}
                                </td>

                                <td className="p-4">
                                    {schedule.segment || "All"}
                                </td>

                                <td className="p-4">
                                    {schedule.send_at}
                                </td>

                                <td className="p-4">
                                    <span
                                        className={`
                                            px-3 py-1 rounded-full text-white text-sm
                                            ${schedule.status === "completed"
                                                ? "bg-green-500"
                                                : schedule.status === "failed"
                                                    ? "bg-red-500"
                                                    : "bg-yellow-500"
                                            }
                                        `}
                                    >
                                        {schedule.status}
                                    </span>
                                </td>

                            </tr>

                        ))}

                    </tbody>

                </table>

            </div>

        </DashboardLayout>
    );
}

export default Schedule;