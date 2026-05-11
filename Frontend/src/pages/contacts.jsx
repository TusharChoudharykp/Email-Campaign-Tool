import { useEffect, useState } from "react";

import API from "../api/axios";
import DashboardLayout from "../components/DashboardLayout";

function Contacts() {

    const [contacts, setContacts] = useState([]);

    const [form, setForm] = useState({
        name: "",
        email: "",
        segment: "",
    });

    useEffect(() => {
        fetchContacts();
    }, []);

    const fetchContacts = async () => {
        try {

            const response = await API.get("/contacts");

            setContacts(response.data.data);

        } catch (err) {
            console.log(err);
        }
    };

    // =========================
    // Create Contact
    // =========================

    const createContact = async (e) => {
        e.preventDefault();

        try {

            await API.post("/contacts", form);

            setForm({
                name: "",
                email: "",
                segment: "",
            });

            fetchContacts();

        } catch (err) {
            console.log(err);
        }
    };

    // =========================
    // Delete Contact
    // =========================

    const deleteContact = async (id) => {

        const confirmDelete = window.confirm(
            "Are you sure you want to delete this contact?"
        );

        if (!confirmDelete) return;

        try {

            await API.delete(`/contacts/${id}`);

            fetchContacts();

        } catch (err) {
            console.log(err);
        }
    };

    return (
        <DashboardLayout>

            <h1 className="text-3xl font-bold mb-8">
                Contacts
            </h1>

            {/* =========================
                CREATE CONTACT FORM
            ========================== */}

            <form
                onSubmit={createContact}
                className="bg-white p-6 rounded-2xl shadow mb-8"
            >

                <h2 className="text-xl font-bold mb-5">
                    Add Contact
                </h2>

                <div className="grid grid-cols-1 md:grid-cols-3 gap-4">

                    <input
                        type="text"
                        placeholder="Name"
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
                        type="email"
                        placeholder="Email"
                        value={form.email}
                        onChange={(e) =>
                            setForm({
                                ...form,
                                email: e.target.value,
                            })
                        }
                        className="border p-3 rounded-lg"
                        required
                    />

                    <input
                        type="text"
                        placeholder="Segment"
                        value={form.segment}
                        onChange={(e) =>
                            setForm({
                                ...form,
                                segment: e.target.value,
                            })
                        }
                        className="border p-3 rounded-lg"
                        required
                    />

                </div>

                <button
                    className="bg-black text-white px-6 py-3 rounded-lg mt-5 hover:bg-gray-800"
                >
                    Create Contact
                </button>

            </form>

            {/* =========================
                CONTACT TABLE
            ========================== */}

            <div className="bg-white rounded-2xl shadow overflow-hidden">

                <table className="w-full">

                    <thead className="bg-black text-white">

                        <tr>
                            <th className="p-4 text-left">
                                Name
                            </th>

                            <th className="p-4 text-left">
                                Email
                            </th>

                            <th className="p-4 text-left">
                                Segment
                            </th>

                            <th className="p-4 text-left">
                                Action
                            </th>
                        </tr>

                    </thead>

                    <tbody>

                        {contacts.map((contact) => (

                            <tr
                                key={contact.id}
                                className="border-b"
                            >

                                <td className="p-4">
                                    {contact.name}
                                </td>

                                <td className="p-4">
                                    {contact.email}
                                </td>

                                <td className="p-4">
                                    {contact.segment}
                                </td>

                                <td className="p-4">

                                    <button
                                        onClick={() =>
                                            deleteContact(contact.id)
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

export default Contacts;