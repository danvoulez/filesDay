/// <reference types="@sveltejs/kit" />
import type { LogLine } from '$lib/logline/types';

declare global {
  namespace App {
    interface Locals {
      // Adicione suas propriedades personalizadas para Locals aqui
    }
  }
}