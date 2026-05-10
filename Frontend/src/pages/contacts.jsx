import { useEffect, useState } from "react";
import API from "../api/axios";

function Contacts() {
    const [contacts, setContacts] = useState([]);

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

    return (
        <div className="p-10">
            <h1 className="text-3xl font-bold mb-5">
                Contacts
            </h1>

            <div className="space-y-4">
                {contacts.map((contact) => (
                    <div
                        key={contact.id}
                        className="border p-4 rounded"
                    >
                        <p>Name: {contact.name}</p>
                        <p>Email: {contact.email}</p>
                        <p>Segment: {contact.segment}</p>
                    </div>
                ))}
            </div>
        </div>
    );
}

export default Contacts;