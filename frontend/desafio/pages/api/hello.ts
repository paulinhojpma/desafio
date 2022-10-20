// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next'

type Date = {
  name: string
}

export default function handler(
  req: NextApiRequest,
  res: NextApiResponse<Date>
) {
  res.status(200).json({ name: 'John Doe' })
}
