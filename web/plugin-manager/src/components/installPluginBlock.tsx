import { useState } from "react";
import axiosServices from "../services/axios";

import { Plugin } from "../../../shared/interfaces/IPlugin";

import SmallCardExtended from "./SmallCardExtended";
export interface InstallZipProps extends InstallProps {
    plugins: Plugin[];
}

export interface InstallProps {
    getPluginsInfo: () => void;
}

function InstallPlugin(p: InstallZipProps) {
    const [errorPluginInstall, setErrorPluginInstall] = useState("");

    const isValidUrl = (url: string) => {
        const regex = /^(http)[^\s]*\.zip$/i; // fragment locator
        return regex.test(url);
    };

    const installPluginHandler = async (url: string) => {
        if (!isValidUrl(url)) {
            setErrorPluginInstall("Invalid URL");
            return;
        }
        try {
            await axiosServices.installPlugin(url);
            p.getPluginsInfo();
        } catch (err: any) {
            console.error(err.response?.data?.message);
            setErrorPluginInstall("Error fetching plugin URL");
        }
    };

    return (
        <>
            <SmallCardExtended
                label2={"Install a plugin"}
                text3={"Install a plugin using .zip URL"}
                propsLabelButton={{
                    callbackToParent: function (data: string): void {
                        installPluginHandler(data);
                    },
                    label: "Plugin link",
                    placeholder: "Set .zip Url Link",
                    buttonValue: "Install",
                    error: errorPluginInstall,
                }}
            />
        </>
    );
}
export default InstallPlugin;
