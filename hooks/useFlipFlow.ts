import { useState } from "react";
import { LogLine } from "../types/logline";

export function useFlipFlow() {
  const [step, setStep] = useState(0);
  const [data, setData] = useState<Partial<LogLine>>({});

  function nextStep(newData: Partial<LogLine>) {
    setData((prev) => ({ ...prev, ...newData }));
    setStep((s) => s + 1);
  }
  function reset() {
    setStep(0);
    setData({});
  }

  return { step, data, nextStep, reset };
}