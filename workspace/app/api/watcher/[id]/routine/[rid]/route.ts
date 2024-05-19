import { NextRequest, NextResponse } from 'next/server'

export async function GET(request: NextRequest, { params }: {
  params: {
    id: string
    rid: string
  }
}) {
  const res = await fetch("http://localhost:8181/v1/watcher/" + params.id + "/routine/" + params.rid + "/obtain")
  return NextResponse.json(await res.json());
}
