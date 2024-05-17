'use client'

import { useEffect, useState } from "react";
import Dashboard from "@/components/dashboard";
import useSWR from "swr";

const fetcher = (url: string) => fetch(url).then(r => r.json())

export default function Painel() {
  const { data, error, isLoading } = useSWR('/api/watcher/84cbfa32-cfa2-4f75-99d6-5423eca60985', fetcher)

  if (error) return <div>falhou ao carregar</div>
  if (isLoading) return <div>carregando...</div>

  return <Dashboard watcher={data} />
}
