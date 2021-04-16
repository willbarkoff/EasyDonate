const baseurl = "/api/";

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

	return paramStr;
}

async function handleResponse<T>(response: Response) {
	let json: response<T>;

	try {
		json = await response.json();
	} catch {
		throw new Error("network_issue");
	}

	if (json.status != "ok") {
		throw new Error(json.error);
	}

	return json as response<T>;
}

export async function GET<T>(url: string): Promise<T> {
	const response = await fetch(baseurl + url, {
		credentials: "include"
	});
	return (await handleResponse<T>(response)).data;
}

export async function POST<T>(url: string, params: Record<string, string>): Promise<T> {
	const response = await fetch(baseurl + url, {
		method: "POST",
		body: formURLEncode(params),
		headers: {
			"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8"
		},
		credentials: "include"
	});
	return (await handleResponse<T>(response)).data;
}

export interface contextOrg {
	stripe_pk: string
	name: string
	site: string
	imprint: string
	phone: string
	email: string
	donate_page_title: string
	donate_page_subtitle: string
	home_page_text: string
	credit_page_text: string
	disable_confetti: boolean
	donation_success_redirect: string
	donation_success_headline: string
	donation_success_subheadline: string
	donation_success_text: string
}

export interface paymentIntentResponse {
	status: string
	client_secret: string
}

export interface me {
	first_name: string
	last_name: string
	email: string
	level: number
}

export function errorToEnglish(error: string): string {
	switch (error) {
		case "internal_server_error":
			return "An internal server error occurred while processing your request. This error has automatically been logged. Please try again later.";
		case "network_issue":
			return "An issue occurred while attempting to communicate with the server. Please try again later.";
		case "bad_request":
			return "The server received a malformed request. Make sure that all of the parameters are correct.";
		case "missing_params":
			return "Required parameters were missing from the request.";
		case "unauthorized":
			return "You aren't authorized to perform that action.";
		case "logged_out":
			return "You need to log in to perform that action.";
		case "user_missing":
			return "Your user ID is missing from the database. Please contact us for assistance.";
		case "not_found":
			return "The page you requested could not be found.";
		case "invalid_login":
			return "The username or password you entered was incorrect.";
		default:
			return `An unknown error occurred: ${error}`;
	}
}