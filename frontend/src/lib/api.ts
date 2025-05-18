const BASE_URL = 'http://localhost:9400';

// Helper: build headers with optional token
function getAuthHeaders(): HeadersInit {
	const token = localStorage.getItem('access_token');
	const headers: HeadersInit = {
		'Content-Type': 'application/json',
	};
	if (token) {
		headers['Authorization'] = `Bearer ${token}`;
	}
	return headers;
}

export async function apiGet<T>(path: string): Promise<T> {
	const res = await fetch(`${BASE_URL}${path}`, {
		method: 'GET',
		headers: getAuthHeaders()
	});
	if (!res.ok) throw new Error(await res.text());
	return await res.json();
}

export async function apiPost<T>(path: string, body: unknown): Promise<T> {
	const res = await fetch(`${BASE_URL}${path}`, {
		method: 'POST',
		headers: getAuthHeaders(),
		body: JSON.stringify(body)
	});
	if (!res.ok) throw new Error(await res.text());
	return await res.json();
}

export async function apiPut<T>(path: string, body: unknown): Promise<T> {
	const res = await fetch(`${BASE_URL}${path}`, {
		method: 'PUT',
		headers: getAuthHeaders(),
		body: JSON.stringify(body)
	});
	if (!res.ok) throw new Error(await res.text());
	return await res.json();
}

export async function apiGetAuth(path: string) {
	const res = await fetch(`${BASE_URL}${path}`, {
		method: 'GET',
		headers: getAuthHeaders()
	});

	if (!res.ok) throw new Error(await res.text());

	const response = await res.json();
	const role_id = response?.data?.user?.role_id;

	if (role_id) {
		localStorage.setItem('role_id', role_id);
	}

	return response;
}
