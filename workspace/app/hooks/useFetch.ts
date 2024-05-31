import useSWR from "swr";

const fetcher = (url: string) => fetch(url, { cache: 'no-store' }).then(r => r.json())

export function useFetch(url: string) {
  const { data, error, mutate, isLoading, isValidating } = useSWR(url, fetcher, {
    refreshInterval: 5000,
    revalidateOnMount: true,
    revalidateIfStale: true,
    revalidateOnReconnect: true
  });

  return { data, error, mutate, isLoading, isValidating };
}
