declare global {
  interface Window {
    type: boolean;
    _port_: string | undefined;
  }
}

export {};
