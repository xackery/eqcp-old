import { createContext } from 'react';

const HostContext = createContext({})

export const HostProvider = HostContext.Provider;
export const HostConsumer = HostContext.Consumer;
export default HostContext;