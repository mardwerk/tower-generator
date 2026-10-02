export async function api<T>(url: string, input?: unknown): Promise<T> {
  const response = await fetch(
    url,
    input === undefined
      ? undefined
      : {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(input),
        },
  );
  const value = (await response.json()) as T & { error?: string };
  if (!response.ok) throw new Error(value.error || 'The local request failed.');
  return value;
}
