export interface LogLine {
  who: string;
  did: string;
  this?: string | Record<string, any>;
  confirmed_by?: string;
  when?: string;
  status?: "valid" | "ghost" | "pending";
  template_id?: string;
  flow_id?: string;
  consequence?: string;
}