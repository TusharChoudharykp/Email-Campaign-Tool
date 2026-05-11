import { useEffect, useState } from "react";

import API from "../api/axios";
import DashboardLayout from "../components/DashboardLayout";

function Campaigns() {

    const [campaigns, setCampaigns] = useState([]);

    const [form, setForm] = useState({
        name: "",
        subject: "",
        template: "",
        status: "draft",
    });

    const [segment, setSegment] = useState("");

    useEffect(() => {
        fetchCampaigns();
    }, []);


    // FETCH CAMPAIGNS
    const fetchCampaigns = async () => {
        try {

            const response = await API.get("/campaigns");

            setCampaigns(response.data.data);

        } catch (err) {
            console.log(err);
        }
    };


    // CREATE CAMPAIGN
    const createCampaign = async (e) => {
        e.preventDefault();

        try {

            await API.post("/campaigns", form);

            setForm({
                name: "",
                subject: "",
                template: "",
                status: "draft",
            });

            fetchCampaigns();

        } catch (err) {
            console.log(err);
        }
    };


    // DELETE CAMPAIGN
    const deleteCampaign = async (id) => {

        const confirmDelete = window.confirm(
            "Are you sure you want to delete this campaign?"
        );

        if (!confirmDelete) return;

        try {

            await API.delete(`/campaigns/${id}`);

            fetchCampaigns();

        } catch (err) {
            console.log(err);
        }
    };


    // SEND CAMPAIGN
    const sendCampaign = async (id) => {

        try {

            await API.post(
                `/campaigns/${id}/send`,
                {
                    segment,
                }
            );

            alert("Campaign sent successfully");

            fetchCampaigns();

        } catch (err) {
            console.log(err);
        }
    };

    return (
        <DashboardLayout>

            <h1 className="text-3xl font-bold mb-8">
                Campaigns
            </h1>

            {/* =========================
                CREATE FORM
            ========================== */}

            <form
                onSubmit={createCampaign}
                className="bg-white p-6 rounded-2xl shadow mb-8"
            >

                <h2 className="text-xl font-bold mb-5">
                    Create Campaign
                </h2>

                <div className="grid grid-cols-1 gap-4">

                    <input
                        type="text"
                        placeholder="Campaign Name"
                        value={form.name}
                        onChange={(e) =>
                            setForm({
                                ...form,
                                name: e.target.value,
                            })
                        }
                        className="border p-3 rounded-lg"
                        required
                    />

                    <input
                        type="text"
                        placeholder="Subject"
                        value={form.subject}
                        onChange={(e) =>
                            setForm({
                                ...form,
                                subject: e.target.value,
                            })
                        }
                        className="border p-3 rounded-lg"
                        required
                    />

                    <textarea
                        rows="6"
                        placeholder="Email Template"
                        value={form.template}
                        onChange={(e) =>
                            setForm({
                                ...form,
                                template: e.target.value,
                            })
                        }
                        className="border p-3 rounded-lg"
                        required
                    />

                </div>

                <button
                    className="bg-black text-white px-6 py-3 rounded-lg mt-5 hover:bg-gray-800"
                >
                    Create Campaign
                </button>

            </form>

            {/* =========================
                SEND BY SEGMENT
            ========================== */}

            <div className="bg-white p-6 rounded-2xl shadow mb-8">

                <h2 className="text-xl font-bold mb-5">
                    Send By Segment
                </h2>

                <input
                    type="text"
                    placeholder="Enter Segment (example: premium)"
                    value={segment}
                    onChange={(e) =>
                        setSegment(e.target.value)
                    }
                    className="border p-3 rounded-lg w-full"
                />

            </div>

            {/* =========================
                CAMPAIGNS TABLE
            ========================== */}

            <div className="bg-white rounded-2xl shadow overflow-hidden">

                <table className="w-full">

                    <thead className="bg-black text-white">

                        <tr>

                            <th className="p-4 text-left">
                                Name
                            </th>

                            <th className="p-4 text-left">
                                Subject
                            </th>

                            <th className="p-4 text-left">
                                Status
                            </th>

                            <th className="p-4 text-left">
                                Actions
                            </th>

                        </tr>

                    </thead>

                    <tbody>

                        {campaigns.map((campaign) => (

                            <tr
                                key={campaign.id}
                                className="border-b"
                            >

                                <td className="p-4">
                                    {campaign.name}
                                </td>

                                <td className="p-4">
                                    {campaign.subject}
                                </td>

                                <td className="p-4">
                                    {campaign.status}
                                </td>

                                <td className="p-4 flex gap-3">

                                    {/* SEND BUTTON */}

                                    <button
                                        onClick={() =>
                                            sendCampaign(campaign.id)
                                        }
                                        className="bg-green-500 text-white px-4 py-2 rounded-lg hover:bg-green-600"
                                    >
                                        Send
                                    </button>

                                    {/* DELETE BUTTON */}

                                    <button
                                        onClick={() =>
                                            deleteCampaign(campaign.id)
                                        }
                                        className="bg-red-500 text-white px-4 py-2 rounded-lg hover:bg-red-600"
                                    >
                                        Delete
                                    </button>

                                </td>

                            </tr>

                        ))}

                    </tbody>

                </table>

            </div>

        </DashboardLayout>
    );
}

export default Campaigns;