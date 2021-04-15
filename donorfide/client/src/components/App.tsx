import * as React from "react";

import "./App.sass"
import * as api from "../api"

import {BrowserRouter as Router} from "react-router-dom";
import Nav from "./Nav";
import Footer from "./Footer";
import LoadingScreen from "./LoadingScreen";

function useOrgContext() {
    const [isLoading, setIsLoading] = React.useState(true);
    const [contextData, setContextData] = React.useState(null as unknown as api.contextOrg);

    React.useEffect(() => {
        async function fetchContext() {
            setContextData(await api.GET<api.contextOrg>("context/org"))
            setIsLoading(false)
        }

        fetchContext()
    }, []);

    return [contextData, isLoading];
}

export const OrgContext = React.createContext(null as unknown as api.contextOrg)

export default function App() {
    const [contextData, isLoading] = useOrgContext()

    if (isLoading) {
        return <LoadingScreen/>
    }

    return <OrgContext.Provider value={contextData as api.contextOrg}>
        <Router>
            <Nav/>
            <Footer/>
        </Router>
    </OrgContext.Provider>
}