'use client'

import { useFetch } from "@/app/hooks/useFetch";
import Dashboard from "@/components/dashboard";

export default function Painel() {
  const { data, error, isLoading } = useFetch('/api/watcher/84cbfa32-cfa2-4f75-99d6-5423eca60985')
  if (error) return <div>falhou ao carregar</div>
  return <Dashboard watcher={data} isLoading={isLoading} />
}
