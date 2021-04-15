const baseurl = "/api/"

export interface response<T> {
    status: string
    error?: string
    data: T
}

function formURLEncode(body: Record<string, string>): string {
    if (!body) {
        return "";
    }

    let paramStr = "";

    let first = true;
    for (const key in body) {
        if (body.hasOwnProperty(key)) {
            const value = body[key];
            if (first) {
                first = false;
            } else {
                paramStr += "&";
            }
            paramStr += key;
            paramStr += "=";
            paramStr += encodeURIComponent(value);
        }
    }

    return paramStr;
};

async function handleResponse<T>(response: Response) {
    if (!response.ok) {
        throw new Error("network_issue")
    }

    let json = await response.json()
    console.log(json)
    if (json.status != "ok") {
        throw new Error(json.error)
    }

    return json as response<T>
}

export async function GET<T>(url: string) {
    let response = await fetch(baseurl + url, {
        credentials: "include"
    });
    return (await handleResponse<T>(response)).data
}

export async function POST<T>(url: string, params: Record<string, string>) {
    let response = await fetch(baseurl + url, {
        method: "POST",
        body: formURLEncode(params),
        headers: {
            "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8"
        },
        credentials: "include"
    })
    return (await handleResponse<T>(response)).data
}

export interface contextOrg {
    stripe_pk: string
    name: string
    site: string
    imprint: string
    phone: string
    email: string
}