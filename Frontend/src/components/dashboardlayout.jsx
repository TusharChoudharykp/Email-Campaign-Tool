import Sidebar from "./Sidebar";
import Navbar from "./Navbar";
import PropTypes from "prop-types";

function DashboardLayout({ children }) {
    return (
        <div className="flex">

            <Sidebar />

            <div className="ml-64 w-full min-h-screen bg-gray-100">

                <Navbar />

                <div className="p-6">
                    {children}
                </div>

            </div>

        </div>
    );
}

DashboardLayout.propTypes = {
    children: PropTypes.node.isRequired,
};

export default DashboardLayout;