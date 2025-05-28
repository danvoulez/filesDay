export type LogLine = {
  id: string;           // UUID
  tenant: string;
  who: string;
  did: string;
  this: any;
  when: string;
  confirmed_by: string;
  if_ok: string;
  if_doubt: string;
  if_not: string;
  status: string;
  valid: boolean;
};