
namespace ZigBee {
    export interface IType {
        type: string;
        size: number;
        format?: (t: IType, val: ValueType) => string;
        bits?: {[bit: number]: string};
        values?: {[val: number]: string};
    }

    export interface IVal extends IType {
        value?: ValueType
        scale?: number
        unit?: IUnit
        numberValue?: () => number|undefined
        valueName?: () => string|undefined
    }
    export type IValInit<T,V> = (value?: V) => T;

    export interface IArgument extends IType {
        value?: any
        name: string
        description?: string
        mnf?: number
        arrayType?: () => IType
        payload?: {[name: string]: IValInit<IArgument,ValueType>}
        unit?: IUnit
        scale?: number
        bits?: {[bit: number]: string}
        values?: {[value: number]: string}
        cond?: ICondition[]
        [n: string]: any
    }
    export interface IAttribute extends IArgument {
        id: number
        report?: boolean
        read?: boolean
        write?: boolean
        require?: boolean

    }
    export interface ICommand extends IType {
        value: {[name:string]: any}
        id: number
        name: string
        description?: string
        payload: {[name: string]: IValInit<IArgument,any>}
    }

    export interface ICondition {
        field: string
        description?: string
        value: number
        mask?: number
        invert?: boolean
    }

    export interface ICluster {
        ID: number
        Name: string
        Desc: string
        Server: {
            Attribute: { [id: number]: IValInit<IAttribute,any> }
            Command: { [id: number]: IValInit<ICommand,any> }
        }
        Client: {
            Attribute: { [id: number]: IValInit<IAttribute,any> }
            Command: { [id: number]: IValInit<ICommand,any> }
        }
        [cmdName: string]: any
    }


    const command: IType = {
        type: "command",
        size: -1
    };

    export interface IUnit {
        unit: string
        format(v: ValueType): string
    }
    export type ValueType = string|object|number|Uint8Array;

    const lp = (n: number|string, z: number) => {
        if (typeof n === 'number') n = n.toString();
        while (n.length < z) n = `0${n}`;
        return n;
    };

    export function makeType<T extends IType,V>(...parts:any[]): IValInit<T,V> {
        return (value) => {
            let proto: any = {};
            for (let obj of parts) {
                proto = {
                    ...proto,
                    ...(typeof obj === 'function' ? obj() : obj)
                };
            }

            class Value {
                value: V;

                constructor(value?: V) {
                    for (let k in proto) {
                        if (!proto.hasOwnProperty(k)) continue;
                        this[k] = proto[k];
                    }
                    if (typeof this['parse'] === 'function') {
                        let nv = this['parse'](this, value);
                        // Ignore NaN
                        if (typeof nv !== 'number' || !isNaN(nv)) {
                            value = nv;
                        }
                    }

                    this.value = value;
                }

                valueName() {
                    if (typeof o.value === 'number' && typeof o.values === 'object' && typeof o.values[o.value] !== 'undefined') {
                        return `${o.values[o.value]}`;
                    }
                    return undefined;
                }

                numberValue() {
                    let val = o.value;
                    switch (typeof val) {
                        case 'string': return parseFloat(val);
                        case 'number': return val;
                        case 'object':
                            if (Array.isArray(val) && val.length === 1 && typeof val[0] === 'number') {
                                return val[0];
                            }
                            if (Array.isArray(val) && val.length === 1 && typeof val[0] === 'string') {
                                return parseFloat(val[0]);
                            }
                            break;
                    }
                }

                toString() {
                    let val = o.value;
                    let nv = o.numberValue();
                    let vn = o.valueName();

                    if (vn) {
                        return vn;
                    }
                    if (nv && typeof o.scale !== 'undefined') {
                        val = nv = (Math.round(100 * nv / o.scale) / 100);
                    }
                    if (typeof o.format === 'function') {
                        return o.format(o, val);
                    }
                    if (nv &&typeof o.unit !== 'undefined') {
                        val = o.unit.format(nv);
                    }
                    return "" + val;
                }
            }

            let o: IVal = <IVal><unknown>(new Value(value));
            return <T>o;
        }
    }

    export const native = {
        bool: { native: "bool" },
        array: { native: "array" },
        struct: { native: "struct" },
        int: { native: "int" },
        uint: { native: "uint" },
        bitmap: {
            native: "bitmap",
            format: (a: IType, v) => (Array.isArray(v) ? v.map(v => {
                    if (typeof a.bits !== 'undefined' && typeof a.bits[v] !== 'undefined') {
                        return a.bits[v]
                    }
                    return `${v}`
                }).join(", ") : ''),
        },
        byte: { native: "byte",
            parse: (a: IType, v: ValueType) => {
                if (typeof v !== 'string') { return v; }
                try {
                    return atob(v);
                } catch(e) {}
                return v;
            }
        },
        float: { native: "float" },
        string: { native: "string" },
        enum: { native: "enum" },
    };

    export const base = {
        local: makeType<IType,Object>(native.struct, {size: -1, type: "local"}),
        array: makeType<IType,any[]>(native.array, {size: -1, type: "array"}),
        struct: makeType<IType,Object>(native.struct, {size: -1, type: "struct"}),
        set: makeType<IType,any[]>(native.array, {size: -1, type: "set"}),
        bag: makeType<IType,any[]>(native.array, {size: -1, type: "bag"}),
        list: makeType<IType,any[]>(native.array, {size: -1, type: "list"}),
        bmp8: makeType<IType,number[]>(native.bitmap, {size: 1, type: "bmp8"}),
        bmp16: makeType<IType,number[]>(native.bitmap, {size: 2, type: "bmp16"}),
        bmp24: makeType<IType,number[]>(native.bitmap, {size: 3, type: "bmp24"}),
        bmp32: makeType<IType,number[]>(native.bitmap, {size: 4, type: "bmp32"}),
        bmp40: makeType<IType,number[]>(native.bitmap, {size: 5, type: "bmp40"}),
        bmp48: makeType<IType,number[]>(native.bitmap, {size: 6, type: "bmp48"}),
        bmp56: makeType<IType,number[]>(native.bitmap, {size: 7, type: "bmp56"}),
        bmp64: makeType<IType,number[]>(native.bitmap, {size: 8, type: "bmp64"}),
        dat8: makeType<IType,string>(native.byte, {size: 1, type: "dat8"}),
        dat16: makeType<IType,string>(native.byte, {size: 2, type: "dat16"}),
        dat24: makeType<IType,string>(native.byte, {size: 3, type: "dat24"}),
        dat32: makeType<IType,string>(native.byte, {size: 4, type: "dat32"}),
        dat40: makeType<IType,string>(native.byte, {size: 5, type: "dat40"}),
        dat48: makeType<IType,string>(native.byte, {size: 6, type: "dat48"}),
        dat56: makeType<IType,string>(native.byte, {size: 7, type: "dat56"}),
        dat64: makeType<IType,string>(native.byte, {size: 8, type: "dat64"}),
        bytes: makeType<IType,string>(native.byte, {size: -1, type: "bytes"}),
        enum8: makeType<IType,number>(native.enum, {size: 1, type: "enum8"}),
        enum16: makeType<IType,number>(native.enum, {size: 2, type: "enum16"}),
        bool: makeType<IType,number>(native.bool, {size: 1, type: "bool"}),
        cid: makeType<IType,number>(native.uint, {size: 2, type: "cid"}),
        aid: makeType<IType,number>(native.uint, {size: 2, type: "aid"}),
        oid: makeType<IType,number>(native.uint, {size: 4, type: "oid"}),
        uid: makeType<IType,number>(native.string, {size: 8, type: "uid"}),
        seckey: makeType<IType,string>(native.byte, {size: 16, type: "seckey"}),
        u8: makeType<IType,number>(native.uint, {size: 1, type: "u8"}),
        u16: makeType<IType,number>(native.uint, {size: 2, type: "u16"}),
        u24: makeType<IType,number>(native.uint, {size: 4, type: "u24"}),
        u32: makeType<IType,number>(native.uint, {size: 4, type: "u32"}),
        u40: makeType<IType,number>(native.uint, {size: 8, type: "u40"}),
        u48: makeType<IType,number>(native.uint, {size: 8, type: "u48"}),
        u56: makeType<IType,number>(native.uint, {size: 8, type: "u56"}),
        u64: makeType<IType,number>(native.uint, {size: 8, type: "u64"}),
        s8: makeType<IType,number>(native.int, {size: 1, type: "s8"}),
        s16: makeType<IType,number>(native.int, {size: 2, type: "s16"}),
        s24: makeType<IType,number>(native.int, {size: 4, type: "s24"}),
        s32: makeType<IType,number>(native.int, {size: 4, type: "s32"}),
        s40: makeType<IType,number>(native.int, {size: 8, type: "s40"}),
        s48: makeType<IType,number>(native.int, {size: 8, type: "s48"}),
        s56: makeType<IType,number>(native.int, {size: 8, type: "s56"}),
        s64: makeType<IType,number>(native.int, {size: 8, type: "s64"}),
        semi: makeType<IType,number>(native.float, {size: 4, type: "semi"}),
        float: makeType<IType,number>(native.float, {size: 4, type: "float"}),
        double: makeType<IType,number>(native.float, {size: 8, type: "double"}),
        ostring: makeType<IType,string>(native.byte, {size: -1, type: "ostring"}),
        cstring: makeType<IType,string>(native.string, {size: -1, type: "cstring"}),
        lostring: makeType<IType,string>(native.byte, {size: -2, type: "lostring"}),
        lcstring: makeType<IType,string>(native.string, {size: -2, type: "lcstring"}),
        date: makeType<IType,number>(native.string, {size: 4, type: "date",
            format: (v) => {
                if (v === null) return null;
                if (typeof v === 'string') v = parseInt(v);
                if (typeof v !== 'number') return v;
                // Can we stop using tiny numbers for dates? Y2K, Y2038 and now 2157...
                let yr = ((v & 0xFF000000) >> 24) + 1900
                    ,mo = ((v & 0x00FF0000) >> 16) - 1
                    ,day = (v & 0x0000FF00) >> 8
                    //,wd = (v & 0x000000FF);
                return (new Date(yr, mo, day, 0, 0, 0)).toLocaleDateString();
            }
        }),
        time: makeType<IType,number>(native.uint, {size: 4, type: "time",
            format: (v) => {
                if (v === null) return null;
                if (typeof v === 'string') v = parseInt(v);
                if (typeof v !== 'number') return v;
                let h = (v & 0xFF000000) >> 24,
                    m = (v & 0x00FF0000) >> 16,
                    s = (v & 0x0000FF00) >> 8,
                    c = (v & 0x000000FF);
                return `${lp(h,2)}:${lp(m,2)}:${lp(s,2)}.${lp(c,2)}`
            }
        }),
        utc: makeType<IType,number>(native.uint, {size: 4, type: "utc",
            format: (v) => {
                if (v === null) return null;
                if (typeof v === 'string') v = parseInt(v);
                if (typeof v !== 'number') return v;
                // Unixtime-like, but starts on 2000-01-01, should last until ~2137
                return (new Date(1000 * (v + 946684800))).toLocaleString();
            }
        }),
        Status: makeType<IType,number>(native.enum, { size: 1, type: "status",
            values: {
                0x00: "Success",
                0x01: "Failure",
                0x7e: "Not Authorized",
                0x7f: "Reserved Field Not Zero",
                0x80: "Malformed Command",
                0x81: "Unsupported Cluster Command",
                0x82: "Unsupported General Command",
                0x83: "Unsupported Manuf. Cluster Command",
                0x84: "Unsupported Manuf. General Command",
                0x85: "Invalid Field",
                0x86: "Unsupported Attribute",
                0x87: "Invalid Value",
                0x88: "Read Only",
                0x89: "Insufficient Space",
                0x8a: "Duplicate Exists",
                0x8b: "Not Found",
                0x8c: "Unreportable Attribute",
                0x8d: "Invalid Data Type",
                0x8e: "Invalid Sector",
                0x8f: "Write Only",
                0x90: "Inconsistent Startup State",
                0x91: "Defined Out Of Band",
                0xc0: "Hardware Failure",
                0xc1: "Software Failure",
                0xc2: "Calibration Error",
            } 
        }),
        SceneExtensionSet: makeType<IType,Object>(native.struct, { type: "sceneExtensionSet" }),
        EngineeringUnit: makeType<IType,number>(native.uint, {size: 2, type: "engineeringUnit",
            parse: (v) => {

                return v;
            }
        }),
    };

    export const ZDP = { 
        Types: {
            
            UnknownU8: makeType<ZigBee.IZDP.IArgUnknownU8, ZigBee.IZDP.IArgUnknownU8Payload>(base.u8, ()=>({
                name: `Unknown u8`,
                description: ``,
                
            })),
            ActiveEndpointList: makeType<ZigBee.IZDP.IArgActiveEndpointList, ZigBee.IZDP.IArgActiveEndpointListPayload>(base.set, ()=>({
                name: `Active Endpoint List`,
                description: `List of active endpoints`,
                arrayType: base.u8,
                
            })),
            ActiveEndpointSize: makeType<ZigBee.IZDP.IArgActiveEndpointSize, ZigBee.IZDP.IArgActiveEndpointSizePayload>(base.u8, ()=>({
                name: `Active Endpoint Size`,
                description: `Size in bytes of the Active Endpoints List`,
                
            })),
            RequestType: makeType<ZigBee.IZDP.IArgRequestType, ZigBee.IZDP.IArgRequestTypePayload>(base.enum8, ()=>({
                name: `Request Type`,
                description: `should be set to 1 if extended response is requested, 0 otherwise`,
                values: { 
                0x00: `Single Device Response`, 
                0x01: `Extended Response`,  },
                
            })),
            AddressMode: makeType<ZigBee.IZDP.IArgAddressMode, ZigBee.IZDP.IArgAddressModePayload>(base.enum8, ()=>({
                name: `Address Mode`,
                description: ``,
                values: { 
                0x01: `Group`, 
                0x02: `NWK`, 
                0x03: `IEEE`,  },
                
            })),
            ApsFlags: makeType<ZigBee.IZDP.IArgApsFlags, ZigBee.IZDP.IArgApsFlagsPayload>(base.u8, ()=>({
                name: `APS Flags`,
                description: ``,
                
            })),
            AssociatedDevices: makeType<ZigBee.IZDP.IArgAssociatedDevices, ZigBee.IZDP.IArgAssociatedDevicesPayload>(base.list, ()=>({
                name: `Associated Devices`,
                description: ``,
                arrayType: base.u16,
                
            })),
            BindingTable: makeType<ZigBee.IZDP.IArgBindingTable, ZigBee.IZDP.IArgBindingTablePayload>(base.array, ()=>({
                name: `Binding Table`,
                description: ``,
                arrayType: ZigBee.ZDP.Types.BindingTableEntry,
                
            })),
            BindingTableEntry: makeType<ZigBee.IZDP.IArgBindingTableEntry, ZigBee.IZDP.IArgBindingTableEntryPayload>(base.local, ()=>({
                name: `Binding Table Entry`,
                description: ``,
                payload: { 
                    SourceAddress: ZigBee.ZDP.Types.SourceAddress,
                    SourceEndpoint: ZigBee.ZDP.Types.SourceEndpoint,
                    ClusterId: ZigBee.ZDP.Types.ClusterId,
                    AddressMode: ZigBee.ZDP.Types.AddressMode,
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    IeeeAddress: ZigBee.ZDP.Types.IeeeAddress,
                    Endpoint: ZigBee.ZDP.Types.Endpoint,
                 },
                
            })),
            BindingTarget: makeType<ZigBee.IZDP.IArgBindingTarget, ZigBee.IZDP.IArgBindingTargetPayload>(base.u16, ()=>({
                name: `Binding Target`,
                description: `NWK Address`,
                
            })),
            ClusterId: makeType<ZigBee.IZDP.IArgClusterId, ZigBee.IZDP.IArgClusterIdPayload>(base.u16, ()=>({
                name: `Cluster ID`,
                description: ``,
                
            })),
            ComplexDescriptor: makeType<ZigBee.IZDP.IArgComplexDescriptor, ZigBee.IZDP.IArgComplexDescriptorPayload>(base.ostring, ()=>({
                name: `Complex Descriptor`,
                description: ``,
                
            })),
            ComplexDescriptorAvailable: makeType<ZigBee.IZDP.IArgComplexDescriptorAvailable, ZigBee.IZDP.IArgComplexDescriptorAvailablePayload>(base.bool, ()=>({
                name: `Complex Descriptor Available`,
                description: ``,
                
            })),
            Depth: makeType<ZigBee.IZDP.IArgDepth, ZigBee.IZDP.IArgDepthPayload>(base.u8, ()=>({
                name: `Depth`,
                description: `of the neighbor device. A value of 0 indicates that the device is the coordinator for the network`,
                
            })),
            DescriptorCapability: makeType<ZigBee.IZDP.IArgDescriptorCapability, ZigBee.IZDP.IArgDescriptorCapabilityPayload>(base.bmp8, ()=>({
                name: `Descriptor Capability`,
                description: ``,
                bits: { 
                0: `Extended Active Endpoint List Available`, 
                1: `Extended Simple Descriptor List Available`,  },
                
            })),
            Capability: makeType<ZigBee.IZDP.IArgCapability, ZigBee.IZDP.IArgCapabilityPayload>(base.bmp8, ()=>({
                name: `Capability`,
                description: `specifies the device:s capabilities`,
                bits: { 
                0: `Pan Coordinator`, 
                1: `Full Function`, 
                2: `Mains Power`, 
                3: `RX on Idle`, 
                6: `Security`, 
                7: `Allocate Address`,  },
                
            })),
            DeviceType: makeType<ZigBee.IZDP.IArgDeviceType, ZigBee.IZDP.IArgDeviceTypePayload>(base.enum16, ()=>({
                name: `Device Type`,
                description: ``,
                values: { 
                0x0000: `On/Off Switch`, 
                0x0001: `Level Control Switch`, 
                0x0002: `On/Off Output`, 
                0x0003: `Level Controllable Output`, 
                0x0004: `Scene Selector`, 
                0x0005: `Configuration Tool`, 
                0x0006: `Remote Control`, 
                0x0007: `Combined Interface`, 
                0x0008: `Range Extender`, 
                0x0009: `Mains Power Outlet`, 
                0x000a: `Door lock`, 
                0x000c: `Simple sensor`, 
                0x0010: `On/off plug-in unit`, 
                0x0051: `Smart plug`, 
                0x0060: `GP Proxy`, 
                0x0061: `GP Proxy Basic`, 
                0x0062: `GP Target Plus`, 
                0x0063: `GP Target`, 
                0x0064: `GP Commissioning Tool`, 
                0x0065: `GP Combo`, 
                0x0066: `GP Combo Basic`, 
                0x0100: `On/Off Light`, 
                0x0101: `Dimmable Light`, 
                0x0102: `Color Dimmable Light`, 
                0x0103: `On/Off Light Switch`, 
                0x0104: `Dimmer Switch`, 
                0x0105: `Color Dimmer Switch`, 
                0x0106: `Light Sensor`, 
                0x0107: `Occupancy Sensor`, 
                0x0108: `On/off ballast`, 
                0x0109: `Dimmable ballast`, 
                0x010a: `On/off plugin unit`, 
                0x010b: `Dimmable plugin unit`, 
                0x010c: `Color temperature light`, 
                0x010d: `Extended color light`, 
                0x010e: `Light level sensor`, 
                0x0110: `Dimmable plug-in unit`, 
                0x0200: `Shade`, 
                0x0201: `Shade Controller`, 
                0x0202: `Window Covering Device`, 
                0x0210: `Extended color light 2`, 
                0x0220: `Color temperature light 2`, 
                0x0300: `Heating/Cooling Unit`, 
                0x0301: `Thermostat`, 
                0x0302: `Temperature Sensor`, 
                0x0303: `Pump`, 
                0x0304: `Pump Controller`, 
                0x0305: `Pressure Sensor`, 
                0x0306: `Flow Sensor`, 
                0x0400: `IAS Control and Indicating Equipment`, 
                0x0401: `IAS Ancillary Control Equipment`, 
                0x0402: `IAS Zone`, 
                0x0403: `IAS Warning Device`, 
                0x0500: `Energy Service Portal`, 
                0x0501: `Metering Device`, 
                0x0502: `In-Premise Display`, 
                0x0503: `Programmable Communicating Thermostat (PCT)`, 
                0x0504: `Load Control Device`, 
                0x0505: `Smart Appliance`, 
                0x0506: `Prepayment Terminal`, 
                0x0507: `Device Management`, 
                0x0800: `Color controller`, 
                0x0810: `Color scene controller`, 
                0x0820: `Non color controller`, 
                0x0830: `Non color scene controller`, 
                0x0840: `Control bridge`, 
                0x0850: `On/off sensor`,  },
                
            })),
            DeviceVersion: makeType<ZigBee.IZDP.IArgDeviceVersion, ZigBee.IZDP.IArgDeviceVersionPayload>(base.u8, ()=>({
                name: `Device Version`,
                description: `is dependant on profile`,
                
            })),
            Endpoint: makeType<ZigBee.IZDP.IArgEndpoint, ZigBee.IZDP.IArgEndpointPayload>(base.u8, ()=>({
                name: `Endpoint`,
                description: ``,
                
            })),
            EndpointList: makeType<ZigBee.IZDP.IArgEndpointList, ZigBee.IZDP.IArgEndpointListPayload>(base.set, ()=>({
                name: `Endpoint List`,
                description: ``,
                arrayType: base.u8,
                
            })),
            EnergyValues: makeType<ZigBee.IZDP.IArgEnergyValues, ZigBee.IZDP.IArgEnergyValuesPayload>(base.array, ()=>({
                name: `Energy Values`,
                description: ``,
                arrayType: base.u8,
                
            })),
            FrequencyBands: makeType<ZigBee.IZDP.IArgFrequencyBands, ZigBee.IZDP.IArgFrequencyBandsPayload>(base.bmp8, ()=>({
                name: `Frequency Bands`,
                description: ``,
                bits: { 
                0: `868 MHz`, 
                2: `902-928 MHz`, 
                3: `2.4 GHz`, 
                4: `EU Sub-GHz`,  },
                
            })),
            IeeeAddress: makeType<ZigBee.IZDP.IArgIeeeAddress, ZigBee.IZDP.IArgIeeeAddressPayload>(base.uid, ()=>({
                name: `IEEE Address`,
                description: `is a 64-bit MAC address`,
                
            })),
            InClusterList: makeType<ZigBee.IZDP.IArgInClusterList, ZigBee.IZDP.IArgInClusterListPayload>(base.set, ()=>({
                name: `In Cluster List`,
                description: `is a list of input clusters`,
                arrayType: base.u16,
                
            })),
            LogicalType: makeType<ZigBee.IZDP.IArgLogicalType, ZigBee.IZDP.IArgLogicalTypePayload>(base.enum8, ()=>({
                name: `Logical Type`,
                description: ``,
                values: { 
                0x00: `Coordinator`, 
                0x01: `Router`, 
                0x02: `End Device`,  },
                
            })),
            Lqi: makeType<ZigBee.IZDP.IArgLqi, ZigBee.IZDP.IArgLqiPayload>(base.u8, ()=>({
                name: `Lqi`,
                description: `is the estimated link quality for RF transmissions from this device`,
                unit: units.Percent,
                scale: 2.55,
                
            })),
            MacCapabilityFlags: makeType<ZigBee.IZDP.IArgMacCapabilityFlags, ZigBee.IZDP.IArgMacCapabilityFlagsPayload>(base.bmp8, ()=>({
                name: `MAC Capability Flags`,
                description: ``,
                bits: { 
                0: `Alternate PAN Coordinator`, 
                1: `Full function device`, 
                2: `Mains powered`, 
                3: `Receiver on when idle`, 
                6: `Supports secured frames`, 
                7: `Allocate Address`,  },
                
            })),
            MaxBufferSize: makeType<ZigBee.IZDP.IArgMaxBufferSize, ZigBee.IZDP.IArgMaxBufferSizePayload>(base.u8, ()=>({
                name: `Max Buffer Size`,
                description: `specifies the maximum size, in octets, of the network sub-layer data unit (NSDU) for this node.
This is the maximum size of data or commands passed to or from the application by the application support
sub-layer, before any fragmentation or re-assembly.`,
                
            })),
            MaxRxSize: makeType<ZigBee.IZDP.IArgMaxRxSize, ZigBee.IZDP.IArgMaxRxSizePayload>(base.u16, ()=>({
                name: `Max RX size`,
                description: `specifies the maximum size, in octets, of the application sub-layer data unit (ASDU) that can be transferred to this node in one single message transfer.`,
                
            })),
            MaxTxSize: makeType<ZigBee.IZDP.IArgMaxTxSize, ZigBee.IZDP.IArgMaxTxSizePayload>(base.u16, ()=>({
                name: `Max TX size`,
                description: `specifies the maximum size, in octets, of the application sub-layer data unit (ASDU) that can be transferred from
this node in one single message transfer.`,
                
            })),
            ManufacturerCode: makeType<ZigBee.IZDP.IArgManufacturerCode, ZigBee.IZDP.IArgManufacturerCodePayload>(base.enum16, ()=>({
                name: `Manufacturer Code`,
                description: ``,
                values: { 
                0x1000: `Cirronet`, 
                0x1001: `Chipcon`, 
                0x1002: `Ember`, 
                0x1003: `National Tech`, 
                0x1004: `Freescale`, 
                0x1005: `IPCom`, 
                0x1006: `San Juan Software`, 
                0x1007: `TUV`, 
                0x1008: `CompXs`, 
                0x1009: `BM SpA`, 
                0x100a: `AwarePoint`, 
                0x100b: `Philips`, 
                0x100c: `Luxoft`, 
                0x100d: `Korvin`, 
                0x100e: `One RF`, 
                0x100f: `Software Technology Group`, 
                0x1010: `Telegesis`, 
                0x1011: `Visionic`, 
                0x1012: `Insta`, 
                0x1013: `Atalum`, 
                0x1014: `Atmel`, 
                0x1015: `Develco`, 
                0x1016: `Honeywell`, 
                0x1017: `RadioPulse`, 
                0x1018: `Renesas`, 
                0x1019: `Xanadu Wireless`, 
                0x101a: `NEC Engineering`, 
                0x101b: `Yamatake`, 
                0x101c: `Tendril`, 
                0x101d: `Assa Abloy`, 
                0x101e: `Maxstream`, 
                0x101f: `Neurocom`, 
                0x1020: `Institute for Information Industry`, 
                0x1021: `Vantage Controls`, 
                0x1022: `iControl`, 
                0x1023: `Raymarine`, 
                0x1024: `LS Research`, 
                0x1025: `Onity`, 
                0x1026: `Mono Products`, 
                0x1027: `RF Tech`, 
                0x1028: `Itron`, 
                0x1029: `Tritech`, 
                0x102a: `Embedit`, 
                0x102b: `S3C`, 
                0x102c: `Siemens`, 
                0x102d: `Mindtech`, 
                0x102e: `LG Electronics`, 
                0x102f: `Mitsubishi`, 
                0x1030: `Johnson Controls`, 
                0x1031: `PRI`, 
                0x1032: `Knick`, 
                0x1033: `Viconics`, 
                0x1034: `Flexipanel`, 
                0x1035: `Piasim Corporation`, 
                0x1036: `Trane`, 
                0x1037: `Jennic`, 
                0x1038: `Living Independently`, 
                0x1039: `AlertMe`, 
                0x103a: `Daintree`, 
                0x103b: `Aiji`, 
                0x103c: `Telecom Italia`, 
                0x103d: `Mikrokrets`, 
                0x103e: `Oki Semi`, 
                0x103f: `Newport Electronics`, 
                0x1040: `Control4`, 
                0x1041: `STMicro`, 
                0x1042: `Ad-Sol Nissin`, 
                0x1043: `DCSI`, 
                0x1044: `France Telecom`, 
                0x1045: `muNet`, 
                0x1046: `Autani`, 
                0x1047: `Colorado vNet`, 
                0x1048: `Aerocomm`, 
                0x1049: `Silicon Labs`, 
                0x104F: `Crane`, 
                0x104a: `Inncom`, 
                0x104b: `Cannon`, 
                0x104c: `Synapse`, 
                0x104d: `Fisher Pierce/Sunrise`, 
                0x104e: `CentraLite`, 
                0x1050: `Mobilarm`, 
                0x1051: `iMonitor`, 
                0x1052: `Bartech`, 
                0x1053: `Meshnetics`, 
                0x1054: `LS Industrial`, 
                0x1055: `Cason`, 
                0x1056: `Wireless Glue`, 
                0x1057: `Elster`, 
                0x1058: `SMS Tec`, 
                0x1059: `Onset Computer`, 
                0x105a: `Riga Development`, 
                0x105b: `Energate`, 
                0x105c: `ConMed Linvatec`, 
                0x105d: `PowerMand`, 
                0x105e: `Schneider Electric`, 
                0x105f: `Eaton`, 
                0x1060: `Telular`, 
                0x1061: `Delphi Medical`, 
                0x1062: `EpiSensor`, 
                0x1063: `Landis+Gyr`, 
                0x1064: `Kaba Group`, 
                0x1065: `Shure`, 
                0x1066: `Comverge`, 
                0x1067: `DBS Lodging`, 
                0x1068: `Energy Aware`, 
                0x1069: `Hidalgo`, 
                0x106a: `Air2App`, 
                0x106b: `AMX`, 
                0x106c: `EDMI Pty`, 
                0x106d: `Cyan Ltd`, 
                0x106e: `System SPA`, 
                0x106f: `Telit`, 
                0x1070: `Kaga Electronics`, 
                0x1071: `4-noks s.r.l.`, 
                0x1072: `Certicom`, 
                0x1073: `Gridpoint`, 
                0x1074: `Profile Systems`, 
                0x1075: `Compacta International`, 
                0x1076: `Freestyle Technology`, 
                0x1077: `Alektrona`, 
                0x1078: `Computime`, 
                0x1079: `Remote Technologies`, 
                0x107a: `Wavecom`, 
                0x107b: `Energy Optimizers`, 
                0x107c: `GE`, 
                0x107d: `Jetlun`, 
                0x107e: `Cipher Systems`, 
                0x107f: `Corporate Systems Eng`, 
                0x1080: `ecobee`, 
                0x1081: `SMK`, 
                0x1082: `Meshworks Wireless`, 
                0x1083: `Ellips B.V.`, 
                0x1084: `Secure electrans`, 
                0x1085: `CEDO`, 
                0x1086: `Toshiba`, 
                0x1087: `Digi International`, 
                0x1088: `Ubilogix`, 
                0x1089: `Echelon`, 
                0x1090: `Green Energy Options`, 
                0x1091: `Silver Spring Networks`, 
                0x1092: `Black & Decker`, 
                0x1093: `Aztech AssociatesInc.`, 
                0x1094: `A&D Co`, 
                0x1095: `Rainforest Automation`, 
                0x1096: `Carrier Electronics`, 
                0x1097: `SyChip/Murata`, 
                0x1098: `OpenPeak`, 
                0x1099: `Passive Systems`, 
                0x109a: `MMBResearch`, 
                0x109b: `Leviton`, 
                0x109c: `Korea Electric Power Data Network`, 
                0x109d: `Comcast`, 
                0x109e: `NEC Electronics`, 
                0x109f: `Netvox`, 
                0x10a0: `U-Control`, 
                0x10a1: `Embedia Technologies`, 
                0x10a2: `Sensus`, 
                0x10a3: `SunriseTechnologies`, 
                0x10a4: `MemtechCorp`, 
                0x10a5: `Freebox`, 
                0x10a6: `M2 Labs`, 
                0x10a7: `BritishGas`, 
                0x10a8: `Sentec`, 
                0x10a9: `Navetas`, 
                0x10aa: `Lightspeed Technologies`, 
                0x10ab: `Oki Electric`, 
                0x10ac: `Sistemas Inteligentes`, 
                0x10ad: `Dometic`, 
                0x10ae: `Alps`, 
                0x10af: `EnergyHub`, 
                0x10b0: `Kamstrup`, 
                0x10b1: `EchoStar`, 
                0x10b2: `EnerNOC`, 
                0x10b3: `Eltav`, 
                0x10b4: `Belkin`, 
                0x10b5: `XStreamHD Wireless`, 
                0x10b6: `Saturn South`, 
                0x10b7: `GreenTrapOnline`, 
                0x10b8: `SmartSynch`, 
                0x10b9: `Nyce Control`, 
                0x10ba: `ICM Controls`, 
                0x10bb: `Millennium Electronics`, 
                0x10bc: `Motorola`, 
                0x10bd: `EmersonWhite-Rodgers`, 
                0x10be: `Radio Thermostat`, 
                0x10bf: `OMRONCorporation`, 
                0x10c0: `GiiNii GlobalLimited`, 
                0x10c1: `Fujitsu GeneralLimited`, 
                0x10c2: `Peel Technologies`, 
                0x10c3: `Accent`, 
                0x10c4: `ByteSnap Design`, 
                0x10c5: `NEC TOKIN Corporation`, 
                0x10c6: `G4S JusticeServices`, 
                0x10c7: `Trilliant Networks`, 
                0x10c8: `Electrolux Italia`, 
                0x10c9: `OnzoLtd`, 
                0x10ca: `EnTekSystems`, 
                0x10cb: `Philips 2`, 
                0x10cc: `MainstreamEngineering`, 
                0x10cd: `IndesitCompany`, 
                0x10ce: `THINKECO`, 
                0x10cf: `2D2C`, 
                0x10d0: `GreenPeak`, 
                0x10d1: `InterCEL`, 
                0x10d2: `LG Electronics 2`, 
                0x10d3: `Mitsumi Electric`, 
                0x10d4: `Mitsumi Electric 2`, 
                0x10d5: `Zentrum Mikroelektronik Dresden`, 
                0x10d6: `Nest Labs`, 
                0x10d7: `Exegin Technologies`, 
                0x10d8: `Honeywell 2`, 
                0x10d9: `Takahata Precision`, 
                0x10da: `Sumitomo Electric Networks`, 
                0x10db: `GE Energy`, 
                0x10dc: `GE Appliances`, 
                0x10dd: `Radiocrafts AS`, 
                0x10de: `Ceiva`, 
                0x10df: `TEC CO Co., Ltd`, 
                0x10e0: `Chameleon Technology (UK) Ltd`, 
                0x10e1: `Samsung`, 
                0x10e2: `ruwido austria gmbh`, 
                0x10e3: `Huawei Technologies Co., Ltd.`, 
                0x10e4: `Huawei Technologies Co., Ltd. 2`, 
                0x10e5: `Greenwave Reality`, 
                0x10e6: `BGlobal Metering Ltd`, 
                0x10e7: `Mindteck`, 
                0x10e8: `Ingersoll-Rand`, 
                0x10e9: `Dius Computing Pty Ltd`, 
                0x10ea: `Embedded Automation, Inc.`, 
                0x10eb: `ABB`, 
                0x10ec: `Sony`, 
                0x10ed: `Genus Power Infrastructures Limited`, 
                0x10ee: `Universal Devices`, 
                0x10ef: `Universal Devices 2`, 
                0x10f0: `Metrum Technologies, LLC`, 
                0x10f1: `Cisco`, 
                0x10f2: `Ubisys technologies GmbH`, 
                0x10f3: `Consert`, 
                0x10f4: `Crestron Electronics`, 
                0x10f5: `Enphase Energy`, 
                0x10f6: `Invensys Controls`, 
                0x10f7: `Mueller Systems, LLC`, 
                0x10f8: `AAC Technologies Holding`, 
                0x10f9: `U-NEXT Co., Ltd`, 
                0x10fa: `Steelcase Inc.`, 
                0x10fb: `Telematics Wireless`, 
                0x10fc: `Samil Power Co., Ltd`, 
                0x10fd: `Pace Plc`, 
                0x10fe: `Osborne Coinage Co.`, 
                0x10ff: `Powerwatch`, 
                0x1100: `CANDELED GmbH`, 
                0x1101: `FlexGrid S.R.L`, 
                0x1102: `Humax`, 
                0x1103: `Universal Electronics, Inc.`, 
                0x1104: `Advanced Energy`, 
                0x1105: `BEGA Gantenbrink-Leuchten`, 
                0x1106: `Brunel University`, 
                0x1107: `Panasonic R&D Center Singapore`, 
                0x1108: `eSystems Research`, 
                0x1109: `Panamax`, 
                0x110a: `Physical Graph Corporation`, 
                0x110b: `EM-Lite Ltd.`, 
                0x110c: `Osram Sylvania`, 
                0x110d: `2 Save Energy Ltd.`, 
                0x110e: `Planet Innovation Products Pty Ltd`, 
                0x110f: `Ambient Devices, Inc.`, 
                0x1110: `Profalux`, 
                0x1111: `Billion Electric Company (BEC)`, 
                0x1112: `Embertec Pty Ltd`, 
                0x1113: `IT Watchdogs`, 
                0x1114: `Reloc`, 
                0x1115: `Intel Corporation`, 
                0x1116: `Trend Electronics Limited`, 
                0x1117: `Moxa`, 
                0x1118: `QEES`, 
                0x1119: `SAYME Wireless Sensor Networks`, 
                0x111a: `Pentair Aquatic Systems`, 
                0x111b: `Orbit Irrigation`, 
                0x111c: `California Eastern Laboratories`, 
                0x111d: `Comcast 2`, 
                0x111e: `IDT Technology Limited`, 
                0x111f: `Pixela`, 
                0x1120: `TiVo`, 
                0x1121: `Fidure`, 
                0x1122: `Marvell Semiconductor`, 
                0x1123: `Wasion Group`, 
                0x1124: `Jasco Products`, 
                0x1125: `Shenzhen Kaifa Technology`, 
                0x1126: `Netcomm Wireless`, 
                0x1127: `Define Instruments`, 
                0x1128: `In Home Displays`, 
                0x1129: `Miele & Cie. KG`, 
                0x112a: `Televes S.A.`, 
                0x112b: `Labelec`, 
                0x112c: `China Electronics Standardization Institute`, 
                0x112d: `Vectorform`, 
                0x112e: `Busch-Jaeger Elektro`, 
                0x112f: `Redpine Signals`, 
                0x1130: `Bridges Electronic Technology`, 
                0x1131: `Sercomm`, 
                0x1132: `WSH GmbH wirsindheller`, 
                0x1133: `Bosch Security Systems`, 
                0x1134: `eZEX Corporation`, 
                0x1135: `Dresden Elektronik Ingenieurtechnik GmbH`, 
                0x1136: `MEAZON S.A.`, 
                0x1137: `Crow Electronic Engineering`, 
                0x1138: `Harvard Engineering`, 
                0x1139: `Andson(Beijing) Technology`, 
                0x113a: `Adhoco AG`, 
                0x113b: `Waxman Consumer Products Group`, 
                0x113c: `Owon Technology`, 
                0x113d: `Hitron Technologies`, 
                0x113e: `Scemtec Steuerungstechnik GmbH`, 
                0x113f: `Webee`, 
                0x1140: `Grid2Home`, 
                0x1141: `Telink Micro`, 
                0x1142: `Jasmine Systems`, 
                0x1143: `Bidgely`, 
                0x1144: `Lutron`, 
                0x1145: `IJENKO`, 
                0x1146: `Starfield Electronic`, 
                0x1147: `TCP`, 
                0x1148: `Rogers Communications Partnership`, 
                0x1149: `Cree`, 
                0x114a: `Robert Bosch LLC`, 
                0x114b: `Ibis Networks`, 
                0x114c: `Quirky`, 
                0x114d: `Efergy Technologies`, 
                0x114e: `Smartlabs`, 
                0x114f: `Everspring Industry`, 
                0x1150: `Swann Communications`, 
                0x1151: `Soneter`, 
                0x1152: `Samsung SDS`, 
                0x1153: `Uniband Electronic Corporation`, 
                0x1154: `Accton Technology Corporation`, 
                0x1155: `Bosch Thermotechnik GmbH`, 
                0x1156: `Wincor Nixdorf Inc.`, 
                0x1157: `Ohsung Electronics`, 
                0x1158: `Zen Within, Inc.`, 
                0x1159: `Tech4home, Lda.`, 
                0x115A: `Nanoleaf`, 
                0x115B: `Keen Home, Inc.`, 
                0x115C: `Poly-Control APS`, 
                0x115D: `Eastfield Lighting Co., Ltd Shenzhen`, 
                0x115E: `IP Datatel, Inc.`, 
                0x115F: `Lumi United Techology, Ltd Shenzhen`, 
                0x1160: `Sengled Optoelectronics Corp`, 
                0x1161: `Remote Solution Co., Ltd.`, 
                0x1162: `ABB Genway Xiamen Electrical Equipment Co., Ltd.`, 
                0x1163: `Zhejiang Rexense Tech`, 
                0x1164: `ForEE Technology`, 
                0x1165: `Open Access Technology Intl.`, 
                0x1166: `INNR Lighting BV`, 
                0x1167: `Techworld Industries`, 
                0x1168: `Leedarson Lighting Co., Ltd.`, 
                0x1169: `Arzel Zoning`, 
                0x116A: `Holley Technology`, 
                0x116B: `Beldon Technologies`, 
                0x116C: `Flextronics`, 
                0x116D: `Shenzhen Meian`, 
                0x116E: `Lowes`, 
                0x116F: `Sigma Connectivity`, 
                0x1171: `Wulian`, 
                0x1172: `Plugwise B.V.`, 
                0x1173: `Titan Products`, 
                0x1174: `Ecospectral`, 
                0x1175: `D-Link`, 
                0x1176: `Technicolor Home USA`, 
                0x1177: `Opple Lighting`, 
                0x1178: `Wistron NeWeb Corp.`, 
                0x1179: `QMotion Shades`, 
                0x117A: `Insta Elektro GmbH`, 
                0x117B: `Shanghai Vancount`, 
                0x117C: `Ikea of Sweden`, 
                0x117D: `RT-RK`, 
                0x117E: `Shenzhen Feibit`, 
                0x117F: `EuControls`, 
                0x1180: `Telkonet`, 
                0x1181: `Thermal Solution Resources`, 
                0x1182: `PomCube`, 
                0x1183: `Ei Electronics`, 
                0x1184: `Optoga`, 
                0x1185: `Stelpro`, 
                0x1186: `Lynxus Technologies Corp.`, 
                0x1187: `Semiconductor Components`, 
                0x1188: `TP-Link`, 
                0x1189: `LEDVANCE LLC.`, 
                0x118A: `Nortek`, 
                0x118B: `iRevo/Assa Abbloy Korea`, 
                0x118C: `Midea`, 
                0x118D: `ZF Friedrichshafen`, 
                0x118E: `Checkit`, 
                0x118F: `Aclara`, 
                0x1190: `Nokia`, 
                0x1191: `Goldcard High-tech Co., Ltd.`, 
                0x1192: `George Wilson Industries Ltd.`, 
                0x1193: `EASY SAVER CO.,INC`, 
                0x1194: `ZTE Corporation`, 
                0x1195: `ARRIS`, 
                0x1196: `Reliance BIG TV`, 
                0x1197: `Insight Energy Ventures/Powerley`, 
                0x1198: `Thomas Research Products (Hubbell Lighting Inc.)`, 
                0x1199: `Li Seng Technology`, 
                0x119A: `System Level Solutions Inc.`, 
                0x119B: `Matrix Labs`, 
                0x119C: `Sinope Technologies`, 
                0x119D: `Jiuzhou Greeble`, 
                0x119E: `Guangzhou Lanvee Tech. Co. Ltd.`, 
                0x119F: `Venstar`, 
                0x1200: `SLV`, 
                0x1201: `Halo Smart Labs`, 
                0x1202: `Scout Security Inc.`, 
                0x1203: `Alibaba China Inc.`, 
                0x1204: `Resolution Products, Inc.`, 
                0x1205: `Smartlok Inc.`, 
                0x1206: `Lux Products Corp.`, 
                0x1207: `Vimar SpA`, 
                0x1208: `Universal Lighting Technologies`, 
                0x1209: `Robert Bosch, GmbH`, 
                0x120A: `Accenture`, 
                0x120B: `Heiman Technology Co., Ltd.`, 
                0x120C: `Shenzhen HOMA Technology Co., Ltd.`, 
                0x120D: `Vision-Electronics Technology`, 
                0x120E: `Lenovo`, 
                0x120F: `Presciense R&D`, 
                0x1210: `Shenzhen Seastar Intelligence Co., Ltd.`, 
                0x1211: `Sensative AB`, 
                0x1212: `SolarEdge`, 
                0x1213: `Zipato`, 
                0x1214: `China Fire & Security Sensing Manufacturing (iHorn)`, 
                0x1215: `Quby BV`, 
                0x1216: `Hangzhou Roombanker Technology Co., Ltd.`, 
                0x1217: `Amazon Lab126`, 
                0x1218: `Paulmann Licht GmbH`, 
                0x1219: `Shenzhen Orvibo Electronics Co. Ltd.`, 
                0x121A: `TCI Telecommunications`, 
                0x121B: `Mueller-Licht International Inc.`, 
                0x121C: `Aurora Limited`, 
                0x121D: `SmartDCC`, 
                0x121E: `Shanghai UMEinfo Co. Ltd.`, 
                0x121F: `carbonTRACK`, 
                0x1220: `Somfy`, 
                0x1221: `Viessmann Elektronik GmbH`, 
                0x1222: `Hildebrand Technology Ltd`, 
                0x1223: `Onkyo Technology Corporation`, 
                0x1224: `Shenzhen Sunricher Technology Ltd.`, 
                0x1225: `Xiu Xiu Technology Co., Ltd`, 
                0x1226: `Zumtobel Group`, 
                0x1227: `Shenzhen Kaadas Intelligent Technology Co. Ltd`, 
                0x1228: `Shanghai Xiaoyan Technology Co. Ltd`, 
                0x1229: `Cypress Semiconductor`, 
                0x122A: `XAL GmbH`, 
                0x122B: `Inergy Systems LLC`, 
                0x122C: `Alfred Karcher GmbH & Co KG`, 
                0x122D: `Adurolight Manufacturing`, 
                0x122E: `Groupe Muller`, 
                0x122F: `V-Mark Enterprises Inc.`, 
                0x1230: `Lead Energy AG`, 
                0x1231: `UIOT Group`, 
                0x1232: `Axxess Industries Inc.`, 
                0x1233: `Third Reality Inc.`, 
                0x1234: `DSR Corporation`, 
                0x1235: `Guangzhou Vensi Intelligent Technology Co. Ltd.`, 
                0x1236: `Schlage Lock (Allegion)`, 
                0x1237: `Net2Grid`, 
                0x1238: `Airam Electric Oy Ab`, 
                0x1239: `IMMAX WPB CZ`, 
                0x123A: `ZIV Automation`, 
                0x123B: `HangZhou iMagicTechnology Co., Ltd`, 
                0x123C: `Xiamen Leelen Technology Co. Ltd.`, 
                0x123D: `Overkiz SAS`, 
                0x123E: `Flonidan A/S`, 
                0x123F: `HDL Automation Co., Ltd.`, 
                0x1240: `Ardomus Networks Corporation`, 
                0x1241: `Samjin Co., Ltd.`, 
                0x1242: `Sprue Aegis PLC`, 
                0x1243: `Indra Sistemas, S.A.`, 
                0x1244: `Shenzhen JBT Smart Lighting Co., Ltd.`, 
                0x1245: `GE Lighting & Current`, 
                0x1246: `Danfoss A/S`, 
                0x1247: `NIVISS PHP Sp. z o.o. Sp.k.`, 
                0x1248: `Shenzhen Fengliyuan Energy Conservating Technology Co. Ltd`, 
                0x1249: `NEXELEC`, 
                0x124A: `Sichuan Behome Prominent Technology Co., Ltd`, 
                0x124B: `Fujian Star-net Communication Co., Ltd.`, 
                0x124C: `Toshiba Visual Solutions Corporation`, 
                0x124D: `Latchable, Inc.`, 
                0x124E: `L&S Deutschland GmbH`, 
                0x124F: `Gledopto Co., Ltd.`, 
                0x1250: `The Home Depot`, 
                0x1251: `Neonlite International Ltd.`, 
                0x1252: `Arlo Technologies, Inc.`, 
                0x1253: `Xingluo Technology Co., Ltd.`, 
                0x1254: `Simon Electric (China) Co., Ltd.`, 
                0x1255: `Hangzhou Greatstar Industrial Co., Ltd.`, 
                0x1256: `Sequentric Energy Systems, LLC`, 
                0x1257: `Solum Co., Ltd.`, 
                0x1258: `Eaglerise Electric & Electronic (China) Co., Ltd.`, 
                0x1259: `Fantem Technologies (Shenzhen) Co., Ltd.`, 
                0x125A: `Yunding Network Technology (Beijing) Co., Ltd.`, 
                0x125B: `Atlantic Group`, 
                0x125C: `Xiamen Intretech, Inc.`, 
                0x125D: `Tuya Global Inc.`, 
                0x125E: `Xiamen Dnake Intelligent Technology Co., Ltd`, 
                0x125F: `Niko nv`, 
                0x1260: `Emporia Energy`, 
                0x1261: `Sikom AS`, 
                0x1262: `AXIS Labs, Inc.`, 
                0x1263: `Current Products Corporation`, 
                0x1264: `MeteRSit SRL`, 
                0x1265: `HORNBACH Baumarkt AG`, 
                0x1266: `DiCEworld s.r.l. a socio unico`, 
                0x1267: `ARC Technology Co., Ltd`, 
                0x1268: `Hangzhou Konke Information Technology Co., Ltd.`, 
                0x1269: `SALTO Systems S.L.`, 
                0x126A: `Shenzhen Shyugj Technology Co., Ltd`, 
                0x126B: `Brayden Automation Corporation`, 
                0x126C: `Environexus Pty. Ltd.`, 
                0x126D: `Eltra nv/sa`, 
                0x126E: `Xiaomi Communications Co., Ltd.`, 
                0x126F: `Shanghai Shuncom Electronic Technology Co., Ltd.`, 
                0x1270: `Voltalis S.A`, 
                0x1271: `FEELUX Co., Ltd.`, 
                0x1272: `SmartPlus Inc.`, 
                0x1273: `Halemeier GmbH`, 
                0x1274: `Trust International BBV`, 
                0x1275: `Duke Energy Business Services LLC`, 
                0x1276: `Calix, Inc.`, 
                0x1277: `ADEO`, 
                0x1278: `Connected Response Limited`, 
                0x1279: `StroyEnergoKom`, 
                0x127A: `Lumitech Lighting Solution GmbH`, 
                0x127B: `Verdant Environmental Technologies`, 
                0x127C: `Alfred International`, 
                0x127D: `Sansi LED Lighting`, 
                0x127E: `Mindtree`, 
                0x127F: `Nordic Semiconductor ASA`, 
                0x1280: `Siterwell Electronics`, 
                0x1281: `Briloner Leuchten GmbH`, 
                0x1282: `Shenzhen SEI Technology`, 
                0x1283: `Copper Labs`, 
                0x1284: `Delta Dore`, 
                0x1285: `Hager Group`, 
                0x1286: `Shenzhen CoolKit Technology`, 
                0x1287: `Hangzhou Sky-Lighting`, 
                0x1288: `E.ON SE`, 
                0x1289: `Lidl Stiftung`, 
                0x128A: `Sichuan Changhong Network Technologies`, 
                0x128B: `NodOn`, 
                0x128C: `Jiangxi Innotech Technology`, 
                0x128D: `Mercator Pty`, 
                0x128E: `Beijing Ruying Tech`, 
                0x128F: `EGLO Leuchten GmbH`, 
                0x1290: `Pietro Fiorentini S.p.A`, 
                0x1291: `Zehnder Group Vaux-Andigny`, 
                0x1292: `BRK Brands`, 
                0x1293: `Askey Computer`, 
                0x1294: `PassiveBolt`, 
                0x1295: `AVM Audiovisuelles`, 
                0x1296: `Ningbo Suntech Lighting Tech`, 
                0x1297: `Societe en Commandite Stello`, 
                0x1298: `Vivint Smart Home`, 
                0x1299: `Namron`, 
                0x129A: `RADEMACHER Geraete Elektronik GmbH`, 
                0x129B: `OMO Systems`, 
                0x129C: `Siglis`, 
                0x129D: `IMHOTEP CREATION`, 
                0x129E: `icasa`, 
                0x129F: `Level Home`, 
                0x1300: `TIS Control`, 
                0x1301: `Radisys India`, 
                0x1302: `Veea`, 
                0x1303: `FELL Technology`, 
                0x1304: `Sowilo Design Services`, 
                0x1305: `Lexi Devices`, 
                0x1306: `Lifi Labs`, 
                0x1307: `GRUNDFOS Holding`, 
                0x1308: `SOURCING & CREATION`, 
                0x1309: `Kraken Technologies`, 
                0x130A: `EVE SYSTEMS`, 
                0x130B: `LITE-ON TECHNOLOGY CORPORATION`, 
                0x130C: `Focalcrest`, 
                0x130D: `Bouffalo Lab (Nanjing)`, 
                0x130E: `Wyze Labs`, 
                0x1337: `Datek Wireless AS`, 
                0x1994: `Gewiss S.p.A.`, 
                0x2794: `Climax Technology Cp., Ltd.`,  },
                
            })),
            NeighborTable: makeType<ZigBee.IZDP.IArgNeighborTable, ZigBee.IZDP.IArgNeighborTablePayload>(base.array, ()=>({
                name: `Neighbor Table`,
                description: ``,
                arrayType: ZigBee.ZDP.Types.NeighborTableEntry,
                
            })),
            NeighborTableEntry: makeType<ZigBee.IZDP.IArgNeighborTableEntry, ZigBee.IZDP.IArgNeighborTableEntryPayload>(base.local, ()=>({
                name: `Neighbor Table Entry`,
                description: ``,
                payload: { 
                    PanId: ZigBee.ZDP.Types.IeeeAddress,
                    IeeeAddress: ZigBee.ZDP.Types.IeeeAddress,
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    NeighborType: ZigBee.ZDP.Types.NeighborType,
                    RxOnWhenIdle: ZigBee.ZDP.Types.RxOnWhenIdle,
                    Relationship: ZigBee.ZDP.Types.Relationship,
                    PermitJoining: ZigBee.ZDP.Types.PermitJoining,
                    Depth: ZigBee.ZDP.Types.Depth,
                    Lqi: ZigBee.ZDP.Types.Lqi,
                 },
                
            })),
            NeighborType: makeType<ZigBee.IZDP.IArgNeighborType, ZigBee.IZDP.IArgNeighborTypePayload>(base.enum8, ()=>({
                name: `Neighbor Type`,
                description: ``,
                values: { 
                0x00: `Coordinator`, 
                0x01: `Router`, 
                0x03: `End device`, 
                0x04: `Unknown`,  },
                
            })),
            NodeDescSize: makeType<ZigBee.IZDP.IArgNodeDescSize, ZigBee.IZDP.IArgNodeDescSizePayload>(base.u8, ()=>({
                name: `Node Desc Size`,
                description: `Size in bytes of the Node Descriptor`,
                
            })),
            NodeDescriptor: makeType<ZigBee.IZDP.IArgNodeDescriptor, ZigBee.IZDP.IArgNodeDescriptorPayload>(base.local, ()=>({
                name: `Node Descriptor`,
                description: ``,
                payload: { 
                    LogicalType: ZigBee.ZDP.Types.LogicalType,
                    ComplexDescriptorAvailable: ZigBee.ZDP.Types.ComplexDescriptorAvailable,
                    UserDescriptorAvailable: ZigBee.ZDP.Types.UserDescriptorAvailable,
                    ApsFlags: ZigBee.ZDP.Types.ApsFlags,
                    FrequencyBands: ZigBee.ZDP.Types.FrequencyBands,
                    MacCapabilityFlags: ZigBee.ZDP.Types.MacCapabilityFlags,
                    ManufacturerCode: ZigBee.ZDP.Types.ManufacturerCode,
                    MaxBufferSize: ZigBee.ZDP.Types.MaxBufferSize,
                    MaxRxSize: ZigBee.ZDP.Types.MaxRxSize,
                    ServerMask: ZigBee.ZDP.Types.ServerMask,
                    MaxTxSize: ZigBee.ZDP.Types.MaxTxSize,
                    DescriptorCapability: ZigBee.ZDP.Types.DescriptorCapability,
                 },
                
            })),
            PowerLevel: makeType<ZigBee.IZDP.IArgPowerLevel, ZigBee.IZDP.IArgPowerLevelPayload>(base.enum8, ()=>({
                name: `Power Level`,
                description: ``,
                values: { 
                0x00: `Critical`, 
                0x04: `33%`, 
                0x08: `66%`, 
                0x0C: `100%`,  },
                
            })),
            PowerMode: makeType<ZigBee.IZDP.IArgPowerMode, ZigBee.IZDP.IArgPowerModePayload>(base.bmp8, ()=>({
                name: `Power Mode`,
                description: ``,
                bits: { 
                0: `Constant Power Available`, 
                1: `Rechargeable battery Available`, 
                2: `Disposable Battery Available`, 
                4: `Check In Periodically`, 
                5: `Check In on Action`,  },
                
            })),
            PowerSource: makeType<ZigBee.IZDP.IArgPowerSource, ZigBee.IZDP.IArgPowerSourcePayload>(base.bmp8, ()=>({
                name: `Power Source`,
                description: ``,
                bits: { 
                0: `Mains power`, 
                1: `Rechargeable battery`, 
                2: `Disposable battery`,  },
                
            })),
            NwkAddress: makeType<ZigBee.IZDP.IArgNwkAddress, ZigBee.IZDP.IArgNwkAddressPayload>(base.u16, ()=>({
                name: `NWK Address`,
                description: `is a 16-bit Network address`,
                
            })),
            OutClusterList: makeType<ZigBee.IZDP.IArgOutClusterList, ZigBee.IZDP.IArgOutClusterListPayload>(base.set, ()=>({
                name: `Out Cluster List`,
                description: `is a list of output clusters`,
                arrayType: base.u16,
                
            })),
            PermitJoining: makeType<ZigBee.IZDP.IArgPermitJoining, ZigBee.IZDP.IArgPermitJoiningPayload>(base.enum8, ()=>({
                name: `Permit Joining`,
                description: ``,
                values: { 
                0x00: `Not permitted`, 
                0x01: `Permitted`, 
                0x02: `Unknown`,  },
                
            })),
            PowerDescSize: makeType<ZigBee.IZDP.IArgPowerDescSize, ZigBee.IZDP.IArgPowerDescSizePayload>(base.u8, ()=>({
                name: `Power Desc Size`,
                description: `Size in bytes of the Power Descriptor`,
                
            })),
            PowerDescriptor: makeType<ZigBee.IZDP.IArgPowerDescriptor, ZigBee.IZDP.IArgPowerDescriptorPayload>(base.local, ()=>({
                name: `Power Descriptor`,
                description: ``,
                payload: { 
                    PowerMode: ZigBee.ZDP.Types.PowerMode,
                    ActivePowerSource: ZigBee.ZDP.Types.PowerSource,
                    CurrentPowerSource: ZigBee.ZDP.Types.PowerSource,
                    PowerLevel: ZigBee.ZDP.Types.PowerLevel,
                 },
                
            })),
            ProfileId: makeType<ZigBee.IZDP.IArgProfileId, ZigBee.IZDP.IArgProfileIdPayload>(base.u16, ()=>({
                name: `Profile ID`,
                description: ``,
                values: { 
                0x0104: `Home Automation`, 
                0xA1E0: `Green Power`, 
                0xC05E: `Light Link`,  },
                
            })),
            Relationship: makeType<ZigBee.IZDP.IArgRelationship, ZigBee.IZDP.IArgRelationshipPayload>(base.enum8, ()=>({
                name: `Relationship`,
                description: ``,
                values: { 
                0x00: `Parent`, 
                0x01: `Child`, 
                0x02: `Sibling`, 
                0x03: `None`, 
                0x04: `Previous Child`,  },
                
            })),
            RxOnWhenIdle: makeType<ZigBee.IZDP.IArgRxOnWhenIdle, ZigBee.IZDP.IArgRxOnWhenIdlePayload>(base.enum8, ()=>({
                name: `Rx On When Idle`,
                description: ``,
                values: { 
                0x00: `Receiver is off`, 
                0x01: `Receiver is on`, 
                0x02: `Unknown`,  },
                
            })),
            ScannedChannels: makeType<ZigBee.IZDP.IArgScannedChannels, ZigBee.IZDP.IArgScannedChannelsPayload>(base.bmp32, ()=>({
                name: `Scanned Channels`,
                description: ``,
                
            })),
            ServerMask: makeType<ZigBee.IZDP.IArgServerMask, ZigBee.IZDP.IArgServerMaskPayload>(base.bmp16, ()=>({
                name: `Server Mask`,
                description: ``,
                bits: { 
                0: `Primary Trust Center`, 
                1: `Backup Trust Center`, 
                2: `Primary Binding Table Cache`, 
                3: `Backup Binding Table Cache`, 
                4: `Primary Discovery Cache`, 
                5: `Backup Discovery Cache`, 
                6: `Network Manager`,  },
                
            })),
            SimpleDescSizeList: makeType<ZigBee.IZDP.IArgSimpleDescSizeList, ZigBee.IZDP.IArgSimpleDescSizeListPayload>(base.set, ()=>({
                name: `Simple Desc Size List`,
                description: `List of sizes for the different Simple Descriptors`,
                arrayType: base.u8,
                
            })),
            SimpleDescriptor: makeType<ZigBee.IZDP.IArgSimpleDescriptor, ZigBee.IZDP.IArgSimpleDescriptorPayload>(base.local, ()=>({
                name: `Simple Descriptor`,
                description: ``,
                payload: { 
                    Endpoint: ZigBee.ZDP.Types.Endpoint,
                    ProfileId: ZigBee.ZDP.Types.ProfileId,
                    DeviceType: ZigBee.ZDP.Types.DeviceType,
                    DeviceVersion: ZigBee.ZDP.Types.DeviceVersion,
                    InClusterList: ZigBee.ZDP.Types.InClusterList,
                    OutClusterList: ZigBee.ZDP.Types.OutClusterList,
                 },
                
            })),
            SourceAddress: makeType<ZigBee.IZDP.IArgSourceAddress, ZigBee.IZDP.IArgSourceAddressPayload>(base.uid, ()=>({
                name: `Source Address`,
                description: `of device generating the request`,
                
            })),
            SourceEndpoint: makeType<ZigBee.IZDP.IArgSourceEndpoint, ZigBee.IZDP.IArgSourceEndpointPayload>(base.u8, ()=>({
                name: `Source Endpoint`,
                description: `of device generating the request`,
                
            })),
            StartIndex: makeType<ZigBee.IZDP.IArgStartIndex, ZigBee.IZDP.IArgStartIndexPayload>(base.u8, ()=>({
                name: `Start Index`,
                description: `provides the starting index for the requested elements of the associated list.`,
                
            })),
            Status: makeType<ZigBee.IZDP.IArgStatus, ZigBee.IZDP.IArgStatusPayload>(base.enum8, ()=>({
                name: `Status`,
                description: "Code, command is normally empty unless status is `Success`",
                values: { 
                0x00: `Success`, 
                0x80: `Invalid Request Type`, 
                0x81: `Device Not Found`, 
                0x82: `Invalid Endpoint`, 
                0x83: `Not Active`, 
                0x84: `Not Supported`, 
                0x85: `Timeout`, 
                0x86: `No Match`, 
                0x88: `No Entry`, 
                0x89: `No Descriptor`, 
                0x8A: `Insufficient Space`, 
                0x8B: `Not Permitted`, 
                0x8C: `Table Full`, 
                0x8D: `Not Authorized`,  },
                
            })),
            TotalEntries: makeType<ZigBee.IZDP.IArgTotalEntries, ZigBee.IZDP.IArgTotalEntriesPayload>(base.u8, ()=>({
                name: `Total Entries`,
                description: `is the total number of entries that can be queried for`,
                
            })),
            TotalTransmissions: makeType<ZigBee.IZDP.IArgTotalTransmissions, ZigBee.IZDP.IArgTotalTransmissionsPayload>(base.u16, ()=>({
                name: `Total Transmissions`,
                description: ``,
                
            })),
            TransmissionFailures: makeType<ZigBee.IZDP.IArgTransmissionFailures, ZigBee.IZDP.IArgTransmissionFailuresPayload>(base.u16, ()=>({
                name: `Transmission Failures`,
                description: ``,
                
            })),
            UserDescriptor: makeType<ZigBee.IZDP.IArgUserDescriptor, ZigBee.IZDP.IArgUserDescriptorPayload>(base.cstring, ()=>({
                name: `User Descriptor`,
                description: `is a user provided ASCII string of 16 characters set on a remote device.
If the string is shorter than 16 characters it should be padded with space characters (0x20).
Control characters (0x00-0x1f) are not allowed.`,
                
            })),
            UserDescriptorAvailable: makeType<ZigBee.IZDP.IArgUserDescriptorAvailable, ZigBee.IZDP.IArgUserDescriptorAvailablePayload>(base.bool, ()=>({
                name: `User Descriptor Available`,
                description: ``,
                
            })),
        },
        Commands: { 
            NwkAddressRequest: makeType<ZigBee.IZDP.ICmdNwkAddressRequest, ZigBee.IZDP.ICmdNwkAddressRequestPayload>(command, () => ({
                name: `NWK Address Request`,
                description: `queries the 16-bit address of the Remote Device based on its known IEEE address.
The destination addressing on this command shall be unicast or broadcast to all
devices supporting RX on when Idle (0xFFFD)`,
                id: 0x0000,
                payload: { 
                    IeeeAddress: ZigBee.ZDP.Types.IeeeAddress,
                    RequestType: ZigBee.ZDP.Types.RequestType,
                    StartIndex: ZigBee.ZDP.Types.StartIndex,
                }
            })),

            NwkAddressResponse: makeType<ZigBee.IZDP.ICmdNwkAddressResponse, ZigBee.IZDP.ICmdNwkAddressResponsePayload>(command, () => ({
                name: `NWK Address Response`,
                description: ``,
                id: 0x8000,
                payload: { 
                    Status: ZigBee.ZDP.Types.Status,
                    IeeeAddress: ZigBee.ZDP.Types.IeeeAddress,
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    StartIndex: ZigBee.ZDP.Types.StartIndex,
                    AssociatedDevices: ZigBee.ZDP.Types.AssociatedDevices,
                }
            })),

            IeeeAddressRequest: makeType<ZigBee.IZDP.ICmdIeeeAddressRequest, ZigBee.IZDP.ICmdIeeeAddressRequestPayload>(command, () => ({
                name: `IEEE Address Request`,
                description: `queries the 64-bit IEEE address of the Remote Device based on its known 16-bit NWK address.
The command should always be sent by unicast.`,
                id: 0x0001,
                payload: { 
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    RequestType: ZigBee.ZDP.Types.RequestType,
                    StartIndex: ZigBee.ZDP.Types.StartIndex,
                }
            })),

            IeeeAddressResponse: makeType<ZigBee.IZDP.ICmdIeeeAddressResponse, ZigBee.IZDP.ICmdIeeeAddressResponsePayload>(command, () => ({
                name: `IEEE Address Response`,
                description: ``,
                id: 0x8001,
                payload: { 
                    Status: ZigBee.ZDP.Types.Status,
                    IeeeAddress: ZigBee.ZDP.Types.IeeeAddress,
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    StartIndex: ZigBee.ZDP.Types.StartIndex,
                    AssociatedDevices: ZigBee.ZDP.Types.AssociatedDevices,
                }
            })),

            NodeDescRequest: makeType<ZigBee.IZDP.ICmdNodeDescRequest, ZigBee.IZDP.ICmdNodeDescRequestPayload>(command, () => ({
                name: `Node Desc Request`,
                description: `queries the node descriptor of a remote device. Should be unicast to the remote device directly,
or to a device that contains the discovery information of the remote device.`,
                id: 0x0002,
                payload: { 
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                }
            })),

            NodeDescResponse: makeType<ZigBee.IZDP.ICmdNodeDescResponse, ZigBee.IZDP.ICmdNodeDescResponsePayload>(command, () => ({
                name: `Node Desc Response`,
                description: ``,
                id: 0x8002,
                payload: { 
                    Status: ZigBee.ZDP.Types.Status,
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    NodeDescriptor: ZigBee.ZDP.Types.NodeDescriptor,
                }
            })),

            PowerDescRequest: makeType<ZigBee.IZDP.ICmdPowerDescRequest, ZigBee.IZDP.ICmdPowerDescRequestPayload>(command, () => ({
                name: `Power Desc Request`,
                description: `queries the power descriptor of a remote device. Should be unicast to the remote device directly,
or to a device that contains the discovery information of the remote device.`,
                id: 0x0003,
                payload: { 
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                }
            })),

            PowerDescResponse: makeType<ZigBee.IZDP.ICmdPowerDescResponse, ZigBee.IZDP.ICmdPowerDescResponsePayload>(command, () => ({
                name: `Power Desc Response`,
                description: ``,
                id: 0x8003,
                payload: { 
                    Status: ZigBee.ZDP.Types.Status,
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    PowerDescriptor: ZigBee.ZDP.Types.PowerDescriptor,
                }
            })),

            SimpleDescRequest: makeType<ZigBee.IZDP.ICmdSimpleDescRequest, ZigBee.IZDP.ICmdSimpleDescRequestPayload>(command, () => ({
                name: `Simple Desc Request`,
                description: `queries the Simple Descriptor of a remote device on a specific endpoint.
Simple Descriptor contains information about which clusters the device supports on the given endpoint.
Should be unicast to the remote device directly, or to a device that contains the discovery information
of the remote device.`,
                id: 0x0004,
                payload: { 
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    Endpoint: ZigBee.ZDP.Types.Endpoint,
                }
            })),

            SimpleDescResponse: makeType<ZigBee.IZDP.ICmdSimpleDescResponse, ZigBee.IZDP.ICmdSimpleDescResponsePayload>(command, () => ({
                name: `Simple Desc Response`,
                description: ``,
                id: 0x8004,
                payload: { 
                    Status: ZigBee.ZDP.Types.Status,
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    SimpleDescriptor: ZigBee.ZDP.Types.SimpleDescriptor,
                }
            })),

            ActiveEndpointRequest: makeType<ZigBee.IZDP.ICmdActiveEndpointRequest, ZigBee.IZDP.ICmdActiveEndpointRequestPayload>(command, () => ({
                name: `Active Endpoint Request`,
                description: `queries the remote device for a list of active endpoints. Should be unicast to the remote device directly,
or to a device that contains the discovery information of the remote device.`,
                id: 0x0005,
                payload: { 
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                }
            })),

            ActiveEndpointResponse: makeType<ZigBee.IZDP.ICmdActiveEndpointResponse, ZigBee.IZDP.ICmdActiveEndpointResponsePayload>(command, () => ({
                name: `Active Endpoint Response`,
                description: ``,
                id: 0x8005,
                payload: { 
                    Status: ZigBee.ZDP.Types.Status,
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    EndpointList: ZigBee.ZDP.Types.EndpointList,
                }
            })),

            MatchDescRequest: makeType<ZigBee.IZDP.ICmdMatchDescRequest, ZigBee.IZDP.ICmdMatchDescRequestPayload>(command, () => ({
                name: `Match Desc Request`,
                description: `is used to find remote devices supporting a specific simple descriptor match criterion. Normally broadcast
to all devices supporting RX on when Idle (0xFFFD)`,
                id: 0x0006,
                payload: { 
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    ProfileId: ZigBee.ZDP.Types.ProfileId,
                    InClusterList: ZigBee.ZDP.Types.InClusterList,
                    OutClusterList: ZigBee.ZDP.Types.OutClusterList,
                }
            })),

            MatchDescResponse: makeType<ZigBee.IZDP.ICmdMatchDescResponse, ZigBee.IZDP.ICmdMatchDescResponsePayload>(command, () => ({
                name: `Match Desc Response`,
                description: ``,
                id: 0x8006,
                payload: { 
                    Status: ZigBee.ZDP.Types.Status,
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    EndpointList: ZigBee.ZDP.Types.EndpointList,
                }
            })),

            ComplexDescRequest: makeType<ZigBee.IZDP.ICmdComplexDescRequest, ZigBee.IZDP.ICmdComplexDescRequestPayload>(command, () => ({
                name: `Complex Desc Request`,
                description: `queries the Complex Descriptor of a remote device. Should be unicast to the remote device directly,
or to a device that contains the discovery information of the remote device.`,
                id: 0x0010,
                payload: { 
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                }
            })),

            ComplexDescResponse: makeType<ZigBee.IZDP.ICmdComplexDescResponse, ZigBee.IZDP.ICmdComplexDescResponsePayload>(command, () => ({
                name: `Complex Desc Response`,
                description: ``,
                id: 0x8010,
                payload: { 
                    Status: ZigBee.ZDP.Types.Status,
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    ComplexDescriptor: ZigBee.ZDP.Types.ComplexDescriptor,
                }
            })),

            UserDescRequest: makeType<ZigBee.IZDP.ICmdUserDescRequest, ZigBee.IZDP.ICmdUserDescRequestPayload>(command, () => ({
                name: `User Desc Request`,
                description: `queries the User Descriptor of a remote device. Should be unicast to the remote device directly,
or to a device that contains the discovery information of the remote device.`,
                id: 0x0011,
                payload: { 
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                }
            })),

            UserDescResponse: makeType<ZigBee.IZDP.ICmdUserDescResponse, ZigBee.IZDP.ICmdUserDescResponsePayload>(command, () => ({
                name: `User Desc Response`,
                description: ``,
                id: 0x8011,
                payload: { 
                    Status: ZigBee.ZDP.Types.Status,
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    UserDescriptor: ZigBee.ZDP.Types.UserDescriptor,
                }
            })),

            DiscoveryCacheRequest: makeType<ZigBee.IZDP.ICmdDiscoveryCacheRequest, ZigBee.IZDP.ICmdDiscoveryCacheRequestPayload>(command, () => ({
                name: `Discovery Cache Request`,
                description: `locates a Primary Discovery Cache on the network. Should be addressed to broadcast RXOnWhenIdle (0xFFFD)`,
                id: 0x0012,
                payload: { 
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    IeeeAddress: ZigBee.ZDP.Types.IeeeAddress,
                }
            })),

            DeviceAnnounce: makeType<ZigBee.IZDP.ICmdDeviceAnnounce, ZigBee.IZDP.ICmdDeviceAnnouncePayload>(command, () => ({
                name: `Device Announce`,
                description: `is sent by a joining or returning device, identifying it's NWK address, IEEE address and capabilities.
Normally sent to broadcast RXOnWhenIdle (0xFFFD)`,
                id: 0x0013,
                payload: { 
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    IeeeAddress: ZigBee.ZDP.Types.IeeeAddress,
                    Capability: ZigBee.ZDP.Types.Capability,
                }
            })),

            UserDescSetRequest: makeType<ZigBee.IZDP.ICmdUserDescSetRequest, ZigBee.IZDP.ICmdUserDescSetRequestPayload>(command, () => ({
                name: `User Desc Set Request`,
                description: `requests the remote device to update it's user descriptor.`,
                id: 0x0014,
                payload: { 
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    UserDescriptor: ZigBee.ZDP.Types.UserDescriptor,
                }
            })),

            SystemServerDiscoverRequest: makeType<ZigBee.IZDP.ICmdSystemServerDiscoverRequest, ZigBee.IZDP.ICmdSystemServerDiscoverRequestPayload>(command, () => ({
                name: `System Server Discover Request`,
                description: `discovers the location of a particular system server or servers as indicated by the Server Mask. Should be sent to broadcast RXOnWhenIdle (0xFFFD)`,
                id: 0x0015,
                payload: { 
                    ServerMask: ZigBee.ZDP.Types.ServerMask,
                }
            })),

            DiscoveryStoreRequest: makeType<ZigBee.IZDP.ICmdDiscoveryStoreRequest, ZigBee.IZDP.ICmdDiscoveryStoreRequestPayload>(command, () => ({
                name: `Discovery Store Request`,
                description: `sent to a Discovery Cache Node to allocate memory of the provided sizes for cache storage.
If the node already exists in cache, it will be removed to allow storage of updated values.
Should be sent to the unicast address of a discovery cache device.`,
                id: 0x0016,
                payload: { 
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    IeeeAddress: ZigBee.ZDP.Types.IeeeAddress,
                    NodeDescSize: ZigBee.ZDP.Types.NodeDescSize,
                    PowerDescSize: ZigBee.ZDP.Types.PowerDescSize,
                    ActiveEndpointSize: ZigBee.ZDP.Types.ActiveEndpointSize,
                    SimpleDescSizeList: ZigBee.ZDP.Types.SimpleDescSizeList,
                }
            })),

            NodeDescStoreRequest: makeType<ZigBee.IZDP.ICmdNodeDescStoreRequest, ZigBee.IZDP.ICmdNodeDescStoreRequestPayload>(command, () => ({
                name: `Node Desc Store Request`,
                description: `sent to a Discovery Cache Node to store the provided Node Descriptor.
A DiscoveryStoreRequest should be sent and acknowledged with successful status before requesting storage.
Should be sent to the unicast address of a discovery cache device.`,
                id: 0x0017,
                payload: { 
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    IeeeAddress: ZigBee.ZDP.Types.IeeeAddress,
                    Capability: ZigBee.ZDP.Types.Capability,
                    ManufacturerCode: ZigBee.ZDP.Types.ManufacturerCode,
                    MaxBufferSize: ZigBee.ZDP.Types.MaxBufferSize,
                    MaxRxSize: ZigBee.ZDP.Types.MaxRxSize,
                    ServerMask: ZigBee.ZDP.Types.ServerMask,
                    MaxTxSize: ZigBee.ZDP.Types.MaxTxSize,
                }
            })),

            PowerDescStoreRequest: makeType<ZigBee.IZDP.ICmdPowerDescStoreRequest, ZigBee.IZDP.ICmdPowerDescStoreRequestPayload>(command, () => ({
                name: `Power Desc Store Request`,
                description: `sent to a Discovery Cache Node to store the provided Power Descriptor.
A DiscoveryStoreRequest should be sent and acknowledged with successful status before requesting storage.
Should be sent to the unicast address of a discovery cache device.`,
                id: 0x0018,
                payload: { 
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    IeeeAddress: ZigBee.ZDP.Types.IeeeAddress,
                    PowerMode: ZigBee.ZDP.Types.PowerMode,
                    PowerSource: ZigBee.ZDP.Types.PowerSource,
                }
            })),

            ActiveEndpointStoreRequest: makeType<ZigBee.IZDP.ICmdActiveEndpointStoreRequest, ZigBee.IZDP.ICmdActiveEndpointStoreRequestPayload>(command, () => ({
                name: `Active Endpoint Store Request`,
                description: `sent to a Discovery Cache Node to store the provided Active Endpoint list.
A DiscoveryStoreRequest should be sent and acknowledged with successful status before requesting storage.
Should be sent to the unicast address of a discovery cache device.`,
                id: 0x0019,
                payload: { 
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    IeeeAddress: ZigBee.ZDP.Types.IeeeAddress,
                    EndpointList: ZigBee.ZDP.Types.EndpointList,
                }
            })),

            SimpleDescStoreRequest: makeType<ZigBee.IZDP.ICmdSimpleDescStoreRequest, ZigBee.IZDP.ICmdSimpleDescStoreRequestPayload>(command, () => ({
                name: `Simple Desc Store Request`,
                description: `sent to a Discovery Cache Node to store the provided Simple Descriptor.
A DiscoveryStoreRequest should be sent and acknowledged with successful status before requesting storage.
Should be sent to the unicast address of a discovery cache device.`,
                id: 0x001A,
                payload: { 
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    IeeeAddress: ZigBee.ZDP.Types.IeeeAddress,
                    SimpleDescriptor: ZigBee.ZDP.Types.SimpleDescriptor,
                }
            })),

            RemoveNodeCacheRequest: makeType<ZigBee.IZDP.ICmdRemoveNodeCacheRequest, ZigBee.IZDP.ICmdRemoveNodeCacheRequestPayload>(command, () => ({
                name: `Remove Node Cache Request`,
                description: `sent to a Discovery Cache Node will request it to remove cached values for the provided node. Should be sent to the unicast address of a discovery cache device.`,
                id: 0x001B,
                payload: { 
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    IeeeAddress: ZigBee.ZDP.Types.IeeeAddress,
                }
            })),

            FindNodeCacheRequest: makeType<ZigBee.IZDP.ICmdFindNodeCacheRequest, ZigBee.IZDP.ICmdFindNodeCacheRequestPayload>(command, () => ({
                name: `Find Node Cache Request`,
                description: `broadcast to the network will generate a response from the Primary Discovery Cache holding information
for the device of interest`,
                id: 0x001C,
                payload: { 
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    IeeeAddress: ZigBee.ZDP.Types.IeeeAddress,
                }
            })),

            ExtendedSimpleDescRequest: makeType<ZigBee.IZDP.ICmdExtendedSimpleDescRequest, ZigBee.IZDP.ICmdExtendedSimpleDescRequestPayload>(command, () => ({
                name: `Extended Simple Desc Request`,
                description: `should be unicast to the remote device or a discovery cache node. It is used to request information from
nodes that implements a larger number of clusters than can be described by a SimpleDescRequest`,
                id: 0x001D,
                payload: { 
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    Endpoint: ZigBee.ZDP.Types.Endpoint,
                    StartIndex: ZigBee.ZDP.Types.StartIndex,
                }
            })),

            ExtendedActiveEndpointRequest: makeType<ZigBee.IZDP.ICmdExtendedActiveEndpointRequest, ZigBee.IZDP.ICmdExtendedActiveEndpointRequestPayload>(command, () => ({
                name: `Extended Active Endpoint Request`,
                description: `should be unicast to the remote device or a discovery cache node. It is used to request information from
nodes that implements a larger number of endpoints than can be described by a ActiveEndpointRequest`,
                id: 0x001E,
                payload: { 
                    NwkAddress: ZigBee.ZDP.Types.NwkAddress,
                    StartIndex: ZigBee.ZDP.Types.StartIndex,
                }
            })),

            EndDeviceBindRequest: makeType<ZigBee.IZDP.ICmdEndDeviceBindRequest, ZigBee.IZDP.ICmdEndDeviceBindRequestPayload>(command, () => ({
                name: `End Device Bind Request`,
                description: "is sent to the ZigBee coordinator from two different devices simultaneously." + "\n" + 
"* If the supplied endpoint is outside the specified range - a reply is sent with status `InvalidEndpoint`" + "\n" + 
"* If only one device sends the request within a pre-configure timeout - a reply is sent with status `Timeout`" + "\n" + 
"* If the ProfileID doesn't match or none of the In/OutClusterList elements match - a reply is sent with status" + "\n" + 
"  `NoMatch`" + "\n" + 
"* Otherwise, a reply is sent with status `Success` to each device" + "\n" + 
"" + "\n" + 
"The Coordinator then needs to either send `BindRequest` or `UnbindRequest` for each matched `ClusterID`." + "\n" + 
"This is done by first issuing a `UnbindRequest` with any of the matched `ClusterID`:" + "\n" + 
"* If the reply status is `NoEntry` - `BindRequest` will instead be sent for each matched `ClusterID`" + "\n" + 
"* If the reply status is `Success` - additional unbind requests are sent the rest of the matched cluster" + "\n" + 
"" + "\n" + 
"This enables the request to toggle the binding.",
                id: 0x0020,
                payload: { 
                    BindingTarget: ZigBee.ZDP.Types.BindingTarget,
                    SourceAddress: ZigBee.ZDP.Types.SourceAddress,
                    SourceEndpoint: ZigBee.ZDP.Types.SourceEndpoint,
                    ProfileId: ZigBee.ZDP.Types.ProfileId,
                    InClusterList: ZigBee.ZDP.Types.InClusterList,
                    OutClusterList: ZigBee.ZDP.Types.OutClusterList,
                }
            })),

            BindRequest: makeType<ZigBee.IZDP.ICmdBindRequest, ZigBee.IZDP.ICmdBindRequestPayload>(command, () => ({
                name: `Bind Request`,
                description: ``,
                id: 0x0021,
                payload: { 
                    BindingTableEntry: ZigBee.ZDP.Types.BindingTableEntry,
                }
            })),

            BindResponse: makeType<ZigBee.IZDP.ICmdBindResponse, ZigBee.IZDP.ICmdBindResponsePayload>(command, () => ({
                name: `Bind Response`,
                description: ``,
                id: 0x8021,
                payload: { 
                    Status: ZigBee.ZDP.Types.Status,
                }
            })),

            MgmtBindRequest: makeType<ZigBee.IZDP.ICmdMgmtBindRequest, ZigBee.IZDP.ICmdMgmtBindRequestPayload>(command, () => ({
                name: `Mgmt Bind Request`,
                description: ``,
                id: 0x0033,
                payload: { 
                    StartIndex: ZigBee.ZDP.Types.StartIndex,
                }
            })),

            MgmtBindResponse: makeType<ZigBee.IZDP.ICmdMgmtBindResponse, ZigBee.IZDP.ICmdMgmtBindResponsePayload>(command, () => ({
                name: `Mgmt Bind Response`,
                description: ``,
                id: 0x8033,
                payload: { 
                    Status: ZigBee.ZDP.Types.Status,
                    TotalEntries: ZigBee.ZDP.Types.TotalEntries,
                    StartIndex: ZigBee.ZDP.Types.StartIndex,
                    BindingTable: ZigBee.ZDP.Types.BindingTable,
                }
            })),

            UnbindRequest: makeType<ZigBee.IZDP.ICmdUnbindRequest, ZigBee.IZDP.ICmdUnbindRequestPayload>(command, () => ({
                name: `Unbind Request`,
                description: ``,
                id: 0x0022,
                payload: { 
                    BindingTableEntry: ZigBee.ZDP.Types.BindingTableEntry,
                }
            })),

            UnbindResponse: makeType<ZigBee.IZDP.ICmdUnbindResponse, ZigBee.IZDP.ICmdUnbindResponsePayload>(command, () => ({
                name: `Unbind Response`,
                description: ``,
                id: 0x8022,
                payload: { 
                    Status: ZigBee.ZDP.Types.Status,
                }
            })),

            MgmtLqiRequest: makeType<ZigBee.IZDP.ICmdMgmtLqiRequest, ZigBee.IZDP.ICmdMgmtLqiRequestPayload>(command, () => ({
                name: `Mgmt Lqi Request`,
                description: ``,
                id: 0x0031,
                payload: { 
                    StartIndex: ZigBee.ZDP.Types.StartIndex,
                }
            })),

            MgmtLqiResponse: makeType<ZigBee.IZDP.ICmdMgmtLqiResponse, ZigBee.IZDP.ICmdMgmtLqiResponsePayload>(command, () => ({
                name: `Mgmt Lqi Response`,
                description: ``,
                id: 0x8031,
                payload: { 
                    Status: ZigBee.ZDP.Types.Status,
                    TotalEntries: ZigBee.ZDP.Types.TotalEntries,
                    StartIndex: ZigBee.ZDP.Types.StartIndex,
                    NeighborTable: ZigBee.ZDP.Types.NeighborTable,
                }
            })),

            MgmtNwkUpdate: makeType<ZigBee.IZDP.ICmdMgmtNwkUpdate, ZigBee.IZDP.ICmdMgmtNwkUpdatePayload>(command, () => ({
                name: `Mgmt Nwk Update`,
                description: ``,
                id: 0x8038,
                payload: { 
                    Status: ZigBee.ZDP.Types.Status,
                    ScannedChannels: ZigBee.ZDP.Types.ScannedChannels,
                    TotalTransmissions: ZigBee.ZDP.Types.TotalTransmissions,
                    TransmissionFailures: ZigBee.ZDP.Types.TransmissionFailures,
                    EnergyValues: ZigBee.ZDP.Types.EnergyValues,
                }
            })),
 }
    };

    export namespace Profile {

    }

    
    export const Closures = {
        Types: { 
            ConfigStatus: makeType<ZigBee.IClosures.IArgConfigStatus, ZigBee.IClosures.IArgConfigStatusPayload>(base.bmp8, ()=>({
                name: `Config / Status`,
                description: ``,
                id: 0x0007,
                report: false,
                read: true,
                write: false,
                require: false,
                bits: { 
                0: `Operational`, 
                1: `Online`, 
                2: `Commands Reversed`, 
                3: `Lift control is Closed Loop`, 
                4: `Tilt control is Closed Loop`, 
                5: `Lift: Encoder Controlled`, 
                6: `Tilt: Encoder Controlled`,  },
                
            })),
            LiftAccelerationTime: makeType<ZigBee.IClosures.IArgLiftAccelerationTime, ZigBee.IClosures.IArgLiftAccelerationTimePayload>(base.u16, ()=>({
                name: `Lift - Acceleration Time`,
                description: ``,
                id: 0x0015,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            LiftCurrentPosition: makeType<ZigBee.IClosures.IArgLiftCurrentPosition, ZigBee.IClosures.IArgLiftCurrentPositionPayload>(base.u16, ()=>({
                name: `Lift - Current Position`,
                description: ``,
                id: 0x0003,
                report: true,
                read: true,
                write: false,
                require: false,
                
            })),
            LiftCurrentPositionPercentage: makeType<ZigBee.IClosures.IArgLiftCurrentPositionPercentage, ZigBee.IClosures.IArgLiftCurrentPositionPercentagePayload>(base.u8, ()=>({
                name: `Lift Current Position Percentage`,
                description: ``,
                id: 0x0008,
                report: true,
                read: true,
                write: false,
                require: false,
                unit: units.Percent,
                
            })),
            LiftDecelerationTime: makeType<ZigBee.IClosures.IArgLiftDecelerationTime, ZigBee.IClosures.IArgLiftDecelerationTimePayload>(base.u16, ()=>({
                name: `Lift - Deceleration Time`,
                description: ``,
                id: 0x0016,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            LiftInstalledClosedLimit: makeType<ZigBee.IClosures.IArgLiftInstalledClosedLimit, ZigBee.IClosures.IArgLiftInstalledClosedLimitPayload>(base.u16, ()=>({
                name: `Lift - Installed Closed Limit`,
                description: ``,
                id: 0x0011,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            LiftInstalledOpenLimit: makeType<ZigBee.IClosures.IArgLiftInstalledOpenLimit, ZigBee.IClosures.IArgLiftInstalledOpenLimitPayload>(base.u16, ()=>({
                name: `Lift - Installed Open Limit`,
                description: ``,
                id: 0x0010,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            LiftIntermediateSetpoints: makeType<ZigBee.IClosures.IArgLiftIntermediateSetpoints, ZigBee.IClosures.IArgLiftIntermediateSetpointsPayload>(base.ostring, ()=>({
                name: `Lift - Intermediate Setpoints`,
                description: ``,
                id: 0x0018,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            LiftNumberOfActuations: makeType<ZigBee.IClosures.IArgLiftNumberOfActuations, ZigBee.IClosures.IArgLiftNumberOfActuationsPayload>(base.u16, ()=>({
                name: `Lift - Number of Actuations`,
                description: ``,
                id: 0x0005,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            LiftPhysicalClosedLimit: makeType<ZigBee.IClosures.IArgLiftPhysicalClosedLimit, ZigBee.IClosures.IArgLiftPhysicalClosedLimitPayload>(base.u16, ()=>({
                name: `Lift - Physical Closed Limit`,
                description: ``,
                id: 0x0001,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            LiftVelocity: makeType<ZigBee.IClosures.IArgLiftVelocity, ZigBee.IClosures.IArgLiftVelocityPayload>(base.u16, ()=>({
                name: `Lift - Velocity`,
                description: ``,
                id: 0x0014,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            Percentage: makeType<ZigBee.IClosures.IArgPercentage, ZigBee.IClosures.IArgPercentagePayload>(base.u8, ()=>({
                name: `Percentage`,
                description: ``,
                unit: units.Percent,
                
            })),
            Position: makeType<ZigBee.IClosures.IArgPosition, ZigBee.IClosures.IArgPositionPayload>(base.u16, ()=>({
                name: `Position`,
                description: ``,
                
            })),
            TiltAInstalledOpenLimit: makeType<ZigBee.IClosures.IArgTiltAInstalledOpenLimit, ZigBee.IClosures.IArgTiltAInstalledOpenLimitPayload>(base.u16, ()=>({
                name: `Tilt A - Installed Open Limit`,
                description: ``,
                id: 0x0012,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            TiltBInstalledOpenLimit: makeType<ZigBee.IClosures.IArgTiltBInstalledOpenLimit, ZigBee.IClosures.IArgTiltBInstalledOpenLimitPayload>(base.u16, ()=>({
                name: `Tilt B - Installed Open Limit`,
                description: ``,
                id: 0x0013,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            TiltCurrentPosition: makeType<ZigBee.IClosures.IArgTiltCurrentPosition, ZigBee.IClosures.IArgTiltCurrentPositionPayload>(base.u16, ()=>({
                name: `Tilt - Current Position`,
                description: ``,
                id: 0x0004,
                report: true,
                read: true,
                write: false,
                require: false,
                
            })),
            TiltCurrentPositionPercentage: makeType<ZigBee.IClosures.IArgTiltCurrentPositionPercentage, ZigBee.IClosures.IArgTiltCurrentPositionPercentagePayload>(base.u8, ()=>({
                name: `Tilt Current Position Percentage`,
                description: ``,
                id: 0x0009,
                report: true,
                read: true,
                write: false,
                require: false,
                unit: units.Percent,
                
            })),
            TiltIntermediateSetpoints: makeType<ZigBee.IClosures.IArgTiltIntermediateSetpoints, ZigBee.IClosures.IArgTiltIntermediateSetpointsPayload>(base.ostring, ()=>({
                name: `Tilt - Intermediate Setpoints`,
                description: ``,
                id: 0x0019,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            TiltNumberOfActuations: makeType<ZigBee.IClosures.IArgTiltNumberOfActuations, ZigBee.IClosures.IArgTiltNumberOfActuationsPayload>(base.u16, ()=>({
                name: `Tilt - Number of Actuations`,
                description: ``,
                id: 0x0006,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            TiltPhysicalClosedLimit: makeType<ZigBee.IClosures.IArgTiltPhysicalClosedLimit, ZigBee.IClosures.IArgTiltPhysicalClosedLimitPayload>(base.u16, ()=>({
                name: `Tilt - Physical Closed Limit`,
                description: ``,
                id: 0x0002,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            WindowCoveringMode: makeType<ZigBee.IClosures.IArgWindowCoveringMode, ZigBee.IClosures.IArgWindowCoveringModePayload>(base.bmp8, ()=>({
                name: `Window Covering Mode`,
                description: ``,
                id: 0x0017,
                report: false,
                read: true,
                write: true,
                require: false,
                bits: { 
                0: `Reversed`, 
                1: `Calibration Mode`, 
                2: `Maintenance Mode`, 
                3: `LED feedback`,  },
                
            })),
            WindowCoveringType: makeType<ZigBee.IClosures.IArgWindowCoveringType, ZigBee.IClosures.IArgWindowCoveringTypePayload>(base.enum8, ()=>({
                name: `Window Covering Type`,
                description: ``,
                id: 0x0000,
                report: false,
                read: true,
                write: false,
                require: false,
                values: { 
                0x00: `Rollershade`, 
                0x01: `Rollershade - 2 Motor`, 
                0x02: `Rollershade - Exterior`, 
                0x03: `Rollershade - Exterior - 2 Motor`, 
                0x04: `Drapery`, 
                0x05: `Awning`, 
                0x06: `Shutter`, 
                0x07: `Tilt Blind - Tilt Only`, 
                0x08: `Tilt Blind - Lift and Tilt`, 
                0x09: `Projector Screen`,  },
                
            })), },
        DoorLock: {
            ID: 0x0101,
            Name: `Door Lock`,
            Desc: ``,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        WindowCovering: {
            ID: 0x0102,
            Name: `Window Covering`,
            Desc: `The window covering cluster provides an interface for controlling and adjusting automatic window coverings such as drapery motors, automatic shades, and blinds.`,
            
            UpOpen: makeType<ZigBee.IClosures.WindowCovering.ICmdUpOpen, ZigBee.IClosures.WindowCovering.ICmdUpOpenPayload>(command, () => ({
                name: `Up / Open`,
                description: ``,
                id: 0x0000,
                payload: {}
            })),

            DownClose: makeType<ZigBee.IClosures.WindowCovering.ICmdDownClose, ZigBee.IClosures.WindowCovering.ICmdDownClosePayload>(command, () => ({
                name: `Down / Close`,
                description: ``,
                id: 0x0001,
                payload: {}
            })),

            Stop: makeType<ZigBee.IClosures.WindowCovering.ICmdStop, ZigBee.IClosures.WindowCovering.ICmdStopPayload>(command, () => ({
                name: `Stop`,
                description: ``,
                id: 0x0002,
                payload: {}
            })),

            GoToLiftValue: makeType<ZigBee.IClosures.WindowCovering.ICmdGoToLiftValue, ZigBee.IClosures.WindowCovering.ICmdGoToLiftValuePayload>(command, () => ({
                name: `Go To Lift Value`,
                description: ``,
                id: 0x0004,
                payload: { 
                    Position: ZigBee.Closures.Types.Position,
                }
            })),

            GoToLiftPercentage: makeType<ZigBee.IClosures.WindowCovering.ICmdGoToLiftPercentage, ZigBee.IClosures.WindowCovering.ICmdGoToLiftPercentagePayload>(command, () => ({
                name: `Go to Lift Percentage`,
                description: ``,
                id: 0x0005,
                payload: { 
                    Percentage: ZigBee.Closures.Types.Percentage,
                }
            })),

            GoToTiltValue: makeType<ZigBee.IClosures.WindowCovering.ICmdGoToTiltValue, ZigBee.IClosures.WindowCovering.ICmdGoToTiltValuePayload>(command, () => ({
                name: `Go to Tilt Value`,
                description: ``,
                id: 0x0007,
                payload: { 
                    Position: ZigBee.Closures.Types.Position,
                }
            })),

            GoToTiltPercentage: makeType<ZigBee.IClosures.WindowCovering.ICmdGoToTiltPercentage, ZigBee.IClosures.WindowCovering.ICmdGoToTiltPercentagePayload>(command, () => ({
                name: `Go to Tilt Percentage`,
                description: ``,
                id: 0x0008,
                payload: { 
                    Percentage: ZigBee.Closures.Types.Percentage,
                }
            })),

            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        }
    };
    
    ZigBee.Closures.DoorLock.Server.Attribute = { 
    };
    ZigBee.Closures.DoorLock.Client.Attribute = { 
    };
    ZigBee.Closures.DoorLock.Server.Command = { 
    };
    ZigBee.Closures.DoorLock.Client.Command = { 
    };
    ZigBee.Closures.WindowCovering.Server.Attribute = { 
        0x0000: ZigBee.Closures.Types.WindowCoveringType,
        0x0001: ZigBee.Closures.Types.LiftPhysicalClosedLimit,
        0x0002: ZigBee.Closures.Types.TiltPhysicalClosedLimit,
        0x0003: ZigBee.Closures.Types.LiftCurrentPosition,
        0x0004: ZigBee.Closures.Types.TiltCurrentPosition,
        0x0005: ZigBee.Closures.Types.LiftNumberOfActuations,
        0x0006: ZigBee.Closures.Types.TiltNumberOfActuations,
        0x0007: ZigBee.Closures.Types.ConfigStatus,
        0x0008: ZigBee.Closures.Types.LiftCurrentPositionPercentage,
        0x0009: ZigBee.Closures.Types.TiltCurrentPositionPercentage,
        0x0010: ZigBee.Closures.Types.LiftInstalledOpenLimit,
        0x0011: ZigBee.Closures.Types.LiftInstalledClosedLimit,
        0x0012: ZigBee.Closures.Types.TiltAInstalledOpenLimit,
        0x0013: ZigBee.Closures.Types.TiltBInstalledOpenLimit,
        0x0014: ZigBee.Closures.Types.LiftVelocity,
        0x0015: ZigBee.Closures.Types.LiftAccelerationTime,
        0x0016: ZigBee.Closures.Types.LiftDecelerationTime,
        0x0017: ZigBee.Closures.Types.WindowCoveringMode,
        0x0018: ZigBee.Closures.Types.LiftIntermediateSetpoints,
        0x0019: ZigBee.Closures.Types.TiltIntermediateSetpoints,
    };
    ZigBee.Closures.WindowCovering.Client.Attribute = { 
    };
    ZigBee.Closures.WindowCovering.Server.Command = { 
        0x00: ZigBee.Closures.WindowCovering.UpOpen,
        0x01: ZigBee.Closures.WindowCovering.DownClose,
        0x02: ZigBee.Closures.WindowCovering.Stop,
        0x04: ZigBee.Closures.WindowCovering.GoToLiftValue,
        0x05: ZigBee.Closures.WindowCovering.GoToLiftPercentage,
        0x07: ZigBee.Closures.WindowCovering.GoToTiltValue,
        0x08: ZigBee.Closures.WindowCovering.GoToTiltPercentage,
    };
    ZigBee.Closures.WindowCovering.Client.Command = { 
    };
    export const General = {
        Types: { 
            AlarmCode: makeType<ZigBee.IGeneral.IArgAlarmCode, ZigBee.IGeneral.IArgAlarmCodePayload>(base.enum8, ()=>({
                name: `Alarm code`,
                description: ``,
                
            })),
            AlarmCount: makeType<ZigBee.IGeneral.IArgAlarmCount, ZigBee.IGeneral.IArgAlarmCountPayload>(base.u16, ()=>({
                name: `Alarm Count`,
                description: `Number of alarms currently defined`,
                id: 0x0000,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            AlarmMask: makeType<ZigBee.IGeneral.IArgAlarmMask, ZigBee.IGeneral.IArgAlarmMaskPayload>(base.bmp8, ()=>({
                name: `Alarm Mask`,
                description: ``,
                id: 0x0013,
                report: false,
                read: true,
                write: true,
                require: false,
                bits: { 
                0: `General Hardware Fault`, 
                1: `General Software Fault`,  },
                
            })),
            AnalogMaxPresentValue: makeType<ZigBee.IGeneral.IArgAnalogMaxPresentValue, ZigBee.IGeneral.IArgAnalogMaxPresentValuePayload>(base.float, ()=>({
                name: `Analog Max Present Value`,
                description: ``,
                id: 0x0041,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            AnalogMinPresentValue: makeType<ZigBee.IGeneral.IArgAnalogMinPresentValue, ZigBee.IGeneral.IArgAnalogMinPresentValuePayload>(base.float, ()=>({
                name: `Analog Min Present Value`,
                description: ``,
                id: 0x0045,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            AnalogPresentValue: makeType<ZigBee.IGeneral.IArgAnalogPresentValue, ZigBee.IGeneral.IArgAnalogPresentValuePayload>(base.float, ()=>({
                name: `Analog Present value`,
                description: ``,
                id: 0x0055,
                report: true,
                read: true,
                write: true,
                require: false,
                
            })),
            AnalogPriority: makeType<ZigBee.IGeneral.IArgAnalogPriority, ZigBee.IGeneral.IArgAnalogPriorityPayload>(base.local, ()=>({
                name: `Analog Priority`,
                description: ``,
                payload: { 
                 },
                
            })),
            AnalogPriorityArray: makeType<ZigBee.IGeneral.IArgAnalogPriorityArray, ZigBee.IGeneral.IArgAnalogPriorityArrayPayload>(base.array, ()=>({
                name: `Analog Priority Array`,
                description: ``,
                id: 0x0057,
                report: false,
                read: true,
                write: true,
                require: false,
                arrayType: ZigBee.General.Types.AnalogPriority,
                
            })),
            AnalogRelinquishDefault: makeType<ZigBee.IGeneral.IArgAnalogRelinquishDefault, ZigBee.IGeneral.IArgAnalogRelinquishDefaultPayload>(base.float, ()=>({
                name: `Analog Relinquish Default`,
                description: ``,
                id: 0x0068,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            AnalogResolution: makeType<ZigBee.IGeneral.IArgAnalogResolution, ZigBee.IGeneral.IArgAnalogResolutionPayload>(base.float, ()=>({
                name: `Analog Resolution`,
                description: ``,
                id: 0x006A,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            ApplicationVersion: makeType<ZigBee.IGeneral.IArgApplicationVersion, ZigBee.IGeneral.IArgApplicationVersionPayload>(base.u8, ()=>({
                name: `Application Version`,
                description: ``,
                id: 0x0001,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            ApsDecryptFailures: makeType<ZigBee.IGeneral.IArgApsDecryptFailures, ZigBee.IGeneral.IArgApsDecryptFailuresPayload>(base.u16, ()=>({
                name: `APS Decrypt Failures`,
                description: ``,
                id: 0x0116,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            ApsFcFailure: makeType<ZigBee.IGeneral.IArgApsFcFailure, ZigBee.IGeneral.IArgApsFcFailurePayload>(base.u16, ()=>({
                name: `APS FC Failure`,
                description: ``,
                id: 0x0113,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            ApsRxBcast: makeType<ZigBee.IGeneral.IArgApsRxBcast, ZigBee.IGeneral.IArgApsRxBcastPayload>(base.u16, ()=>({
                name: `APS Rx Bcast`,
                description: ``,
                id: 0x0106,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            ApsRxUcast: makeType<ZigBee.IGeneral.IArgApsRxUcast, ZigBee.IGeneral.IArgApsRxUcastPayload>(base.u16, ()=>({
                name: `APS Rx Ucast`,
                description: ``,
                id: 0x0108,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            ApsTxBcast: makeType<ZigBee.IGeneral.IArgApsTxBcast, ZigBee.IGeneral.IArgApsTxBcastPayload>(base.u16, ()=>({
                name: `APS Tx Bcast`,
                description: ``,
                id: 0x0107,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            ApsTxUcastFail: makeType<ZigBee.IGeneral.IArgApsTxUcastFail, ZigBee.IGeneral.IArgApsTxUcastFailPayload>(base.u16, ()=>({
                name: `APS Tx Ucast Fail`,
                description: ``,
                id: 0x010B,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            ApsTxUcastRetry: makeType<ZigBee.IGeneral.IArgApsTxUcastRetry, ZigBee.IGeneral.IArgApsTxUcastRetryPayload>(base.u16, ()=>({
                name: `APS Tx Ucast Retry`,
                description: ``,
                id: 0x010A,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            ApsTxUcastSuccess: makeType<ZigBee.IGeneral.IArgApsTxUcastSuccess, ZigBee.IGeneral.IArgApsTxUcastSuccessPayload>(base.u16, ()=>({
                name: `APS Tx Ucast Success`,
                description: ``,
                id: 0x0109,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            ApsUnauthorizedKey: makeType<ZigBee.IGeneral.IArgApsUnauthorizedKey, ZigBee.IGeneral.IArgApsUnauthorizedKeyPayload>(base.u16, ()=>({
                name: `APS Unauthorized Key`,
                description: ``,
                id: 0x0114,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            AvgMacRetryPerApsMsgSent: makeType<ZigBee.IGeneral.IArgAvgMacRetryPerApsMsgSent, ZigBee.IGeneral.IArgAvgMacRetryPerApsMsgSentPayload>(base.u16, ()=>({
                name: `Avg MAC Retry per APS Msg Sent`,
                description: ``,
                id: 0x011B,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            BatteryAlarmMask: makeType<ZigBee.IGeneral.IArgBatteryAlarmMask, ZigBee.IGeneral.IArgBatteryAlarmMaskPayload>(base.bmp8, ()=>({
                name: `Battery Alarm Mask`,
                description: ``,
                id: 0x0035,
                report: false,
                read: true,
                write: true,
                require: false,
                bits: { 
                0: `Battery Voltage too low`, 
                1: `Battery Alarm 1`, 
                2: `Battery Alarm 2`, 
                3: `Battery Alarm 3`,  },
                
            })),
            BatteryAlarmState: makeType<ZigBee.IGeneral.IArgBatteryAlarmState, ZigBee.IGeneral.IArgBatteryAlarmStatePayload>(base.bmp32, ()=>({
                name: `Battery Alarm State`,
                description: ``,
                id: 0x003E,
                report: false,
                read: true,
                write: true,
                require: false,
                bits: { 
                0: `Minimum Threshold 1 for Voltage or Percentage reached for Battery Source 1`, 
                1: `Threshold 1 for Voltage or Percentage reached for Battery Source 1`, 
                10: `Minimum Threshold 1 for Voltage or Percentage reached for Battery Source 2`, 
                11: `Threshold 1 for Voltage or Percentage reached for Battery Source 2`, 
                12: `Threshold 2 for Voltage or Percentage reached for Battery Source 2`, 
                13: `Threshold 3 for Voltage or Percentage reached for Battery Source 2`, 
                2: `Threshold 2 for Voltage or Percentage reached for Battery Source 1`, 
                20: `Minimum Threshold 1 for Voltage or Percentage reached for Battery Source 3`, 
                21: `Threshold 1 for Voltage or Percentage reached for Battery Source 3`, 
                22: `Threshold 2 for Voltage or Percentage reached for Battery Source 3`, 
                23: `Threshold 3 for Voltage or Percentage reached for Battery Source 3`, 
                3: `Threshold 3 for Voltage or Percentage reached for Battery Source 1`, 
                30: `Mains power lost/unavailable`,  },
                
            })),
            BatteryCapacity: makeType<ZigBee.IGeneral.IArgBatteryCapacity, ZigBee.IGeneral.IArgBatteryCapacityPayload>(base.u16, ()=>({
                name: `Battery capacity`,
                description: ``,
                id: 0x0032,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.MilliAmpereHours,
                scale: 0.1,
                
            })),
            BatteryManufacturer: makeType<ZigBee.IGeneral.IArgBatteryManufacturer, ZigBee.IGeneral.IArgBatteryManufacturerPayload>(base.cstring, ()=>({
                name: `Battery Manufacturer`,
                description: ``,
                id: 0x0030,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            BatteryPercentageMinThreshold: makeType<ZigBee.IGeneral.IArgBatteryPercentageMinThreshold, ZigBee.IGeneral.IArgBatteryPercentageMinThresholdPayload>(base.u8, ()=>({
                name: `Battery Percentage Min Threshold`,
                description: ``,
                id: 0x003A,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Percent,
                scale: 2,
                
            })),
            BatteryPercentageThreshold1: makeType<ZigBee.IGeneral.IArgBatteryPercentageThreshold1, ZigBee.IGeneral.IArgBatteryPercentageThreshold1Payload>(base.u8, ()=>({
                name: `Battery Percentage Threshold 1`,
                description: ``,
                id: 0x003B,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Percent,
                scale: 2,
                
            })),
            BatteryPercentageThreshold2: makeType<ZigBee.IGeneral.IArgBatteryPercentageThreshold2, ZigBee.IGeneral.IArgBatteryPercentageThreshold2Payload>(base.u8, ()=>({
                name: `Battery Percentage Threshold 2`,
                description: ``,
                id: 0x003C,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Percent,
                scale: 2,
                
            })),
            BatteryPercentageThreshold3: makeType<ZigBee.IGeneral.IArgBatteryPercentageThreshold3, ZigBee.IGeneral.IArgBatteryPercentageThreshold3Payload>(base.u8, ()=>({
                name: `Battery Percentage Threshold 3`,
                description: ``,
                id: 0x003D,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Percent,
                scale: 2,
                
            })),
            BatteryQuantity: makeType<ZigBee.IGeneral.IArgBatteryQuantity, ZigBee.IGeneral.IArgBatteryQuantityPayload>(base.u8, ()=>({
                name: `Battery Quantity`,
                description: ``,
                id: 0x0033,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            BatteryRatedVoltage: makeType<ZigBee.IGeneral.IArgBatteryRatedVoltage, ZigBee.IGeneral.IArgBatteryRatedVoltagePayload>(base.u8, ()=>({
                name: `Battery Rated Voltage`,
                description: ``,
                id: 0x0034,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Volts,
                scale: 10,
                
            })),
            BatteryRemaining: makeType<ZigBee.IGeneral.IArgBatteryRemaining, ZigBee.IGeneral.IArgBatteryRemainingPayload>(base.u8, ()=>({
                name: `Battery Remaining`,
                description: ``,
                id: 0x0021,
                report: true,
                read: true,
                write: false,
                require: false,
                unit: units.Percent,
                scale: 2,
                
            })),
            BatterySize: makeType<ZigBee.IGeneral.IArgBatterySize, ZigBee.IGeneral.IArgBatterySizePayload>(base.enum8, ()=>({
                name: `Battery Size`,
                description: ``,
                id: 0x0031,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x00: `No Battery`, 
                0x01: `Built in`, 
                0x02: `Other`, 
                0x03: `AA`, 
                0x04: `AAA`, 
                0x05: `C`, 
                0x06: `D`, 
                0x07: `CR2`, 
                0x08: `CR123A`, 
                0xFF: `Unknown`,  },
                
            })),
            BatteryVoltage: makeType<ZigBee.IGeneral.IArgBatteryVoltage, ZigBee.IGeneral.IArgBatteryVoltagePayload>(base.u8, ()=>({
                name: `Battery Voltage`,
                description: ``,
                id: 0x0020,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.Volts,
                scale: 10,
                
            })),
            BatteryVoltageMinThreshold: makeType<ZigBee.IGeneral.IArgBatteryVoltageMinThreshold, ZigBee.IGeneral.IArgBatteryVoltageMinThresholdPayload>(base.u8, ()=>({
                name: `Battery Voltage Min Threshold`,
                description: ``,
                id: 0x0036,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Volts,
                scale: 10,
                
            })),
            BatteryVoltageThreshold1: makeType<ZigBee.IGeneral.IArgBatteryVoltageThreshold1, ZigBee.IGeneral.IArgBatteryVoltageThreshold1Payload>(base.u8, ()=>({
                name: `Battery Voltage Threshold 1`,
                description: ``,
                id: 0x0037,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Volts,
                scale: 10,
                
            })),
            BatteryVoltageThreshold2: makeType<ZigBee.IGeneral.IArgBatteryVoltageThreshold2, ZigBee.IGeneral.IArgBatteryVoltageThreshold2Payload>(base.u8, ()=>({
                name: `Battery Voltage Threshold 2`,
                description: ``,
                id: 0x0038,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Volts,
                scale: 10,
                
            })),
            BatteryVoltageThreshold3: makeType<ZigBee.IGeneral.IArgBatteryVoltageThreshold3, ZigBee.IGeneral.IArgBatteryVoltageThreshold3Payload>(base.u8, ()=>({
                name: `Battery Voltage Threshold 3`,
                description: ``,
                id: 0x0039,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Volts,
                scale: 10,
                
            })),
            BinaryActiveText: makeType<ZigBee.IGeneral.IArgBinaryActiveText, ZigBee.IGeneral.IArgBinaryActiveTextPayload>(base.cstring, ()=>({
                name: `Binary Active Text`,
                description: ``,
                id: 0x0004,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            BinaryInactiveText: makeType<ZigBee.IGeneral.IArgBinaryInactiveText, ZigBee.IGeneral.IArgBinaryInactiveTextPayload>(base.cstring, ()=>({
                name: `Binary Inactive Text`,
                description: ``,
                id: 0x002E,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            BinaryMaxOffTime: makeType<ZigBee.IGeneral.IArgBinaryMaxOffTime, ZigBee.IGeneral.IArgBinaryMaxOffTimePayload>(base.u32, ()=>({
                name: `Binary Max Off-time`,
                description: ``,
                id: 0x0043,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            BinaryMinOffTime: makeType<ZigBee.IGeneral.IArgBinaryMinOffTime, ZigBee.IGeneral.IArgBinaryMinOffTimePayload>(base.u32, ()=>({
                name: `Binary Min Off-time`,
                description: ``,
                id: 0x0042,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            BinaryPolarity: makeType<ZigBee.IGeneral.IArgBinaryPolarity, ZigBee.IGeneral.IArgBinaryPolarityPayload>(base.enum8, ()=>({
                name: `Binary Polarity`,
                description: ``,
                id: 0x0054,
                report: false,
                read: true,
                write: false,
                require: false,
                values: { 
                0x00: `Normal`, 
                0x01: `Reverse`,  },
                
            })),
            BinaryPresentValue: makeType<ZigBee.IGeneral.IArgBinaryPresentValue, ZigBee.IGeneral.IArgBinaryPresentValuePayload>(base.bool, ()=>({
                name: `Binary Present Value`,
                description: ``,
                id: 0x0055,
                report: true,
                read: true,
                write: true,
                require: false,
                
            })),
            BinaryPriority: makeType<ZigBee.IGeneral.IArgBinaryPriority, ZigBee.IGeneral.IArgBinaryPriorityPayload>(base.local, ()=>({
                name: `Binary Priority`,
                description: ``,
                payload: { 
                 },
                
            })),
            BinaryPriorityArray: makeType<ZigBee.IGeneral.IArgBinaryPriorityArray, ZigBee.IGeneral.IArgBinaryPriorityArrayPayload>(base.array, ()=>({
                name: `Binary Priority Array`,
                description: ``,
                id: 0x0057,
                report: false,
                read: true,
                write: true,
                require: false,
                arrayType: ZigBee.General.Types.BinaryPriority,
                
            })),
            BinaryRelinquishDefault: makeType<ZigBee.IGeneral.IArgBinaryRelinquishDefault, ZigBee.IGeneral.IArgBinaryRelinquishDefaultPayload>(base.bool, ()=>({
                name: `Binary Relinquish Default`,
                description: ``,
                id: 0x0068,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            CalculationPeriod: makeType<ZigBee.IGeneral.IArgCalculationPeriod, ZigBee.IGeneral.IArgCalculationPeriodPayload>(base.u16, ()=>({
                name: `Calculation Period`,
                description: ``,
                id: 0x0016,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Seconds,
                
            })),
            CheckInInterval: makeType<ZigBee.IGeneral.IArgCheckInInterval, ZigBee.IGeneral.IArgCheckInIntervalPayload>(base.u32, ()=>({
                name: `Check-in Interval`,
                description: ``,
                id: 0x0000,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            CheckInIntervalMin: makeType<ZigBee.IGeneral.IArgCheckInIntervalMin, ZigBee.IGeneral.IArgCheckInIntervalMinPayload>(base.u32, ()=>({
                name: `Check-in Interval Min`,
                description: ``,
                id: 0x0004,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            ChildMoved: makeType<ZigBee.IGeneral.IArgChildMoved, ZigBee.IGeneral.IArgChildMovedPayload>(base.u16, ()=>({
                name: `Child Moved`,
                description: ``,
                id: 0x0111,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            ClusterId: makeType<ZigBee.IGeneral.IArgClusterId, ZigBee.IGeneral.IArgClusterIdPayload>(base.u16, ()=>({
                name: `Cluster Id`,
                description: ``,
                
            })),
            ClusterRevision: makeType<ZigBee.IGeneral.IArgClusterRevision, ZigBee.IGeneral.IArgClusterRevisionPayload>(base.u16, ()=>({
                name: `Cluster Revision`,
                description: ``,
                id: 0xFFFD,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            Configuration: makeType<ZigBee.IGeneral.IArgConfiguration, ZigBee.IGeneral.IArgConfigurationPayload>(base.bmp16, ()=>({
                name: `Configuration`,
                description: ``,
                id: 0x0031,
                report: false,
                read: true,
                write: true,
                require: false,
                mnf: 0x100b,
                bits: { 
                0: `Touchlink enabled 0`, 
                1: `Touchlink enabled 1`, 
                3: `Touchlink enabled 2`,  },
                
            })),
            CurrentGroup: makeType<ZigBee.IGeneral.IArgCurrentGroup, ZigBee.IGeneral.IArgCurrentGroupPayload>(base.u16, ()=>({
                name: `Current Group`,
                description: ``,
                id: 0x0002,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            CurrentLevel: makeType<ZigBee.IGeneral.IArgCurrentLevel, ZigBee.IGeneral.IArgCurrentLevelPayload>(base.u8, ()=>({
                name: `Current Level`,
                description: `represents the current level of this device. Meaning of 'level' is device dependent.`,
                id: 0x0000,
                report: true,
                read: true,
                write: false,
                require: false,
                unit: units.Percent,
                scale: 2.54,
                
            })),
            CurrentScene: makeType<ZigBee.IGeneral.IArgCurrentScene, ZigBee.IGeneral.IArgCurrentScenePayload>(base.u8, ()=>({
                name: `Current Scene`,
                description: ``,
                id: 0x0001,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            CurrentTemperature: makeType<ZigBee.IGeneral.IArgCurrentTemperature, ZigBee.IGeneral.IArgCurrentTemperaturePayload>(base.s16, ()=>({
                name: `Current Temperature`,
                description: ``,
                id: 0x0000,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.DegreesCelsius,
                
            })),
            DateCode: makeType<ZigBee.IGeneral.IArgDateCode, ZigBee.IGeneral.IArgDateCodePayload>(base.cstring, ()=>({
                name: `Date Code`,
                description: ``,
                id: 0x0006,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            DefaultMoveRate: makeType<ZigBee.IGeneral.IArgDefaultMoveRate, ZigBee.IGeneral.IArgDefaultMoveRatePayload>(base.u8, ()=>({
                name: `Default Move Rate`,
                description: ``,
                id: 0x0014,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.PercentPerSecond,
                scale: 2.54,
                
            })),
            Device: makeType<ZigBee.IGeneral.IArgDevice, ZigBee.IGeneral.IArgDevicePayload>(base.uid, ()=>({
                name: `Device`,
                description: ``,
                
            })),
            DeviceEnabled: makeType<ZigBee.IGeneral.IArgDeviceEnabled, ZigBee.IGeneral.IArgDeviceEnabledPayload>(base.bool, ()=>({
                name: `Device Enabled`,
                description: ``,
                id: 0x0012,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            DeviceTempAlarmMask: makeType<ZigBee.IGeneral.IArgDeviceTempAlarmMask, ZigBee.IGeneral.IArgDeviceTempAlarmMaskPayload>(base.bmp8, ()=>({
                name: `Device Temp Alarm Mask`,
                description: ``,
                id: 0x0010,
                report: false,
                read: true,
                write: true,
                require: false,
                bits: { 
                0: `Temperature too low`, 
                1: `Temperature too high`,  },
                
            })),
            DisableLocalConfig: makeType<ZigBee.IGeneral.IArgDisableLocalConfig, ZigBee.IGeneral.IArgDisableLocalConfigPayload>(base.bmp8, ()=>({
                name: `Disable Local Config`,
                description: ``,
                id: 0x0014,
                report: false,
                read: true,
                write: true,
                require: false,
                bits: { 
                0: `Disable Factory Reset`, 
                1: `Disable Device Configuration`,  },
                
            })),
            Distance: makeType<ZigBee.IGeneral.IArgDistance, ZigBee.IGeneral.IArgDistancePayload>(base.u16, ()=>({
                name: `Distance`,
                description: ``,
                unit: units.Meters,
                scale: 10,
                
            })),
            DstEnd: makeType<ZigBee.IGeneral.IArgDstEnd, ZigBee.IGeneral.IArgDstEndPayload>(base.utc, ()=>({
                name: `Dst End`,
                description: ``,
                id: 0x0004,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            DstShift: makeType<ZigBee.IGeneral.IArgDstShift, ZigBee.IGeneral.IArgDstShiftPayload>(base.s32, ()=>({
                name: `Dst Shift`,
                description: ``,
                id: 0x0005,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Seconds,
                
            })),
            DstStart: makeType<ZigBee.IGeneral.IArgDstStart, ZigBee.IGeneral.IArgDstStartPayload>(base.utc, ()=>({
                name: `Dst Start`,
                description: ``,
                id: 0x0003,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            EffectIdentifier: makeType<ZigBee.IGeneral.IArgEffectIdentifier, ZigBee.IGeneral.IArgEffectIdentifierPayload>(base.enum8, ()=>({
                name: `Effect Identifier`,
                description: `when turning lights off`,
                values: { 
                0x00: `Delayed all off`, 
                0x01: `Dying light`,  },
                
            })),
            EffectVariant: makeType<ZigBee.IGeneral.IArgEffectVariant, ZigBee.IGeneral.IArgEffectVariantPayload>(base.enum8, ()=>({
                name: `Effect Variant`,
                description: ``,
                values: { 
                0x00: `Fade to off in 0.8 seconds / 20% dim up in 0.5 then fade to off in 1 second`, 
                0x01: `No Fade`, 
                0x02: `50% dim down in 0.8s then fade off in 12s`,  },
                
            })),
            FastPollTimeout: makeType<ZigBee.IGeneral.IArgFastPollTimeout, ZigBee.IGeneral.IArgFastPollTimeoutPayload>(base.u16, ()=>({
                name: `Fast Poll Timeout`,
                description: ``,
                id: 0x0003,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            FastPollTimeoutMax: makeType<ZigBee.IGeneral.IArgFastPollTimeoutMax, ZigBee.IGeneral.IArgFastPollTimeoutMaxPayload>(base.u16, ()=>({
                name: `Fast Poll Timeout Max`,
                description: ``,
                id: 0x0006,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            GenericDeviceClass: makeType<ZigBee.IGeneral.IArgGenericDeviceClass, ZigBee.IGeneral.IArgGenericDeviceClassPayload>(base.enum8, ()=>({
                name: `Generic Device Class`,
                description: `defines the field of application of the GenericDeviceType attribute`,
                id: 0x0008,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            GenericDeviceType: makeType<ZigBee.IGeneral.IArgGenericDeviceType, ZigBee.IGeneral.IArgGenericDeviceTypePayload>(base.enum8, ()=>({
                name: `Generic Device Type`,
                description: `allows an application to show an icon on a rich user interface (e.g. smartphone app)`,
                id: 0x0009,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            GlobalSceneControl: makeType<ZigBee.IGeneral.IArgGlobalSceneControl, ZigBee.IGeneral.IArgGlobalSceneControlPayload>(base.bool, ()=>({
                name: `Global Scene Control`,
                description: ``,
                id: 0x4000,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            GroupCapacity: makeType<ZigBee.IGeneral.IArgGroupCapacity, ZigBee.IGeneral.IArgGroupCapacityPayload>(base.u8, ()=>({
                name: `Group capacity`,
                description: `specifies remaining number of groups that can be added.
If set to 0xFE, at least one more group can be added (exact number unknown)
If set to 0xFF, it's unknown if any more groups can be added`,
                
            })),
            GroupId: makeType<ZigBee.IGeneral.IArgGroupId, ZigBee.IGeneral.IArgGroupIdPayload>(base.u16, ()=>({
                name: `Group ID`,
                description: ``,
                
            })),
            GroupList: makeType<ZigBee.IGeneral.IArgGroupList, ZigBee.IGeneral.IArgGroupListPayload>(base.set, ()=>({
                name: `Group list`,
                description: ``,
                arrayType: base.u16,
                
            })),
            GroupName: makeType<ZigBee.IGeneral.IArgGroupName, ZigBee.IGeneral.IArgGroupNamePayload>(base.cstring, ()=>({
                name: `Group name`,
                description: ``,
                
            })),
            GroupNameSupport: makeType<ZigBee.IGeneral.IArgGroupNameSupport, ZigBee.IGeneral.IArgGroupNameSupportPayload>(base.bmp8, ()=>({
                name: `Group Name Support`,
                description: ``,
                id: 0x0000,
                report: false,
                read: true,
                write: false,
                require: false,
                bits: { 
                7: `Names Supported`,  },
                
            })),
            HighTempDwellTripPoint: makeType<ZigBee.IGeneral.IArgHighTempDwellTripPoint, ZigBee.IGeneral.IArgHighTempDwellTripPointPayload>(base.u24, ()=>({
                name: `High Temp Dwell Trip Point`,
                description: ``,
                id: 0x0014,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Seconds,
                
            })),
            HighTempThreshold: makeType<ZigBee.IGeneral.IArgHighTempThreshold, ZigBee.IGeneral.IArgHighTempThresholdPayload>(base.s16, ()=>({
                name: `High Temp Threshold`,
                description: `If the current temperature goes above the threshold for longer
than the time specified by high temp dwell, an alarm will be triggered`,
                id: 0x0012,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.DegreesCelsius,
                
            })),
            HwVersion: makeType<ZigBee.IGeneral.IArgHwVersion, ZigBee.IGeneral.IArgHwVersionPayload>(base.u8, ()=>({
                name: `HW Version`,
                description: ``,
                id: 0x0003,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            IOApplicationType: makeType<ZigBee.IGeneral.IArgIOApplicationType, ZigBee.IGeneral.IArgIOApplicationTypePayload>(base.u32, ()=>({
                name: `I/O Application Type`,
                description: ``,
                id: 0x0100,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            IODescription: makeType<ZigBee.IGeneral.IArgIODescription, ZigBee.IGeneral.IArgIODescriptionPayload>(base.cstring, ()=>({
                name: `I/O Description`,
                description: ``,
                id: 0x001C,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            IOOutOfService: makeType<ZigBee.IGeneral.IArgIOOutOfService, ZigBee.IGeneral.IArgIOOutOfServicePayload>(base.bool, ()=>({
                name: `I/O Out of service`,
                description: ``,
                id: 0x0051,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            IOReliability: makeType<ZigBee.IGeneral.IArgIOReliability, ZigBee.IGeneral.IArgIOReliabilityPayload>(base.enum8, ()=>({
                name: `I/O Reliability`,
                description: ``,
                id: 0x0067,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x00: `No fault detected`, 
                0x01: `No Sensor`, 
                0x02: `Over Range`, 
                0x03: `Under Range`, 
                0x04: `Open Loop`, 
                0x05: `Shorted Loop`, 
                0x06: `No Output`, 
                0x07: `Unreliable (other)`, 
                0x08: `Process Error`, 
                0x09: `Multi state fault`, 
                0x0A: `Configuration Error`,  },
                
            })),
            IOStatusFlags: makeType<ZigBee.IGeneral.IArgIOStatusFlags, ZigBee.IGeneral.IArgIOStatusFlagsPayload>(base.bmp8, ()=>({
                name: `I/O Status flags`,
                description: ``,
                id: 0x006F,
                report: true,
                read: true,
                write: false,
                require: false,
                bits: { 
                0: `In Alarm`, 
                1: `Fault`, 
                2: `Overidden`, 
                3: `Out of Service`,  },
                
            })),
            IOUnitType: makeType<ZigBee.IGeneral.IArgIOUnitType, ZigBee.IGeneral.IArgIOUnitTypePayload>(base.EngineeringUnit, ()=>({
                name: `I/O Unit Type`,
                description: ``,
                id: 0x0075,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            IdentifyEffect: makeType<ZigBee.IGeneral.IArgIdentifyEffect, ZigBee.IGeneral.IArgIdentifyEffectPayload>(base.enum8, ()=>({
                name: `Identify Effect`,
                description: `The effect identifier field specifies the identify effect to use.`,
                values: { 
                0x00: `Blink`, 
                0x01: `Breathe`, 
                0x02: `Okay`, 
                0x0B: `Channel change`, 
                0xFE: `Finish`, 
                0xFF: `Stop`,  },
                
            })),
            IdentifyEffectVariant: makeType<ZigBee.IGeneral.IArgIdentifyEffectVariant, ZigBee.IGeneral.IArgIdentifyEffectVariantPayload>(base.enum8, ()=>({
                name: `Identify Effect variant`,
                description: `The effect identifier field specifies the identify effect to use.`,
                values: { 
                0x00: `Default`,  },
                
            })),
            IdentifyTime: makeType<ZigBee.IGeneral.IArgIdentifyTime, ZigBee.IGeneral.IArgIdentifyTimePayload>(base.u16, ()=>({
                name: `Identify Time`,
                description: `The time in seconds for which a device will stay in identify mode.`,
                id: 0x0000,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Seconds,
                
            })),
            IdentifyTimeout: makeType<ZigBee.IGeneral.IArgIdentifyTimeout, ZigBee.IGeneral.IArgIdentifyTimeoutPayload>(base.u16, ()=>({
                name: `Identify Timeout`,
                description: `The time in seconds for which a device will stay in identify mode.`,
                unit: units.Seconds,
                
            })),
            IkeaRemoteDirection: makeType<ZigBee.IGeneral.IArgIkeaRemoteDirection, ZigBee.IGeneral.IArgIkeaRemoteDirectionPayload>(base.enum8, ()=>({
                name: `Ikea Remote Direction`,
                description: ``,
                values: { 
                0x00: `Right`, 
                0x01: `Left`,  },
                
            })),
            JoinIndication: makeType<ZigBee.IGeneral.IArgJoinIndication, ZigBee.IGeneral.IArgJoinIndicationPayload>(base.u16, ()=>({
                name: `Join Indication`,
                description: ``,
                id: 0x0110,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            LastMessageLqi: makeType<ZigBee.IGeneral.IArgLastMessageLqi, ZigBee.IGeneral.IArgLastMessageLqiPayload>(base.u8, ()=>({
                name: `Last Message LQI`,
                description: ``,
                id: 0x011C,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            LastMessageRssi: makeType<ZigBee.IGeneral.IArgLastMessageRssi, ZigBee.IGeneral.IArgLastMessageRssiPayload>(base.s8, ()=>({
                name: `Last Message RSSI`,
                description: ``,
                id: 0x011D,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            LastSetTime: makeType<ZigBee.IGeneral.IArgLastSetTime, ZigBee.IGeneral.IArgLastSetTimePayload>(base.utc, ()=>({
                name: `Last Set Time`,
                description: ``,
                id: 0x0008,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            LedIndication: makeType<ZigBee.IGeneral.IArgLedIndication, ZigBee.IGeneral.IArgLedIndicationPayload>(base.bool, ()=>({
                name: `LED Indication`,
                description: ``,
                id: 0x0033,
                report: false,
                read: true,
                write: true,
                require: false,
                mnf: 0x100b,
                
            })),
            Level: makeType<ZigBee.IGeneral.IArgLevel, ZigBee.IGeneral.IArgLevelPayload>(base.u8, ()=>({
                name: `Level`,
                description: ``,
                unit: units.Percent,
                scale: 2.54,
                
            })),
            LevelControlOptions: makeType<ZigBee.IGeneral.IArgLevelControlOptions, ZigBee.IGeneral.IArgLevelControlOptionsPayload>(base.bmp8, ()=>({
                name: `Level Control Options`,
                description: `is a bitmap that determines the default behavior of some cluster commands`,
                id: 0x000F,
                report: false,
                read: true,
                write: true,
                require: false,
                bits: { 
                0x00: `Execute if off`,  },
                
            })),
            LevelDirection: makeType<ZigBee.IGeneral.IArgLevelDirection, ZigBee.IGeneral.IArgLevelDirectionPayload>(base.enum8, ()=>({
                name: `Level direction`,
                description: ``,
                values: { 
                0x00: `Up`, 
                0x01: `Down`,  },
                
            })),
            LocalTime: makeType<ZigBee.IGeneral.IArgLocalTime, ZigBee.IGeneral.IArgLocalTimePayload>(base.u32, ()=>({
                name: `Local Time`,
                description: `Local time`,
                id: 0x0007,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            LocationAge: makeType<ZigBee.IGeneral.IArgLocationAge, ZigBee.IGeneral.IArgLocationAgePayload>(base.u16, ()=>({
                name: `Location Age`,
                description: ``,
                id: 0x0002,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.Seconds,
                
            })),
            LocationDescription: makeType<ZigBee.IGeneral.IArgLocationDescription, ZigBee.IGeneral.IArgLocationDescriptionPayload>(base.cstring, ()=>({
                name: `Location Description`,
                description: ``,
                id: 0x0010,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            LocationFlags: makeType<ZigBee.IGeneral.IArgLocationFlags, ZigBee.IGeneral.IArgLocationFlagsPayload>(base.bmp8, ()=>({
                name: `Location flags`,
                description: ``,
                bits: { 
                0: `Absolute Only`, 
                1: `Recalculate`, 
                2: `Broadcast Indicator`, 
                3: `Broadcast Response`, 
                4: `Compact Response`,  },
                
            })),
            LocationMethod: makeType<ZigBee.IGeneral.IArgLocationMethod, ZigBee.IGeneral.IArgLocationMethodPayload>(base.enum8, ()=>({
                name: `Location Method`,
                description: ``,
                id: 0x0001,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x00: `Lateration`, 
                0x01: `Signposting`, 
                0x02: `RF fingerprinting`, 
                0x03: `Out of band`, 
                0x04: `Centralized`,  },
                
            })),
            LocationType: makeType<ZigBee.IGeneral.IArgLocationType, ZigBee.IGeneral.IArgLocationTypePayload>(base.enum8, ()=>({
                name: `Location Type`,
                description: ``,
                id: 0x0000,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x00: `3D Location`, 
                0x01: `Absolute 3D Location`, 
                0x02: `2D Location`, 
                0x03: `Absolute 2D Location`,  },
                
            })),
            LongPollInterval: makeType<ZigBee.IGeneral.IArgLongPollInterval, ZigBee.IGeneral.IArgLongPollIntervalPayload>(base.u32, ()=>({
                name: `Long Poll Interval`,
                description: ``,
                id: 0x0001,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            LongPollIntervalMin: makeType<ZigBee.IGeneral.IArgLongPollIntervalMin, ZigBee.IGeneral.IArgLongPollIntervalMinPayload>(base.u32, ()=>({
                name: `Long Poll Interval Min`,
                description: ``,
                id: 0x0005,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            LowTempDwellTripPoint: makeType<ZigBee.IGeneral.IArgLowTempDwellTripPoint, ZigBee.IGeneral.IArgLowTempDwellTripPointPayload>(base.u24, ()=>({
                name: `Low Temp Dwell Trip Point`,
                description: ``,
                id: 0x0013,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Seconds,
                
            })),
            LowTempThreshold: makeType<ZigBee.IGeneral.IArgLowTempThreshold, ZigBee.IGeneral.IArgLowTempThresholdPayload>(base.s16, ()=>({
                name: `Low Temp Threshold`,
                description: `If the current temperature drops below the threshold for longer
than the time specified by low temp dwell, an alarm will be triggered`,
                id: 0x0011,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.DegreesCelsius,
                
            })),
            MacRxBcast: makeType<ZigBee.IGeneral.IArgMacRxBcast, ZigBee.IGeneral.IArgMacRxBcastPayload>(base.u32, ()=>({
                name: `Mac Rx Bcast`,
                description: ``,
                id: 0x0100,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            MacRxUcast: makeType<ZigBee.IGeneral.IArgMacRxUcast, ZigBee.IGeneral.IArgMacRxUcastPayload>(base.u32, ()=>({
                name: `Mac Rx Ucast`,
                description: ``,
                id: 0x0102,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            MacTxBcast: makeType<ZigBee.IGeneral.IArgMacTxBcast, ZigBee.IGeneral.IArgMacTxBcastPayload>(base.u32, ()=>({
                name: `Mac Tx Bcast`,
                description: ``,
                id: 0x0101,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            MacTxUcast: makeType<ZigBee.IGeneral.IArgMacTxUcast, ZigBee.IGeneral.IArgMacTxUcastPayload>(base.u32, ()=>({
                name: `Mac Tx Ucast`,
                description: ``,
                id: 0x0103,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            MacTxUcastFail: makeType<ZigBee.IGeneral.IArgMacTxUcastFail, ZigBee.IGeneral.IArgMacTxUcastFailPayload>(base.u16, ()=>({
                name: `Mac Tx Ucast Fail`,
                description: ``,
                id: 0x0105,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            MacTxUcastRetry: makeType<ZigBee.IGeneral.IArgMacTxUcastRetry, ZigBee.IGeneral.IArgMacTxUcastRetryPayload>(base.u16, ()=>({
                name: `Mac Tx Ucast Retry`,
                description: ``,
                id: 0x0104,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            MainsAlarmMask: makeType<ZigBee.IGeneral.IArgMainsAlarmMask, ZigBee.IGeneral.IArgMainsAlarmMaskPayload>(base.bmp8, ()=>({
                name: `Mains Alarm Mask`,
                description: ``,
                id: 0x0010,
                report: false,
                read: true,
                write: true,
                require: false,
                bits: { 
                0: `Mains Voltage too low`, 
                1: `Mains Voltage too high`, 
                2: `Mains power supply lost/unavailable`,  },
                
            })),
            MainsFrequency: makeType<ZigBee.IGeneral.IArgMainsFrequency, ZigBee.IGeneral.IArgMainsFrequencyPayload>(base.u8, ()=>({
                name: `Mains Frequency`,
                description: `Resolution of 2Hz
Special values:
* 0x00 frequency too low to measure
* 0xfe frequency too high to measure
* 0xff unable to measure`,
                id: 0x0001,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.Hertz,
                scale: 2,
                
            })),
            MainsVoltage: makeType<ZigBee.IGeneral.IArgMainsVoltage, ZigBee.IGeneral.IArgMainsVoltagePayload>(base.u16, ()=>({
                name: `Mains Voltage`,
                description: ``,
                id: 0x0000,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.Volts,
                scale: 10,
                
            })),
            MainsVoltageDwellTripPoint: makeType<ZigBee.IGeneral.IArgMainsVoltageDwellTripPoint, ZigBee.IGeneral.IArgMainsVoltageDwellTripPointPayload>(base.u16, ()=>({
                name: `Mains Voltage Dwell Trip Point`,
                description: `Length of time that the value of MainsVoltage MAY exist beyond either
of its thresholds before an alarm is generated`,
                id: 0x0013,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Seconds,
                
            })),
            MainsVoltageMaxThreshold: makeType<ZigBee.IGeneral.IArgMainsVoltageMaxThreshold, ZigBee.IGeneral.IArgMainsVoltageMaxThresholdPayload>(base.u16, ()=>({
                name: `Mains Voltage Max Threshold`,
                description: ``,
                id: 0x0012,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Volts,
                scale: 10,
                
            })),
            MainsVoltageMinThreshold: makeType<ZigBee.IGeneral.IArgMainsVoltageMinThreshold, ZigBee.IGeneral.IArgMainsVoltageMinThresholdPayload>(base.u16, ()=>({
                name: `Mains Voltage Min Threshold`,
                description: ``,
                id: 0x0011,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Volts,
                scale: 10,
                
            })),
            ManufacturerName: makeType<ZigBee.IGeneral.IArgManufacturerName, ZigBee.IGeneral.IArgManufacturerNamePayload>(base.cstring, ()=>({
                name: `Manufacturer Name`,
                description: ``,
                id: 0x0004,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            MaxTempExperienced: makeType<ZigBee.IGeneral.IArgMaxTempExperienced, ZigBee.IGeneral.IArgMaxTempExperiencedPayload>(base.s16, ()=>({
                name: `Max Temp Experienced`,
                description: ``,
                id: 0x0002,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.DegreesCelsius,
                
            })),
            MinTempExperienced: makeType<ZigBee.IGeneral.IArgMinTempExperienced, ZigBee.IGeneral.IArgMinTempExperiencedPayload>(base.s16, ()=>({
                name: `Min Temp Experienced`,
                description: ``,
                id: 0x0001,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.DegreesCelsius,
                
            })),
            ModelIdentifier: makeType<ZigBee.IGeneral.IArgModelIdentifier, ZigBee.IGeneral.IArgModelIdentifierPayload>(base.cstring, ()=>({
                name: `Model Identifier`,
                description: ``,
                id: 0x0005,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            MultistateNumberOfStates: makeType<ZigBee.IGeneral.IArgMultistateNumberOfStates, ZigBee.IGeneral.IArgMultistateNumberOfStatesPayload>(base.u16, ()=>({
                name: `Multistate Number of States`,
                description: ``,
                id: 0x004A,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            MultistatePresentValue: makeType<ZigBee.IGeneral.IArgMultistatePresentValue, ZigBee.IGeneral.IArgMultistatePresentValuePayload>(base.u16, ()=>({
                name: `Multistate Present value`,
                description: ``,
                id: 0x0055,
                report: true,
                read: true,
                write: true,
                require: false,
                
            })),
            MultistatePriority: makeType<ZigBee.IGeneral.IArgMultistatePriority, ZigBee.IGeneral.IArgMultistatePriorityPayload>(base.local, ()=>({
                name: `Multistate Priority`,
                description: ``,
                payload: { 
                 },
                
            })),
            MultistatePriorityArray: makeType<ZigBee.IGeneral.IArgMultistatePriorityArray, ZigBee.IGeneral.IArgMultistatePriorityArrayPayload>(base.array, ()=>({
                name: `Multistate Priority Array`,
                description: ``,
                id: 0x0057,
                report: false,
                read: true,
                write: true,
                require: false,
                arrayType: ZigBee.General.Types.MultistatePriority,
                
            })),
            MultistateRelinquishDefault: makeType<ZigBee.IGeneral.IArgMultistateRelinquishDefault, ZigBee.IGeneral.IArgMultistateRelinquishDefaultPayload>(base.u16, ()=>({
                name: `Multistate Relinquish Default`,
                description: ``,
                id: 0x0068,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            MultistateText: makeType<ZigBee.IGeneral.IArgMultistateText, ZigBee.IGeneral.IArgMultistateTextPayload>(base.cstring, ()=>({
                name: `Multistate Text`,
                description: ``,
                id: 0x000E,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            NeighborAdded: makeType<ZigBee.IGeneral.IArgNeighborAdded, ZigBee.IGeneral.IArgNeighborAddedPayload>(base.u16, ()=>({
                name: `Neighbor Added`,
                description: ``,
                id: 0x010D,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            NeighborInfo: makeType<ZigBee.IGeneral.IArgNeighborInfo, ZigBee.IGeneral.IArgNeighborInfoPayload>(base.local, ()=>({
                name: `Neighbor Info`,
                description: ``,
                payload: { 
                    Device: ZigBee.General.Types.Device,
                    XCoordinate: ZigBee.General.Types.XCoordinate,
                    YCoordinate: ZigBee.General.Types.YCoordinate,
                    ZCoordinate: ZigBee.General.Types.ZCoordinate,
                    Rssi: ZigBee.General.Types.Rssi,
                    NumberRssiMeasurements: ZigBee.General.Types.NumberRssiMeasurements,
                 },
                
            })),
            NeighborRemoved: makeType<ZigBee.IGeneral.IArgNeighborRemoved, ZigBee.IGeneral.IArgNeighborRemovedPayload>(base.u16, ()=>({
                name: `Neighbor Removed`,
                description: ``,
                id: 0x010E,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            NeighborStale: makeType<ZigBee.IGeneral.IArgNeighborStale, ZigBee.IGeneral.IArgNeighborStalePayload>(base.u16, ()=>({
                name: `Neighbor Stale`,
                description: ``,
                id: 0x010F,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            NeighborsInfoList: makeType<ZigBee.IGeneral.IArgNeighborsInfoList, ZigBee.IGeneral.IArgNeighborsInfoListPayload>(base.set, ()=>({
                name: `Neighbors Info List`,
                description: ``,
                arrayType: ZigBee.General.Types.NeighborInfo,
                
            })),
            NumberOfDevices: makeType<ZigBee.IGeneral.IArgNumberOfDevices, ZigBee.IGeneral.IArgNumberOfDevicesPayload>(base.u8, ()=>({
                name: `Number of Devices`,
                description: ``,
                id: 0x0004,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            NumberOfResets: makeType<ZigBee.IGeneral.IArgNumberOfResets, ZigBee.IGeneral.IArgNumberOfResetsPayload>(base.u16, ()=>({
                name: `Number of Resets`,
                description: ``,
                id: 0x0000,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            NumberResponses: makeType<ZigBee.IGeneral.IArgNumberResponses, ZigBee.IGeneral.IArgNumberResponsesPayload>(base.u8, ()=>({
                name: `Number Responses`,
                description: ``,
                
            })),
            NumberRssiMeasurements: makeType<ZigBee.IGeneral.IArgNumberRssiMeasurements, ZigBee.IGeneral.IArgNumberRssiMeasurementsPayload>(base.u8, ()=>({
                name: `Number RSSI Measurements`,
                description: `is the number of measurements to use to generate one location estimate`,
                id: 0x0017,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            NwkDecryptFailures: makeType<ZigBee.IGeneral.IArgNwkDecryptFailures, ZigBee.IGeneral.IArgNwkDecryptFailuresPayload>(base.u16, ()=>({
                name: `NWK Decrypt Failures`,
                description: ``,
                id: 0x0115,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            NwkFcFailure: makeType<ZigBee.IGeneral.IArgNwkFcFailure, ZigBee.IGeneral.IArgNwkFcFailurePayload>(base.u16, ()=>({
                name: `NWK FC Failure`,
                description: ``,
                id: 0x0112,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            OffTransitionTime: makeType<ZigBee.IGeneral.IArgOffTransitionTime, ZigBee.IGeneral.IArgOffTransitionTimePayload>(base.u16, ()=>({
                name: `Off Transition Time`,
                description: ``,
                id: 0x0013,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Seconds,
                scale: 10,
                
            })),
            OffWaitTime: makeType<ZigBee.IGeneral.IArgOffWaitTime, ZigBee.IGeneral.IArgOffWaitTimePayload>(base.u16, ()=>({
                name: `Off Wait Time`,
                description: ``,
                id: 0x4002,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.Seconds,
                scale: 10,
                
            })),
            OnLevel: makeType<ZigBee.IGeneral.IArgOnLevel, ZigBee.IGeneral.IArgOnLevelPayload>(base.u8, ()=>({
                name: `On Level`,
                description: `determines the value that the CurrentLevel attribute is set to when the OnOff attribute of an On/Off cluster
on the same endpoint is set to On. If the OnLevel attribute is not implemented, or is set to 0xff, it has no
effect.`,
                id: 0x0011,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Percent,
                scale: 2.54,
                values: { 
                0xFF: `Previous`,  },
                
            })),
            OnOff: makeType<ZigBee.IGeneral.IArgOnOff, ZigBee.IGeneral.IArgOnOffPayload>(base.bool, ()=>({
                name: `On Off`,
                description: ``,
                id: 0x0000,
                report: true,
                read: true,
                write: false,
                require: false,
                values: { 
                0x00: `Off`, 
                0x01: `On`,  },
                
            })),
            OnOffControl: makeType<ZigBee.IGeneral.IArgOnOffControl, ZigBee.IGeneral.IArgOnOffControlPayload>(base.bmp8, ()=>({
                name: `On/off control`,
                description: ``,
                bits: { 
                0: `Accept only when on`,  },
                
            })),
            OnOffTransistionTime: makeType<ZigBee.IGeneral.IArgOnOffTransistionTime, ZigBee.IGeneral.IArgOnOffTransistionTimePayload>(base.u16, ()=>({
                name: `On/Off Transistion Time`,
                description: `represents the time taken to move to or from the target level when On of Off commands are received
by an On/Off cluster on the same endpoint.
The actual time taken should be as close to OnOffTransitionTime as the device is able.`,
                id: 0x0010,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Seconds,
                scale: 10,
                
            })),
            OnTime: makeType<ZigBee.IGeneral.IArgOnTime, ZigBee.IGeneral.IArgOnTimePayload>(base.u16, ()=>({
                name: `On Time`,
                description: ``,
                id: 0x4001,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.Seconds,
                scale: 10,
                
            })),
            OnTransitionTime: makeType<ZigBee.IGeneral.IArgOnTransitionTime, ZigBee.IGeneral.IArgOnTransitionTimePayload>(base.u16, ()=>({
                name: `On Transition Time`,
                description: ``,
                id: 0x0012,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Seconds,
                scale: 10,
                
            })),
            OverTempTotalDwell: makeType<ZigBee.IGeneral.IArgOverTempTotalDwell, ZigBee.IGeneral.IArgOverTempTotalDwellPayload>(base.u16, ()=>({
                name: `Over Temp Total Dwell`,
                description: `Total time the device has spent above the tmperature specified by High Temp Threshold`,
                id: 0x0003,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            PacketBufferAllocFailures: makeType<ZigBee.IGeneral.IArgPacketBufferAllocFailures, ZigBee.IGeneral.IArgPacketBufferAllocFailuresPayload>(base.u16, ()=>({
                name: `Packet Buffer Alloc Failures`,
                description: ``,
                id: 0x0117,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            PacketValidateDropcount: makeType<ZigBee.IGeneral.IArgPacketValidateDropcount, ZigBee.IGeneral.IArgPacketValidateDropcountPayload>(base.u16, ()=>({
                name: `Packet Validate Dropcount`,
                description: ``,
                id: 0x011A,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            PathLossExponent: makeType<ZigBee.IGeneral.IArgPathLossExponent, ZigBee.IGeneral.IArgPathLossExponentPayload>(base.u16, ()=>({
                name: `Path loss Exponent`,
                description: `is the rate at which the signal power decays with increasing distance`,
                id: 0x0014,
                report: false,
                read: true,
                write: true,
                require: false,
                scale: 0.01,
                
            })),
            PersistensMemoryWrites: makeType<ZigBee.IGeneral.IArgPersistensMemoryWrites, ZigBee.IGeneral.IArgPersistensMemoryWritesPayload>(base.u16, ()=>({
                name: `Persistens Memory Writes`,
                description: ``,
                id: 0x0001,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            PhyToMacQueueLimitReached: makeType<ZigBee.IGeneral.IArgPhyToMacQueueLimitReached, ZigBee.IGeneral.IArgPhyToMacQueueLimitReachedPayload>(base.u16, ()=>({
                name: `Phy to MAC queue limit reached`,
                description: ``,
                id: 0x0119,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            PhysicalEnvironment: makeType<ZigBee.IGeneral.IArgPhysicalEnvironment, ZigBee.IGeneral.IArgPhysicalEnvironmentPayload>(base.enum8, ()=>({
                name: `Physical Environment`,
                description: ``,
                id: 0x0011,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x00: `Unspecified Environment`, 
                0x01: `Atrium`, 
                0x02: `Bar`, 
                0x03: `Courtyard`, 
                0x04: `Bathroom`, 
                0x05: `Bedroom`, 
                0x06: `Billiard Room`, 
                0x07: `Utility Room`, 
                0x08: `Cellar`, 
                0x09: `Storage Closet`, 
                0x0a: `Theater`, 
                0x0b: `Office`, 
                0x0c: `Deck`, 
                0x0d: `Den`, 
                0x0e: `Dining Room`, 
                0x0f: `Electrical Room`, 
                0x10: `Elevator`, 
                0x11: `Entry`, 
                0x12: `Family Room`, 
                0x13: `Main Floor`, 
                0x14: `Upstairs`, 
                0x15: `Downstairs`, 
                0x16: `Basement/Lower Level`, 
                0x17: `Gallery`, 
                0x18: `Game Room`, 
                0x19: `Garage`, 
                0x1a: `Gym`, 
                0x1b: `Hallway`, 
                0x1c: `House`, 
                0x1d: `Kitchen`, 
                0x1e: `Laundry Room`, 
                0x1f: `Library`, 
                0x20: `Master Bedroom`, 
                0x21: `Mud Room (small room for coats and boots)`, 
                0x22: `Nursery`, 
                0x23: `Pantry`, 
                0x24: `Office 2`, 
                0x25: `Outside`, 
                0x26: `Pool`, 
                0x27: `Porch`, 
                0x28: `Sewing Room`, 
                0x29: `Sitting Room`, 
                0x2a: `Stairway`, 
                0x2b: `Yard`, 
                0x2c: `Attic`, 
                0x2d: `Hot Tub`, 
                0x2e: `Living Room`, 
                0x2f: `Sauna`, 
                0x30: `Shop/Workshop`, 
                0x31: `Guest Bedroom`, 
                0x32: `Guest Bath`, 
                0x33: `Powder Room (1/2 bath)`, 
                0x34: `Back Yard`, 
                0x35: `Front Yard`, 
                0x36: `Patio`, 
                0x37: `Driveway`, 
                0x38: `Sun Room`, 
                0x39: `Living Room 2`, 
                0x3a: `Spa`, 
                0x3b: `Whirlpool`, 
                0x3c: `Shed`, 
                0x3d: `Equipment Storage`, 
                0x3e: `Hobby/Craft Room`, 
                0x3f: `Fountain`, 
                0x40: `Pond`, 
                0x41: `Reception Room`, 
                0x42: `Breakfast Room`, 
                0x43: `Nook`, 
                0x44: `Garden`, 
                0x45: `Balcony`, 
                0x46: `Panic Room`, 
                0x47: `Terrace`, 
                0x48: `Roof`, 
                0x49: `Toilet`, 
                0x4a: `Toilet Main`, 
                0x4b: `Outside Toilet`, 
                0x4c: `Shower room`, 
                0x4d: `Study`, 
                0x4e: `Front Garden`, 
                0x4f: `Back Garden`, 
                0x50: `Kettle`, 
                0x51: `Television`, 
                0x52: `Stove`, 
                0x53: `Microwave`, 
                0x54: `Toaster`, 
                0x55: `Vacuum`, 
                0x56: `Appliance`, 
                0x57: `Front Door`, 
                0x58: `Back Door`, 
                0x59: `Fridge Door`, 
                0x60: `Medication Cabinet Door`, 
                0x61: `Wardrobe Door`, 
                0x62: `Front Cupboard Door`, 
                0x63: `Other Door`, 
                0x64: `Waiting Room`, 
                0x65: `Triage Room`, 
                0x66: `Doctor’s Office`, 
                0x67: `Patient’s Private Room`, 
                0x68: `Consultation Room`, 
                0x69: `Nurse Station`, 
                0x6a: `Ward`, 
                0x6b: `Corridor`, 
                0x6c: `Operating Theatre`, 
                0x6d: `Dental Surgery Room`, 
                0x6e: `Medical Imaging Room`, 
                0x6f: `Decontamination Room`, 
                0xFF: `Unknown Environment`,  },
                
            })),
            Power: makeType<ZigBee.IGeneral.IArgPower, ZigBee.IGeneral.IArgPowerPayload>(base.s16, ()=>({
                name: `Power`,
                description: ``,
                id: 0x0013,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.DecibelMilliWatts,
                scale: 100,
                
            })),
            PowerOnLevel: makeType<ZigBee.IGeneral.IArgPowerOnLevel, ZigBee.IGeneral.IArgPowerOnLevelPayload>(base.u8, ()=>({
                name: `Power On level`,
                description: ``,
                id: 0x4000,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Percent,
                scale: 2.54,
                values: { 
                0xFF: `Previous`,  },
                
            })),
            PowerSource: makeType<ZigBee.IGeneral.IArgPowerSource, ZigBee.IGeneral.IArgPowerSourcePayload>(base.enum8, ()=>({
                name: `Power Source`,
                description: ``,
                id: 0x0007,
                report: false,
                read: true,
                write: false,
                require: false,
                values: { 
                0x00: `Unknown`, 
                0x01: `Mains (single phase)`, 
                0x02: `Mains (3 phase)`, 
                0x03: `Battery`, 
                0x04: `DC Source`, 
                0x05: `Emergency mains constantly powered`, 
                0x06: `Emergency mains and transfer switch`, 
                0x80: `Unknown with battery backup`, 
                0x81: `Mains (single phase) with battery backup`, 
                0x82: `Mains (3 phase) with battery backup`, 
                0x83: `Battery with battery backup`, 
                0x84: `DC Source with battery backup`, 
                0x85: `Emergency mains constantly powered with battery backup`, 
                0x86: `Emergency mains and transfer switch with battery backup`,  },
                
            })),
            PoweronOnOff: makeType<ZigBee.IGeneral.IArgPoweronOnOff, ZigBee.IGeneral.IArgPoweronOnOffPayload>(base.enum8, ()=>({
                name: `PowerOn On/Off`,
                description: ``,
                id: 0x4003,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x00: `Off`, 
                0x01: `On`, 
                0xFF: `Previous`,  },
                
            })),
            ProductCode: makeType<ZigBee.IGeneral.IArgProductCode, ZigBee.IGeneral.IArgProductCodePayload>(base.ostring, ()=>({
                name: `Product code`,
                description: `As printed on the product.`,
                id: 0x000A,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            ProductUrl: makeType<ZigBee.IGeneral.IArgProductUrl, ZigBee.IGeneral.IArgProductUrlPayload>(base.cstring, ()=>({
                name: `Product URL`,
                description: `specifies a link to a web page containing specific product information.`,
                id: 0x000B,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            QualityIndex: makeType<ZigBee.IGeneral.IArgQualityIndex, ZigBee.IGeneral.IArgQualityIndexPayload>(base.u16, ()=>({
                name: `Quality index`,
                description: ``,
                
            })),
            QualityMeasure: makeType<ZigBee.IGeneral.IArgQualityMeasure, ZigBee.IGeneral.IArgQualityMeasurePayload>(base.u8, ()=>({
                name: `Quality Measure`,
                description: ``,
                id: 0x0003,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.Percent,
                
            })),
            Rate: makeType<ZigBee.IGeneral.IArgRate, ZigBee.IGeneral.IArgRatePayload>(base.u8, ()=>({
                name: `Rate`,
                description: ``,
                unit: units.PercentPerSecond,
                scale: 2.54,
                
            })),
            RelayedUcast: makeType<ZigBee.IGeneral.IArgRelayedUcast, ZigBee.IGeneral.IArgRelayedUcastPayload>(base.u16, ()=>({
                name: `Relayed Ucast`,
                description: ``,
                id: 0x0118,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            RemainingTime: makeType<ZigBee.IGeneral.IArgRemainingTime, ZigBee.IGeneral.IArgRemainingTimePayload>(base.u16, ()=>({
                name: `Remaining Time`,
                description: `represents the time remaining until the current command is complete. It is specified in 1/10ths of a second.`,
                id: 0x0001,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.Seconds,
                scale: 10,
                
            })),
            ReportingPeriod: makeType<ZigBee.IGeneral.IArgReportingPeriod, ZigBee.IGeneral.IArgReportingPeriodPayload>(base.u16, ()=>({
                name: `Reporting Period`,
                description: ``,
                id: 0x0015,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Seconds,
                
            })),
            Resolution: makeType<ZigBee.IGeneral.IArgResolution, ZigBee.IGeneral.IArgResolutionPayload>(base.enum8, ()=>({
                name: `Resolution`,
                description: ``,
                values: { 
                0x00: `High`, 
                0x01: `Mid`, 
                0x02: `Low`,  },
                
            })),
            RouteDiscInitiated: makeType<ZigBee.IGeneral.IArgRouteDiscInitiated, ZigBee.IGeneral.IArgRouteDiscInitiatedPayload>(base.u16, ()=>({
                name: `Route Disc Initiated`,
                description: ``,
                id: 0x010C,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            Rssi: makeType<ZigBee.IGeneral.IArgRssi, ZigBee.IGeneral.IArgRssiPayload>(base.s8, ()=>({
                name: `RSSI`,
                description: ``,
                unit: units.DecibelMilliWatts,
                
            })),
            SceneCapacity: makeType<ZigBee.IGeneral.IArgSceneCapacity, ZigBee.IGeneral.IArgSceneCapacityPayload>(base.u8, ()=>({
                name: `Scene Capacity`,
                description: `specifies remaining number of scenes that can be added`,
                
            })),
            SceneCopyMode: makeType<ZigBee.IGeneral.IArgSceneCopyMode, ZigBee.IGeneral.IArgSceneCopyModePayload>(base.bmp8, ()=>({
                name: `Scene Copy Mode`,
                description: ``,
                bits: { 
                0: `Copy All Scenes`,  },
                
            })),
            SceneCount: makeType<ZigBee.IGeneral.IArgSceneCount, ZigBee.IGeneral.IArgSceneCountPayload>(base.u8, ()=>({
                name: `Scene Count`,
                description: ``,
                id: 0x0000,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            SceneExtensions: makeType<ZigBee.IGeneral.IArgSceneExtensions, ZigBee.IGeneral.IArgSceneExtensionsPayload>(base.SceneExtensionSet, ()=>({
                name: `Scene Extensions`,
                description: `The format of each extension field set is a 16 bit field carrying the cluster ID,
followed by an 8 bit length field and the set of scene extension fields specified
in  the  relevant  cluster. The length field holds the length in octets of that
extension field set. Extension field set format:
{{clusterId1, length 1, {extension field set 1}}, {clusterId2, length 2, {extension field set 2}} ...}
I.e. the field would be a repeating struct with [ClusterID uint16] [OctetLength uint8] [AttrID ...uint16]`,
                
            })),
            SceneId: makeType<ZigBee.IGeneral.IArgSceneId, ZigBee.IGeneral.IArgSceneIdPayload>(base.u8, ()=>({
                name: `Scene ID`,
                description: ``,
                
            })),
            SceneLastConfiguredBy: makeType<ZigBee.IGeneral.IArgSceneLastConfiguredBy, ZigBee.IGeneral.IArgSceneLastConfiguredByPayload>(base.uid, ()=>({
                name: `Scene Last Configured By`,
                description: ``,
                id: 0x0005,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            SceneList: makeType<ZigBee.IGeneral.IArgSceneList, ZigBee.IGeneral.IArgSceneListPayload>(base.set, ()=>({
                name: `Scene list`,
                description: ``,
                arrayType: base.u8,
                
            })),
            SceneName: makeType<ZigBee.IGeneral.IArgSceneName, ZigBee.IGeneral.IArgSceneNamePayload>(base.cstring, ()=>({
                name: `Scene Name`,
                description: ``,
                
            })),
            SceneNameSupport: makeType<ZigBee.IGeneral.IArgSceneNameSupport, ZigBee.IGeneral.IArgSceneNameSupportPayload>(base.bmp8, ()=>({
                name: `Scene Name Support`,
                description: ``,
                id: 0x0004,
                report: false,
                read: true,
                write: false,
                require: false,
                bits: { 
                7: `Names Supported`,  },
                
            })),
            SceneValid: makeType<ZigBee.IGeneral.IArgSceneValid, ZigBee.IGeneral.IArgSceneValidPayload>(base.bool, ()=>({
                name: `Scene Valid`,
                description: ``,
                id: 0x0003,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            Sensitivity: makeType<ZigBee.IGeneral.IArgSensitivity, ZigBee.IGeneral.IArgSensitivityPayload>(base.enum8, ()=>({
                name: `Sensitivity`,
                description: ``,
                id: 0x0030,
                report: false,
                read: true,
                write: true,
                require: false,
                mnf: 0x100b,
                values: { 
                0x00: `default`, 
                0x01: `high`, 
                0x02: `max`,  },
                
            })),
            ShortPollInterval: makeType<ZigBee.IGeneral.IArgShortPollInterval, ZigBee.IGeneral.IArgShortPollIntervalPayload>(base.u16, ()=>({
                name: `Short Poll Interval`,
                description: ``,
                id: 0x0002,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            StackVersion: makeType<ZigBee.IGeneral.IArgStackVersion, ZigBee.IGeneral.IArgStackVersionPayload>(base.u8, ()=>({
                name: `Stack Version`,
                description: ``,
                id: 0x0002,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            StandardTime: makeType<ZigBee.IGeneral.IArgStandardTime, ZigBee.IGeneral.IArgStandardTimePayload>(base.u32, ()=>({
                name: `Standard Time`,
                description: `Local time (without DST offset)`,
                id: 0x0006,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            Status: makeType<ZigBee.IGeneral.IArgStatus, ZigBee.IGeneral.IArgStatusPayload>(base.Status, ()=>({
                name: `Status`,
                description: ``,
                
            })),
            StepSize: makeType<ZigBee.IGeneral.IArgStepSize, ZigBee.IGeneral.IArgStepSizePayload>(base.u8, ()=>({
                name: `Step size`,
                description: ``,
                unit: units.Percent,
                scale: 2.54,
                
            })),
            SwBuildId: makeType<ZigBee.IGeneral.IArgSwBuildId, ZigBee.IGeneral.IArgSwBuildIdPayload>(base.cstring, ()=>({
                name: `SW Build ID`,
                description: ``,
                id: 0x4000,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            SwitchActions: makeType<ZigBee.IGeneral.IArgSwitchActions, ZigBee.IGeneral.IArgSwitchActionsPayload>(base.enum8, ()=>({
                name: `Switch actions`,
                description: `specifies the commands of the On/Off cluster to be generated when the switch moves between its two states.`,
                id: 0x0010,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x00: `On-Off`, 
                0x01: `Off-On`, 
                0x02: `Toggle`,  },
                
            })),
            SwitchType: makeType<ZigBee.IGeneral.IArgSwitchType, ZigBee.IGeneral.IArgSwitchTypePayload>(base.enum8, ()=>({
                name: `Switch type`,
                description: `specifies the basic functionality of the On/Off switching device.`,
                id: 0x0000,
                report: false,
                read: true,
                write: false,
                require: false,
                values: { 
                0x00: `Toggle`, 
                0x01: `Momentary`, 
                0x02: `Multifunction`,  },
                
            })),
            Time: makeType<ZigBee.IGeneral.IArgTime, ZigBee.IGeneral.IArgTimePayload>(base.utc, ()=>({
                name: `Time`,
                description: ``,
                id: 0x0000,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            TimeStatus: makeType<ZigBee.IGeneral.IArgTimeStatus, ZigBee.IGeneral.IArgTimeStatusPayload>(base.bmp8, ()=>({
                name: `Time Status`,
                description: ``,
                id: 0x0001,
                report: false,
                read: true,
                write: true,
                require: false,
                bits: { 
                0: `Master Clock`, 
                1: `Synchronized`, 
                2: `Master for Timezone and Dst`, 
                3: `Superseding`,  },
                
            })),
            TimeZone: makeType<ZigBee.IGeneral.IArgTimeZone, ZigBee.IGeneral.IArgTimeZonePayload>(base.s32, ()=>({
                name: `Time Zone`,
                description: `Offset during normal time from UTC in seconds`,
                id: 0x0002,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Seconds,
                
            })),
            TransitionTime: makeType<ZigBee.IGeneral.IArgTransitionTime, ZigBee.IGeneral.IArgTransitionTimePayload>(base.u16, ()=>({
                name: `Transition Time`,
                description: ``,
                unit: units.Seconds,
                scale: 10,
                
            })),
            TransitionTimeSec: makeType<ZigBee.IGeneral.IArgTransitionTimeSec, ZigBee.IGeneral.IArgTransitionTimeSecPayload>(base.u16, ()=>({
                name: `Transition time (Sec)`,
                description: ``,
                unit: units.Seconds,
                scale: 10,
                
            })),
            UserTest: makeType<ZigBee.IGeneral.IArgUserTest, ZigBee.IGeneral.IArgUserTestPayload>(base.bool, ()=>({
                name: `User test`,
                description: ``,
                id: 0x0032,
                report: false,
                read: true,
                write: true,
                require: false,
                mnf: 0x100b,
                
            })),
            ValidUntilTime: makeType<ZigBee.IGeneral.IArgValidUntilTime, ZigBee.IGeneral.IArgValidUntilTimePayload>(base.utc, ()=>({
                name: `Valid Until Time`,
                description: ``,
                id: 0x0009,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            XCoordinate: makeType<ZigBee.IGeneral.IArgXCoordinate, ZigBee.IGeneral.IArgXCoordinatePayload>(base.s16, ()=>({
                name: `X Coordinate`,
                description: ``,
                id: 0x0010,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Meters,
                scale: 10,
                
            })),
            YCoordinate: makeType<ZigBee.IGeneral.IArgYCoordinate, ZigBee.IGeneral.IArgYCoordinatePayload>(base.s16, ()=>({
                name: `Y Coordinate`,
                description: ``,
                id: 0x0011,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Meters,
                scale: 10,
                
            })),
            ZCoordinate: makeType<ZigBee.IGeneral.IArgZCoordinate, ZigBee.IGeneral.IArgZCoordinatePayload>(base.s16, ()=>({
                name: `Z Coordinate`,
                description: ``,
                id: 0x0012,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Meters,
                scale: 10,
                
            })),
            ZclVersion: makeType<ZigBee.IGeneral.IArgZclVersion, ZigBee.IGeneral.IArgZclVersionPayload>(base.u8, ()=>({
                name: `ZCL Version`,
                description: ``,
                id: 0x0000,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })), },
        Basic: {
            ID: 0x0000,
            Name: `Basic`,
            Desc: `Attributes for determining basic information about a device, setting user device information
such as description of location, and enabling a device.`,
            
            ResetToFactoryDefaults: makeType<ZigBee.IGeneral.Basic.ICmdResetToFactoryDefaults, ZigBee.IGeneral.Basic.ICmdResetToFactoryDefaultsPayload>(command, () => ({
                name: `Reset to Factory Defaults`,
                description: ``,
                id: 0x0000,
                payload: {}
            })),

            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        PowerConfiguration: {
            ID: 0x0001,
            Name: `Power Configuration`,
            Desc: `Attributes for determining more detailed information about a device’s power source(s), and for configuring under/over voltage alarms.`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        DeviceTemperatureConfiguration: {
            ID: 0x0002,
            Name: `Device Temperature Configuration`,
            Desc: `Attributes for determining information about a device’s internal temperature, and for configuring under/over temperature alarms.`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        Identify: {
            ID: 0x0003,
            Name: `Identify`,
            Desc: `Attributes and commands for putting a device into Identification mode (e.g. flashing a light)`,
            
            Identify: makeType<ZigBee.IGeneral.Identify.ICmdIdentify, ZigBee.IGeneral.Identify.ICmdIdentifyPayload>(command, () => ({
                name: `Identify`,
                description: `Start or stop the device identifying itself.`,
                id: 0x0000,
                payload: { 
                    IdentifyTime: ZigBee.General.Types.IdentifyTime,
                }
            })),

            IdentifyQuery: makeType<ZigBee.IGeneral.Identify.ICmdIdentifyQuery, ZigBee.IGeneral.Identify.ICmdIdentifyQueryPayload>(command, () => ({
                name: `Identify Query`,
                description: `Allows the sending device to request the target or targets to respond if they are currently identifying themselves.`,
                id: 0x0001,
                payload: {}
            })),

            TriggerEffect: makeType<ZigBee.IGeneral.Identify.ICmdTriggerEffect, ZigBee.IGeneral.Identify.ICmdTriggerEffectPayload>(command, () => ({
                name: `Trigger Effect`,
                description: `The trigger effect command allows the support of feedback to the user, such as a certain light effect.`,
                id: 0x0040,
                payload: { 
                    IdentifyEffect: ZigBee.General.Types.IdentifyEffect,
                    IdentifyEffectVariant: ZigBee.General.Types.IdentifyEffectVariant,
                }
            })),

            
            IdentifyQueryResponse: makeType<ZigBee.IGeneral.Identify.ICmdIdentifyQueryResponse, ZigBee.IGeneral.Identify.ICmdIdentifyQueryResponsePayload>(command, () => ({
                name: `Identify Query Response`,
                description: `Response of a identify query command.`,
                id: 0x0000,
                payload: { 
                    IdentifyTimeout: ZigBee.General.Types.IdentifyTimeout,
                }
            })),

            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        Groups: {
            ID: 0x0004,
            Name: `Groups`,
            Desc: `Attributes and commands for group configuration and manipulation.`,
            
            AddGroup: makeType<ZigBee.IGeneral.Groups.ICmdAddGroup, ZigBee.IGeneral.Groups.ICmdAddGroupPayload>(command, () => ({
                name: `Add group`,
                description: `will add a group to the device`,
                id: 0x0000,
                payload: { 
                    GroupId: ZigBee.General.Types.GroupId,
                    GroupName: ZigBee.General.Types.GroupName,
                }
            })),

            ViewGroup: makeType<ZigBee.IGeneral.Groups.ICmdViewGroup, ZigBee.IGeneral.Groups.ICmdViewGroupPayload>(command, () => ({
                name: `View group`,
                description: `requests the name of a group`,
                id: 0x0001,
                payload: { 
                    GroupId: ZigBee.General.Types.GroupId,
                }
            })),

            GetGroupMembership: makeType<ZigBee.IGeneral.Groups.ICmdGetGroupMembership, ZigBee.IGeneral.Groups.ICmdGetGroupMembershipPayload>(command, () => ({
                name: `Get group membership`,
                description: `fetches group membership(s). Request with empty list to request all memberships`,
                id: 0x0002,
                payload: { 
                    GroupList: ZigBee.General.Types.GroupList,
                }
            })),

            RemoveGroup: makeType<ZigBee.IGeneral.Groups.ICmdRemoveGroup, ZigBee.IGeneral.Groups.ICmdRemoveGroupPayload>(command, () => ({
                name: `Remove group`,
                description: `Remove a group from the device.`,
                id: 0x0003,
                payload: { 
                    GroupList: ZigBee.General.Types.GroupList,
                }
            })),

            RemoveAllGroups: makeType<ZigBee.IGeneral.Groups.ICmdRemoveAllGroups, ZigBee.IGeneral.Groups.ICmdRemoveAllGroupsPayload>(command, () => ({
                name: `Remove all groups`,
                description: `Remove all group from the device.`,
                id: 0x0004,
                payload: {}
            })),

            AddGroupIfIdentifying: makeType<ZigBee.IGeneral.Groups.ICmdAddGroupIfIdentifying, ZigBee.IGeneral.Groups.ICmdAddGroupIfIdentifyingPayload>(command, () => ({
                name: `Add group if identifying`,
                description: `add a group to the device if the device is currently identifying itself (using the identify cluster)`,
                id: 0x0005,
                payload: { 
                    GroupId: ZigBee.General.Types.GroupId,
                    GroupName: ZigBee.General.Types.GroupName,
                }
            })),

            
            AddGroupResponse: makeType<ZigBee.IGeneral.Groups.ICmdAddGroupResponse, ZigBee.IGeneral.Groups.ICmdAddGroupResponsePayload>(command, () => ({
                name: `Add group response`,
                description: `The Response to the add group request.`,
                id: 0x0000,
                payload: { 
                    Status: ZigBee.General.Types.Status,
                    GroupId: ZigBee.General.Types.GroupId,
                }
            })),

            ViewGroupResponse: makeType<ZigBee.IGeneral.Groups.ICmdViewGroupResponse, ZigBee.IGeneral.Groups.ICmdViewGroupResponsePayload>(command, () => ({
                name: `View group response`,
                description: `The Response to the view group request.`,
                id: 0x0001,
                payload: { 
                    Status: ZigBee.General.Types.Status,
                    GroupId: ZigBee.General.Types.GroupId,
                    GroupName: ZigBee.General.Types.GroupName,
                }
            })),

            GetGroupMembershipResponse: makeType<ZigBee.IGeneral.Groups.ICmdGetGroupMembershipResponse, ZigBee.IGeneral.Groups.ICmdGetGroupMembershipResponsePayload>(command, () => ({
                name: `Get group membership response`,
                description: `The Response to the get group membership request.`,
                id: 0x0002,
                payload: { 
                    GroupCapacity: ZigBee.General.Types.GroupCapacity,
                    GroupList: ZigBee.General.Types.GroupList,
                }
            })),

            RemoveGroupResponse: makeType<ZigBee.IGeneral.Groups.ICmdRemoveGroupResponse, ZigBee.IGeneral.Groups.ICmdRemoveGroupResponsePayload>(command, () => ({
                name: `Remove group response`,
                description: `The Response to the remove group request.`,
                id: 0x0003,
                payload: { 
                    Status: ZigBee.General.Types.Status,
                    GroupId: ZigBee.General.Types.GroupId,
                }
            })),

            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        Scenes: {
            ID: 0x0005,
            Name: `Scenes`,
            Desc: `Attributes and commands for scene configuration and manipulation.`,
            
            AddScene: makeType<ZigBee.IGeneral.Scenes.ICmdAddScene, ZigBee.IGeneral.Scenes.ICmdAddScenePayload>(command, () => ({
                name: `Add scene`,
                description: `Add a scenes to the group.`,
                id: 0x0000,
                payload: { 
                    GroupId: ZigBee.General.Types.GroupId,
                    SceneId: ZigBee.General.Types.SceneId,
                    TransitionTimeSec: ZigBee.General.Types.TransitionTimeSec,
                    SceneName: ZigBee.General.Types.SceneName,
                    SceneExtensions: ZigBee.General.Types.SceneExtensions,
                }
            })),

            ViewScene: makeType<ZigBee.IGeneral.Scenes.ICmdViewScene, ZigBee.IGeneral.Scenes.ICmdViewScenePayload>(command, () => ({
                name: `View scene`,
                description: `Views the scenes of a group.`,
                id: 0x0001,
                payload: { 
                    GroupId: ZigBee.General.Types.GroupId,
                    SceneId: ZigBee.General.Types.SceneId,
                }
            })),

            RemoveScene: makeType<ZigBee.IGeneral.Scenes.ICmdRemoveScene, ZigBee.IGeneral.Scenes.ICmdRemoveScenePayload>(command, () => ({
                name: `Remove scene`,
                description: `Removes a scenes of a group.`,
                id: 0x0002,
                payload: { 
                    GroupId: ZigBee.General.Types.GroupId,
                    SceneId: ZigBee.General.Types.SceneId,
                }
            })),

            RemoveAllScenes: makeType<ZigBee.IGeneral.Scenes.ICmdRemoveAllScenes, ZigBee.IGeneral.Scenes.ICmdRemoveAllScenesPayload>(command, () => ({
                name: `Remove all scenes`,
                description: `Removes all scenes of a group.`,
                id: 0x0003,
                payload: { 
                    GroupId: ZigBee.General.Types.GroupId,
                }
            })),

            StoreScene: makeType<ZigBee.IGeneral.Scenes.ICmdStoreScene, ZigBee.IGeneral.Scenes.ICmdStoreScenePayload>(command, () => ({
                name: `Store scene`,
                description: `Stores a scene of a group for a device.`,
                id: 0x0004,
                payload: { 
                    GroupId: ZigBee.General.Types.GroupId,
                    SceneId: ZigBee.General.Types.SceneId,
                }
            })),

            RecallScene: makeType<ZigBee.IGeneral.Scenes.ICmdRecallScene, ZigBee.IGeneral.Scenes.ICmdRecallScenePayload>(command, () => ({
                name: `Recall scene`,
                description: `Recalls a scene of a group for a device.`,
                id: 0x0005,
                payload: { 
                    GroupId: ZigBee.General.Types.GroupId,
                    SceneId: ZigBee.General.Types.SceneId,
                }
            })),

            GetSceneMembership: makeType<ZigBee.IGeneral.Scenes.ICmdGetSceneMembership, ZigBee.IGeneral.Scenes.ICmdGetSceneMembershipPayload>(command, () => ({
                name: `Get scene membership`,
                description: `Get the scenes of a group for a device.`,
                id: 0x0006,
                payload: { 
                    GroupId: ZigBee.General.Types.GroupId,
                }
            })),

            EnhancedAddScene: makeType<ZigBee.IGeneral.Scenes.ICmdEnhancedAddScene, ZigBee.IGeneral.Scenes.ICmdEnhancedAddScenePayload>(command, () => ({
                name: `Enhanced Add Scene`,
                description: `Same as Add Scene, except that transition time is specified in 1/10 s`,
                id: 0x0040,
                payload: { 
                    GroupId: ZigBee.General.Types.GroupId,
                    SceneId: ZigBee.General.Types.SceneId,
                    TransitionTime: ZigBee.General.Types.TransitionTime,
                    SceneName: ZigBee.General.Types.SceneName,
                    SceneExtensions: ZigBee.General.Types.SceneExtensions,
                }
            })),

            EnhancedViewScene: makeType<ZigBee.IGeneral.Scenes.ICmdEnhancedViewScene, ZigBee.IGeneral.Scenes.ICmdEnhancedViewScenePayload>(command, () => ({
                name: `Enhanced View Scene`,
                description: `Views the scenes of a group (returning transition time with 1/10s precision).`,
                id: 0x0041,
                payload: { 
                    GroupId: ZigBee.General.Types.GroupId,
                    SceneId: ZigBee.General.Types.SceneId,
                }
            })),

            CopyScene: makeType<ZigBee.IGeneral.Scenes.ICmdCopyScene, ZigBee.IGeneral.Scenes.ICmdCopyScenePayload>(command, () => ({
                name: `Copy Scene`,
                description: ``,
                id: 0x0042,
                payload: { 
                    SceneCopyMode: ZigBee.General.Types.SceneCopyMode,
                    FromGroupId: ZigBee.General.Types.GroupId,
                    FromSceneId: ZigBee.General.Types.SceneId,
                    ToGroupId: ZigBee.General.Types.GroupId,
                    ToSceneId: ZigBee.General.Types.SceneId,
                }
            })),

            IkeaRemotePress: makeType<ZigBee.IGeneral.Scenes.ICmdIkeaRemotePress, ZigBee.IGeneral.Scenes.ICmdIkeaRemotePressPayload>(command, () => ({
                name: `Ikea Remote Press`,
                description: `Custom left/right action of Ikea remote`,
                id: 0x0007,
                mnf: 0x117c,
                payload: { 
                    IkeaRemoteDirection: ZigBee.General.Types.IkeaRemoteDirection,
                }
            })),

            IkeaRemoteLongpressStart: makeType<ZigBee.IGeneral.Scenes.ICmdIkeaRemoteLongpressStart, ZigBee.IGeneral.Scenes.ICmdIkeaRemoteLongpressStartPayload>(command, () => ({
                name: `Ikea Remote Longpress Start`,
                description: ``,
                id: 0x0008,
                mnf: 0x117c,
                payload: { 
                    IkeaRemoteDirection: ZigBee.General.Types.IkeaRemoteDirection,
                }
            })),

            IkeaRemoteLongpressStop: makeType<ZigBee.IGeneral.Scenes.ICmdIkeaRemoteLongpressStop, ZigBee.IGeneral.Scenes.ICmdIkeaRemoteLongpressStopPayload>(command, () => ({
                name: `Ikea Remote Longpress Stop`,
                description: ``,
                id: 0x0009,
                mnf: 0x117c,
                payload: {}
            })),

            
            AddSceneResponse: makeType<ZigBee.IGeneral.Scenes.ICmdAddSceneResponse, ZigBee.IGeneral.Scenes.ICmdAddSceneResponsePayload>(command, () => ({
                name: `Add scene response`,
                description: `Response to the add scene command.`,
                id: 0x0000,
                payload: { 
                    Status: ZigBee.General.Types.Status,
                    GroupId: ZigBee.General.Types.GroupId,
                    SceneId: ZigBee.General.Types.SceneId,
                }
            })),

            ViewSceneResponse: makeType<ZigBee.IGeneral.Scenes.ICmdViewSceneResponse, ZigBee.IGeneral.Scenes.ICmdViewSceneResponsePayload>(command, () => ({
                name: `View scene response`,
                description: `Response to the view scene command.`,
                id: 0x0001,
                payload: { 
                    Status: ZigBee.General.Types.Status,
                    GroupId: ZigBee.General.Types.GroupId,
                    SceneId: ZigBee.General.Types.SceneId,
                    TransitionTimeSec: ZigBee.General.Types.TransitionTimeSec,
                    SceneName: ZigBee.General.Types.SceneName,
                    SceneExtensions: ZigBee.General.Types.SceneExtensions,
                }
            })),

            RemoveSceneResponse: makeType<ZigBee.IGeneral.Scenes.ICmdRemoveSceneResponse, ZigBee.IGeneral.Scenes.ICmdRemoveSceneResponsePayload>(command, () => ({
                name: `Remove scene response`,
                description: `Response to the remove scene command.`,
                id: 0x0002,
                payload: { 
                    Status: ZigBee.General.Types.Status,
                    GroupId: ZigBee.General.Types.GroupId,
                    SceneId: ZigBee.General.Types.SceneId,
                }
            })),

            RemoveAllScenesResponse: makeType<ZigBee.IGeneral.Scenes.ICmdRemoveAllScenesResponse, ZigBee.IGeneral.Scenes.ICmdRemoveAllScenesResponsePayload>(command, () => ({
                name: `Remove all scenes response`,
                description: `Response to the remove all scenes command.`,
                id: 0x0003,
                payload: { 
                    Status: ZigBee.General.Types.Status,
                    GroupId: ZigBee.General.Types.GroupId,
                }
            })),

            StoreSceneResponse: makeType<ZigBee.IGeneral.Scenes.ICmdStoreSceneResponse, ZigBee.IGeneral.Scenes.ICmdStoreSceneResponsePayload>(command, () => ({
                name: `Store scene response`,
                description: `Response to the store scene command.`,
                id: 0x0004,
                payload: { 
                    Status: ZigBee.General.Types.Status,
                    GroupId: ZigBee.General.Types.GroupId,
                    SceneId: ZigBee.General.Types.SceneId,
                }
            })),

            GetSceneMembershipResponse: makeType<ZigBee.IGeneral.Scenes.ICmdGetSceneMembershipResponse, ZigBee.IGeneral.Scenes.ICmdGetSceneMembershipResponsePayload>(command, () => ({
                name: `Get scene membership response`,
                description: `Shows details about scene membership.`,
                id: 0x0006,
                payload: { 
                    Status: ZigBee.General.Types.Status,
                    SceneCapacity: ZigBee.General.Types.SceneCapacity,
                    GroupId: ZigBee.General.Types.GroupId,
                    SceneList: ZigBee.General.Types.SceneList,
                }
            })),

            EnhancedAddSceneResponse: makeType<ZigBee.IGeneral.Scenes.ICmdEnhancedAddSceneResponse, ZigBee.IGeneral.Scenes.ICmdEnhancedAddSceneResponsePayload>(command, () => ({
                name: `Enhanced Add Scene response`,
                description: ``,
                id: 0x0040,
                payload: { 
                    Status: ZigBee.General.Types.Status,
                    GroupId: ZigBee.General.Types.GroupId,
                    SceneId: ZigBee.General.Types.SceneId,
                }
            })),

            EnhancedViewSceneResponse: makeType<ZigBee.IGeneral.Scenes.ICmdEnhancedViewSceneResponse, ZigBee.IGeneral.Scenes.ICmdEnhancedViewSceneResponsePayload>(command, () => ({
                name: `Enhanced View Scene response`,
                description: `A scene description.`,
                id: 0x0041,
                payload: { 
                    Status: ZigBee.General.Types.Status,
                    GroupId: ZigBee.General.Types.GroupId,
                    SceneId: ZigBee.General.Types.SceneId,
                    TransitionTime: ZigBee.General.Types.TransitionTime,
                    SceneName: ZigBee.General.Types.SceneName,
                    SceneExtensions: ZigBee.General.Types.SceneExtensions,
                }
            })),

            CopySceneResponse: makeType<ZigBee.IGeneral.Scenes.ICmdCopySceneResponse, ZigBee.IGeneral.Scenes.ICmdCopySceneResponsePayload>(command, () => ({
                name: `Copy Scene Response`,
                description: ``,
                id: 0x0042,
                payload: { 
                    Status: ZigBee.General.Types.Status,
                    FromGroupId: ZigBee.General.Types.GroupId,
                    FromSceneId: ZigBee.General.Types.SceneId,
                }
            })),

            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        OnOff: {
            ID: 0x0006,
            Name: `On/Off`,
            Desc: `Attributes and commands for switching devices between 'On' and 'Off' states.`,
            
            Off: makeType<ZigBee.IGeneral.OnOff.ICmdOff, ZigBee.IGeneral.OnOff.ICmdOffPayload>(command, () => ({
                name: `Off`,
                description: `On receipt of this command, a device shall enter its 'Off' state. This state is device dependent, but it is recommended that it is used for power off or similar functions.`,
                id: 0x0000,
                payload: {}
            })),

            On: makeType<ZigBee.IGeneral.OnOff.ICmdOn, ZigBee.IGeneral.OnOff.ICmdOnPayload>(command, () => ({
                name: `On`,
                description: `On receipt of this command, a device shall enter its 'On' state. This state is device dependent, but it is recommended that it is used for power on or similar functions.`,
                id: 0x0001,
                payload: {}
            })),

            Toggle: makeType<ZigBee.IGeneral.OnOff.ICmdToggle, ZigBee.IGeneral.OnOff.ICmdTogglePayload>(command, () => ({
                name: `Toggle`,
                description: `On receipt of this command, if a device is in its ‘Off’ state it shall enter its 'On' state. Otherwise, if it is in its ‘On’ state it shall enter its 'Off' state.`,
                id: 0x0002,
                payload: {}
            })),

            OffWithEffect: makeType<ZigBee.IGeneral.OnOff.ICmdOffWithEffect, ZigBee.IGeneral.OnOff.ICmdOffWithEffectPayload>(command, () => ({
                name: `Off with effect`,
                description: ``,
                id: 0x0040,
                payload: { 
                    EffectIdentifier: ZigBee.General.Types.EffectIdentifier,
                    EffectVariant: ZigBee.General.Types.EffectVariant,
                }
            })),

            OnWithRecallGlobalScene: makeType<ZigBee.IGeneral.OnOff.ICmdOnWithRecallGlobalScene, ZigBee.IGeneral.OnOff.ICmdOnWithRecallGlobalScenePayload>(command, () => ({
                name: `On with recall global scene`,
                description: `The on with recall global scene command allows the recall of the light settings when the light was turned off.`,
                id: 0x0041,
                payload: {}
            })),

            OnWithTimedOff: makeType<ZigBee.IGeneral.OnOff.ICmdOnWithTimedOff, ZigBee.IGeneral.OnOff.ICmdOnWithTimedOffPayload>(command, () => ({
                name: `On with timed off`,
                description: `Allows lamps to be turned on for a specific duration with a guarded off duration so that should the lamp be subsequently switched off, further on with timed off commands, received during this time, are prevented from turning the lamps back on.`,
                id: 0x0042,
                payload: { 
                    OnOffControl: ZigBee.General.Types.OnOffControl,
                    OnTime: ZigBee.General.Types.OnTime,
                    OffWaitTime: ZigBee.General.Types.OffWaitTime,
                }
            })),

            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        OnOffSwitchConfiguration: {
            ID: 0x0007,
            Name: `On/Off Switch Configuration`,
            Desc: `Attributes and commands for configuring On/Off switching devices`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        LevelControl: {
            ID: 0x0008,
            Name: `Level Control`,
            Desc: `provides an interface for controlling a characteristic of a device that can be set to a level.
For example the brightness of a light, the degree of closure of a door, or the power output of a heater.`,
            
            MoveToLevel: makeType<ZigBee.IGeneral.LevelControl.ICmdMoveToLevel, ZigBee.IGeneral.LevelControl.ICmdMoveToLevelPayload>(command, () => ({
                name: `Move to Level`,
                description: ``,
                id: 0x0000,
                payload: { 
                    Level: ZigBee.General.Types.Level,
                    TransitionTime: ZigBee.General.Types.TransitionTime,
                }
            })),

            Move: makeType<ZigBee.IGeneral.LevelControl.ICmdMove, ZigBee.IGeneral.LevelControl.ICmdMovePayload>(command, () => ({
                name: `Move`,
                description: ``,
                id: 0x0001,
                payload: { 
                    LevelDirection: ZigBee.General.Types.LevelDirection,
                    Rate: ZigBee.General.Types.Rate,
                }
            })),

            Step: makeType<ZigBee.IGeneral.LevelControl.ICmdStep, ZigBee.IGeneral.LevelControl.ICmdStepPayload>(command, () => ({
                name: `Step`,
                description: ``,
                id: 0x0002,
                payload: { 
                    LevelDirection: ZigBee.General.Types.LevelDirection,
                    StepSize: ZigBee.General.Types.StepSize,
                    TransitionTime: ZigBee.General.Types.TransitionTime,
                }
            })),

            Stop: makeType<ZigBee.IGeneral.LevelControl.ICmdStop, ZigBee.IGeneral.LevelControl.ICmdStopPayload>(command, () => ({
                name: `Stop`,
                description: ``,
                id: 0x0003,
                payload: {}
            })),

            MoveToLevelWithOnOff: makeType<ZigBee.IGeneral.LevelControl.ICmdMoveToLevelWithOnOff, ZigBee.IGeneral.LevelControl.ICmdMoveToLevelWithOnOffPayload>(command, () => ({
                name: `Move to Level (with On/Off)`,
                description: ``,
                id: 0x0004,
                payload: { 
                    Level: ZigBee.General.Types.Level,
                    TransitionTime: ZigBee.General.Types.TransitionTime,
                }
            })),

            MoveWithOnOff: makeType<ZigBee.IGeneral.LevelControl.ICmdMoveWithOnOff, ZigBee.IGeneral.LevelControl.ICmdMoveWithOnOffPayload>(command, () => ({
                name: `Move (with On/Off)`,
                description: ``,
                id: 0x0005,
                payload: { 
                    LevelDirection: ZigBee.General.Types.LevelDirection,
                    Rate: ZigBee.General.Types.Rate,
                }
            })),

            StepWithOnOff: makeType<ZigBee.IGeneral.LevelControl.ICmdStepWithOnOff, ZigBee.IGeneral.LevelControl.ICmdStepWithOnOffPayload>(command, () => ({
                name: `Step (with On/Off)`,
                description: ``,
                id: 0x0006,
                payload: { 
                    LevelDirection: ZigBee.General.Types.LevelDirection,
                    StepSize: ZigBee.General.Types.StepSize,
                    TransitionTime: ZigBee.General.Types.TransitionTime,
                }
            })),

            StopWithOnOff: makeType<ZigBee.IGeneral.LevelControl.ICmdStopWithOnOff, ZigBee.IGeneral.LevelControl.ICmdStopWithOnOffPayload>(command, () => ({
                name: `Stop (with On/Off)`,
                description: ``,
                id: 0x0007,
                payload: {}
            })),

            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        Alarms: {
            ID: 0x0009,
            Name: `Alarms`,
            Desc: `Sending alarm notifications and configuring alarm functionality.`,
            
            ResetAlarm: makeType<ZigBee.IGeneral.Alarms.ICmdResetAlarm, ZigBee.IGeneral.Alarms.ICmdResetAlarmPayload>(command, () => ({
                name: `Reset alarm`,
                description: ``,
                id: 0x0000,
                payload: { 
                    AlarmCode: ZigBee.General.Types.AlarmCode,
                    ClusterId: ZigBee.General.Types.ClusterId,
                }
            })),

            ResetAllAlarms: makeType<ZigBee.IGeneral.Alarms.ICmdResetAllAlarms, ZigBee.IGeneral.Alarms.ICmdResetAllAlarmsPayload>(command, () => ({
                name: `Reset all alarms`,
                description: `Resets all alarms, causing triggered alarms to generate new notification`,
                id: 0x0001,
                payload: {}
            })),

            GetAlarm: makeType<ZigBee.IGeneral.Alarms.ICmdGetAlarm, ZigBee.IGeneral.Alarms.ICmdGetAlarmPayload>(command, () => ({
                name: `Get Alarm`,
                description: `Retrieves the earliest alarm and removes it from the table`,
                id: 0x0002,
                payload: {}
            })),

            ResetAlarmLog: makeType<ZigBee.IGeneral.Alarms.ICmdResetAlarmLog, ZigBee.IGeneral.Alarms.ICmdResetAlarmLogPayload>(command, () => ({
                name: `Reset alarm log`,
                description: `Clears the alarm log`,
                id: 0x0003,
                payload: {}
            })),

            
            Alarm: makeType<ZigBee.IGeneral.Alarms.ICmdAlarm, ZigBee.IGeneral.Alarms.ICmdAlarmPayload>(command, () => ({
                name: `Alarm`,
                description: ``,
                id: 0x0000,
                payload: { 
                    AlarmCode: ZigBee.General.Types.AlarmCode,
                    ClusterId: ZigBee.General.Types.ClusterId,
                }
            })),

            GetAlarmResponse: makeType<ZigBee.IGeneral.Alarms.ICmdGetAlarmResponse, ZigBee.IGeneral.Alarms.ICmdGetAlarmResponsePayload>(command, () => ({
                name: `Get alarm response`,
                description: ``,
                id: 0x0001,
                payload: { 
                    Status: ZigBee.General.Types.Status,
                    AlarmCode: ZigBee.General.Types.AlarmCode,
                    ClusterId: ZigBee.General.Types.ClusterId,
                    Time: ZigBee.General.Types.Time,
                }
            })),

            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        Time: {
            ID: 0x000A,
            Name: `Time`,
            Desc: `This cluster provides a basic interface to a real-time clock.`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        Location: {
            ID: 0x000B,
            Name: `Location`,
            Desc: `provides distance between devices`,
            
            SetAbsoluteLocation: makeType<ZigBee.IGeneral.Location.ICmdSetAbsoluteLocation, ZigBee.IGeneral.Location.ICmdSetAbsoluteLocationPayload>(command, () => ({
                name: `Set Absolute Location`,
                description: ``,
                id: 0x0000,
                payload: { 
                    XCoordinate: ZigBee.General.Types.XCoordinate,
                    YCoordinate: ZigBee.General.Types.YCoordinate,
                    ZCoordinate: ZigBee.General.Types.ZCoordinate,
                    Power: ZigBee.General.Types.Power,
                    PathLossExponent: ZigBee.General.Types.PathLossExponent,
                }
            })),

            SetDeviceConfiguration: makeType<ZigBee.IGeneral.Location.ICmdSetDeviceConfiguration, ZigBee.IGeneral.Location.ICmdSetDeviceConfigurationPayload>(command, () => ({
                name: `Set Device Configuration`,
                description: ``,
                id: 0x0001,
                payload: { 
                    Power: ZigBee.General.Types.Power,
                    PathLossExponent: ZigBee.General.Types.PathLossExponent,
                    CalculationPeriod: ZigBee.General.Types.CalculationPeriod,
                    NumberRssiMeasurements: ZigBee.General.Types.NumberRssiMeasurements,
                    ReportingPeriod: ZigBee.General.Types.ReportingPeriod,
                }
            })),

            GetDeviceConfiguration: makeType<ZigBee.IGeneral.Location.ICmdGetDeviceConfiguration, ZigBee.IGeneral.Location.ICmdGetDeviceConfigurationPayload>(command, () => ({
                name: `Get Device Configuration`,
                description: ``,
                id: 0x0002,
                payload: { 
                    TargetAddress: ZigBee.General.Types.Device,
                }
            })),

            GetLocationData: makeType<ZigBee.IGeneral.Location.ICmdGetLocationData, ZigBee.IGeneral.Location.ICmdGetLocationDataPayload>(command, () => ({
                name: `Get Location Data`,
                description: ``,
                id: 0x0003,
                payload: { 
                    LocationFlags: ZigBee.General.Types.LocationFlags,
                    NumberResponses: ZigBee.General.Types.NumberResponses,
                    TargetAddress: ZigBee.General.Types.Device,
                }
            })),

            RssiResponse: makeType<ZigBee.IGeneral.Location.ICmdRssiResponse, ZigBee.IGeneral.Location.ICmdRssiResponsePayload>(command, () => ({
                name: `RSSI Response`,
                description: ``,
                id: 0x0004,
                payload: { 
                    Device: ZigBee.General.Types.Device,
                    XCoordinate: ZigBee.General.Types.XCoordinate,
                    YCoordinate: ZigBee.General.Types.YCoordinate,
                    ZCoordinate: ZigBee.General.Types.ZCoordinate,
                    Rssi: ZigBee.General.Types.Rssi,
                    NumberRssiMeasurements: ZigBee.General.Types.NumberRssiMeasurements,
                }
            })),

            SendPings: makeType<ZigBee.IGeneral.Location.ICmdSendPings, ZigBee.IGeneral.Location.ICmdSendPingsPayload>(command, () => ({
                name: `Send Pings`,
                description: ``,
                id: 0x0005,
                payload: { 
                    TargetAddress: ZigBee.General.Types.Device,
                    NumberRssiMeasurements: ZigBee.General.Types.NumberRssiMeasurements,
                    CalculationPeriod: ZigBee.General.Types.CalculationPeriod,
                }
            })),

            AnchorNodeAnnounce: makeType<ZigBee.IGeneral.Location.ICmdAnchorNodeAnnounce, ZigBee.IGeneral.Location.ICmdAnchorNodeAnnouncePayload>(command, () => ({
                name: `Anchor Node Announce`,
                description: ``,
                id: 0x0006,
                payload: { 
                    Device: ZigBee.General.Types.Device,
                    XCoordinate: ZigBee.General.Types.XCoordinate,
                    YCoordinate: ZigBee.General.Types.YCoordinate,
                    ZCoordinate: ZigBee.General.Types.ZCoordinate,
                }
            })),

            DistanceMeasure: makeType<ZigBee.IGeneral.Location.ICmdDistanceMeasure, ZigBee.IGeneral.Location.ICmdDistanceMeasurePayload>(command, () => ({
                name: `Distance measure`,
                description: ``,
                id: 0x0040,
                payload: { 
                    TargetAddress: ZigBee.General.Types.Device,
                    Resolution: ZigBee.General.Types.Resolution,
                }
            })),

            
            DeviceConfigurationResponse: makeType<ZigBee.IGeneral.Location.ICmdDeviceConfigurationResponse, ZigBee.IGeneral.Location.ICmdDeviceConfigurationResponsePayload>(command, () => ({
                name: `Device Configuration Response`,
                description: ``,
                id: 0x0000,
                payload: { 
                    Status: ZigBee.General.Types.Status,
                    Power: ZigBee.General.Types.Power,
                    PathLossExponent: ZigBee.General.Types.PathLossExponent,
                    CalculationPeriod: ZigBee.General.Types.CalculationPeriod,
                    NumberRssiMeasurements: ZigBee.General.Types.NumberRssiMeasurements,
                    ReportingPeriod: ZigBee.General.Types.ReportingPeriod,
                }
            })),

            LocationDataResponse: makeType<ZigBee.IGeneral.Location.ICmdLocationDataResponse, ZigBee.IGeneral.Location.ICmdLocationDataResponsePayload>(command, () => ({
                name: `Location Data Response`,
                description: ``,
                id: 0x0001,
                payload: { 
                    Status: ZigBee.General.Types.Status,
                    LocationType: ZigBee.General.Types.LocationType,
                    XCoordinate: ZigBee.General.Types.XCoordinate,
                    YCoordinate: ZigBee.General.Types.YCoordinate,
                    ZCoordinate: ZigBee.General.Types.ZCoordinate,
                    Power: ZigBee.General.Types.Power,
                    PathLossExponent: ZigBee.General.Types.PathLossExponent,
                    LocationMethod: ZigBee.General.Types.LocationMethod,
                    QualityMeasure: ZigBee.General.Types.QualityMeasure,
                    LocationAge: ZigBee.General.Types.LocationAge,
                }
            })),

            LocationDataNotification: makeType<ZigBee.IGeneral.Location.ICmdLocationDataNotification, ZigBee.IGeneral.Location.ICmdLocationDataNotificationPayload>(command, () => ({
                name: `Location Data Notification`,
                description: ``,
                id: 0x0002,
                payload: { 
                    LocationType: ZigBee.General.Types.LocationType,
                    XCoordinate: ZigBee.General.Types.XCoordinate,
                    YCoordinate: ZigBee.General.Types.YCoordinate,
                    ZCoordinate: ZigBee.General.Types.ZCoordinate,
                    Power: ZigBee.General.Types.Power,
                    PathLossExponent: ZigBee.General.Types.PathLossExponent,
                    LocationMethod: ZigBee.General.Types.LocationMethod,
                    QualityMeasure: ZigBee.General.Types.QualityMeasure,
                    LocationAge: ZigBee.General.Types.LocationAge,
                }
            })),

            CompactLocationDataNotification: makeType<ZigBee.IGeneral.Location.ICmdCompactLocationDataNotification, ZigBee.IGeneral.Location.ICmdCompactLocationDataNotificationPayload>(command, () => ({
                name: `Compact Location Data Notification`,
                description: ``,
                id: 0x0003,
                payload: { 
                    LocationType: ZigBee.General.Types.LocationType,
                    XCoordinate: ZigBee.General.Types.XCoordinate,
                    YCoordinate: ZigBee.General.Types.YCoordinate,
                    ZCoordinate: ZigBee.General.Types.ZCoordinate,
                    QualityMeasure: ZigBee.General.Types.QualityMeasure,
                    LocationAge: ZigBee.General.Types.LocationAge,
                }
            })),

            RssiPing: makeType<ZigBee.IGeneral.Location.ICmdRssiPing, ZigBee.IGeneral.Location.ICmdRssiPingPayload>(command, () => ({
                name: `RSSI Ping`,
                description: ``,
                id: 0x0004,
                payload: { 
                    LocationType: ZigBee.General.Types.LocationType,
                }
            })),

            RssiRequest: makeType<ZigBee.IGeneral.Location.ICmdRssiRequest, ZigBee.IGeneral.Location.ICmdRssiRequestPayload>(command, () => ({
                name: `RSSI Request`,
                description: ``,
                id: 0x0005,
                payload: {}
            })),

            ReportRssiMeasurements: makeType<ZigBee.IGeneral.Location.ICmdReportRssiMeasurements, ZigBee.IGeneral.Location.ICmdReportRssiMeasurementsPayload>(command, () => ({
                name: `Report RSSI Measurements`,
                description: ``,
                id: 0x0006,
                payload: { 
                    Device: ZigBee.General.Types.Device,
                    NeighborsInfoList: ZigBee.General.Types.NeighborsInfoList,
                }
            })),

            RequestOwnLocation: makeType<ZigBee.IGeneral.Location.ICmdRequestOwnLocation, ZigBee.IGeneral.Location.ICmdRequestOwnLocationPayload>(command, () => ({
                name: `Request Own Location`,
                description: ``,
                id: 0x0007,
                payload: { 
                    BlindNodeAddress: ZigBee.General.Types.Device,
                }
            })),

            DistanceMeasureResponse: makeType<ZigBee.IGeneral.Location.ICmdDistanceMeasureResponse, ZigBee.IGeneral.Location.ICmdDistanceMeasureResponsePayload>(command, () => ({
                name: `Distance measure response`,
                description: `Returns the result of a distance measure.`,
                id: 0x0040,
                payload: { 
                    TargetAddress: ZigBee.General.Types.Device,
                    Distance: ZigBee.General.Types.Distance,
                    QualityIndex: ZigBee.General.Types.QualityIndex,
                }
            })),

            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        AnalogInput: {
            ID: 0x000C,
            Name: `Analog Input`,
            Desc: `is an interface for reading the value of an analog measurement and accessing various characteristics of that measurement.`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        AnalogOutput: {
            ID: 0x000D,
            Name: `Analog Output`,
            Desc: `is an interface for setting the value of an analog output (typically to the environment) and accessing various characteristics of that value.`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        AnalogValue: {
            ID: 0x000E,
            Name: `Analog Value`,
            Desc: `is an interface for setting an analog value, typically used as a control system parameter, and accessing various characteristics of that value.`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        BinaryInput: {
            ID: 0x000F,
            Name: `Binary Input`,
            Desc: `is an interface for reading the value of a binary measurement and accessing various characteristics of
that measurement. The cluster is typically used to implement a sensor that measures a two-state physical quantity.`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        BinaryOutput: {
            ID: 0x0010,
            Name: `Binary Output`,
            Desc: ``,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        BinaryValue: {
            ID: 0x0011,
            Name: `Binary Value`,
            Desc: ``,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        MultistateInput: {
            ID: 0x0012,
            Name: `Multistate Input`,
            Desc: `is an interface for reading the value of a multistate measurement
and accessing various characteristics of that measurement. The cluster is typically
used to implement a sensor that measures a physical quantity that can take on
one of a number of discrete states.`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        MultistateOutput: {
            ID: 0x0013,
            Name: `Multistate Output`,
            Desc: ``,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        MultistateValue: {
            ID: 0x0014,
            Name: `Multistate Value`,
            Desc: ``,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        PollControl: {
            ID: 0x0020,
            Name: `Poll Control`,
            Desc: `Provides a mechanism for the management of an end device’s MAC Data Request rate.`,
            
            CheckIn: makeType<ZigBee.IGeneral.PollControl.ICmdCheckIn, ZigBee.IGeneral.PollControl.ICmdCheckInPayload>(command, () => ({
                name: `Check-in`,
                description: ``,
                id: 0x0000,
                payload: {}
            })),

            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        MeterIdentification: {
            ID: 0x0B01,
            Name: `Meter Identification`,
            Desc: `Attributes and commands that provide an interface tometer identification`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        Diagnostics: {
            ID: 0x0B05,
            Name: `Diagnostics`,
            Desc: `The diagnostics cluster provides access to information regarding the operation of the ZigBee stack over time.`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        }
    };
    
    ZigBee.General.Basic.Server.Attribute = { 
        0xFFFD: ZigBee.General.Types.ClusterRevision,
        0x0000: ZigBee.General.Types.ZclVersion,
        0x0001: ZigBee.General.Types.ApplicationVersion,
        0x0002: ZigBee.General.Types.StackVersion,
        0x0003: ZigBee.General.Types.HwVersion,
        0x0004: ZigBee.General.Types.ManufacturerName,
        0x0005: ZigBee.General.Types.ModelIdentifier,
        0x0006: ZigBee.General.Types.DateCode,
        0x0007: ZigBee.General.Types.PowerSource,
        0x0008: ZigBee.General.Types.GenericDeviceClass,
        0x0009: ZigBee.General.Types.GenericDeviceType,
        0x0010: ZigBee.General.Types.LocationDescription,
        0x0011: ZigBee.General.Types.PhysicalEnvironment,
        0x0012: ZigBee.General.Types.DeviceEnabled,
        0x0013: ZigBee.General.Types.AlarmMask,
        0x0014: ZigBee.General.Types.DisableLocalConfig,
        0x4000: ZigBee.General.Types.SwBuildId,
        0x000A: ZigBee.General.Types.ProductCode,
        0x000B: ZigBee.General.Types.ProductUrl,
        0x0030: ZigBee.General.Types.Sensitivity,
        0x0031: ZigBee.General.Types.Configuration,
        0x0032: ZigBee.General.Types.UserTest,
        0x0033: ZigBee.General.Types.LedIndication,
    };
    ZigBee.General.Basic.Client.Attribute = { 
    };
    ZigBee.General.Basic.Server.Command = { 
        0x00: ZigBee.General.Basic.ResetToFactoryDefaults,
    };
    ZigBee.General.Basic.Client.Command = { 
    };
    ZigBee.General.PowerConfiguration.Server.Attribute = { 
        0x0000: ZigBee.General.Types.MainsVoltage,
        0x0001: ZigBee.General.Types.MainsFrequency,
        0x0010: ZigBee.General.Types.MainsAlarmMask,
        0x0011: ZigBee.General.Types.MainsVoltageMinThreshold,
        0x0012: ZigBee.General.Types.MainsVoltageMaxThreshold,
        0x0013: ZigBee.General.Types.MainsVoltageDwellTripPoint,
        0x0020: ZigBee.General.Types.BatteryVoltage,
        0x0021: ZigBee.General.Types.BatteryRemaining,
        0x0030: ZigBee.General.Types.BatteryManufacturer,
        0x0031: ZigBee.General.Types.BatterySize,
        0x0032: ZigBee.General.Types.BatteryCapacity,
        0x0033: ZigBee.General.Types.BatteryQuantity,
        0x0034: ZigBee.General.Types.BatteryRatedVoltage,
        0x0035: ZigBee.General.Types.BatteryAlarmMask,
        0x0036: ZigBee.General.Types.BatteryVoltageMinThreshold,
        0x0037: ZigBee.General.Types.BatteryVoltageThreshold1,
        0x0038: ZigBee.General.Types.BatteryVoltageThreshold2,
        0x0039: ZigBee.General.Types.BatteryVoltageThreshold3,
        0x003A: ZigBee.General.Types.BatteryPercentageMinThreshold,
        0x003B: ZigBee.General.Types.BatteryPercentageThreshold1,
        0x003C: ZigBee.General.Types.BatteryPercentageThreshold2,
        0x003D: ZigBee.General.Types.BatteryPercentageThreshold3,
        0x003E: ZigBee.General.Types.BatteryAlarmState,
    };
    ZigBee.General.PowerConfiguration.Client.Attribute = { 
    };
    ZigBee.General.PowerConfiguration.Server.Command = { 
    };
    ZigBee.General.PowerConfiguration.Client.Command = { 
    };
    ZigBee.General.DeviceTemperatureConfiguration.Server.Attribute = { 
        0x0000: ZigBee.General.Types.CurrentTemperature,
        0x0001: ZigBee.General.Types.MinTempExperienced,
        0x0002: ZigBee.General.Types.MaxTempExperienced,
        0x0003: ZigBee.General.Types.OverTempTotalDwell,
        0x0010: ZigBee.General.Types.DeviceTempAlarmMask,
        0x0011: ZigBee.General.Types.LowTempThreshold,
        0x0012: ZigBee.General.Types.HighTempThreshold,
        0x0013: ZigBee.General.Types.LowTempDwellTripPoint,
        0x0014: ZigBee.General.Types.HighTempDwellTripPoint,
    };
    ZigBee.General.DeviceTemperatureConfiguration.Client.Attribute = { 
    };
    ZigBee.General.DeviceTemperatureConfiguration.Server.Command = { 
    };
    ZigBee.General.DeviceTemperatureConfiguration.Client.Command = { 
    };
    ZigBee.General.Identify.Server.Attribute = { 
        0x0000: ZigBee.General.Types.IdentifyTime,
    };
    ZigBee.General.Identify.Client.Attribute = { 
    };
    ZigBee.General.Identify.Server.Command = { 
        0x00: ZigBee.General.Identify.Identify,
        0x01: ZigBee.General.Identify.IdentifyQuery,
        0x40: ZigBee.General.Identify.TriggerEffect,
    };
    ZigBee.General.Identify.Client.Command = { 
        0x00: ZigBee.General.Identify.IdentifyQueryResponse,
    };
    ZigBee.General.Groups.Server.Attribute = { 
        0x0000: ZigBee.General.Types.GroupNameSupport,
    };
    ZigBee.General.Groups.Client.Attribute = { 
    };
    ZigBee.General.Groups.Server.Command = { 
        0x00: ZigBee.General.Groups.AddGroup,
        0x01: ZigBee.General.Groups.ViewGroup,
        0x02: ZigBee.General.Groups.GetGroupMembership,
        0x03: ZigBee.General.Groups.RemoveGroup,
        0x04: ZigBee.General.Groups.RemoveAllGroups,
        0x05: ZigBee.General.Groups.AddGroupIfIdentifying,
    };
    ZigBee.General.Groups.Client.Command = { 
        0x00: ZigBee.General.Groups.AddGroupResponse,
        0x01: ZigBee.General.Groups.ViewGroupResponse,
        0x02: ZigBee.General.Groups.GetGroupMembershipResponse,
        0x03: ZigBee.General.Groups.RemoveGroupResponse,
    };
    ZigBee.General.Scenes.Server.Attribute = { 
        0x0000: ZigBee.General.Types.SceneCount,
        0x0001: ZigBee.General.Types.CurrentScene,
        0x0002: ZigBee.General.Types.CurrentGroup,
        0x0003: ZigBee.General.Types.SceneValid,
        0x0004: ZigBee.General.Types.SceneNameSupport,
        0x0005: ZigBee.General.Types.SceneLastConfiguredBy,
    };
    ZigBee.General.Scenes.Client.Attribute = { 
    };
    ZigBee.General.Scenes.Server.Command = { 
        0x00: ZigBee.General.Scenes.AddScene,
        0x01: ZigBee.General.Scenes.ViewScene,
        0x02: ZigBee.General.Scenes.RemoveScene,
        0x03: ZigBee.General.Scenes.RemoveAllScenes,
        0x04: ZigBee.General.Scenes.StoreScene,
        0x05: ZigBee.General.Scenes.RecallScene,
        0x06: ZigBee.General.Scenes.GetSceneMembership,
        0x40: ZigBee.General.Scenes.EnhancedAddScene,
        0x41: ZigBee.General.Scenes.EnhancedViewScene,
        0x42: ZigBee.General.Scenes.CopyScene,
        0x07: ZigBee.General.Scenes.IkeaRemotePress,
        0x08: ZigBee.General.Scenes.IkeaRemoteLongpressStart,
        0x09: ZigBee.General.Scenes.IkeaRemoteLongpressStop,
    };
    ZigBee.General.Scenes.Client.Command = { 
        0x00: ZigBee.General.Scenes.AddSceneResponse,
        0x01: ZigBee.General.Scenes.ViewSceneResponse,
        0x02: ZigBee.General.Scenes.RemoveSceneResponse,
        0x03: ZigBee.General.Scenes.RemoveAllScenesResponse,
        0x04: ZigBee.General.Scenes.StoreSceneResponse,
        0x06: ZigBee.General.Scenes.GetSceneMembershipResponse,
        0x40: ZigBee.General.Scenes.EnhancedAddSceneResponse,
        0x41: ZigBee.General.Scenes.EnhancedViewSceneResponse,
        0x42: ZigBee.General.Scenes.CopySceneResponse,
    };
    ZigBee.General.OnOff.Server.Attribute = { 
        0x0000: ZigBee.General.Types.OnOff,
        0x4000: ZigBee.General.Types.GlobalSceneControl,
        0x4001: ZigBee.General.Types.OnTime,
        0x4002: ZigBee.General.Types.OffWaitTime,
        0x4003: ZigBee.General.Types.PoweronOnOff,
    };
    ZigBee.General.OnOff.Client.Attribute = { 
    };
    ZigBee.General.OnOff.Server.Command = { 
        0x00: ZigBee.General.OnOff.Off,
        0x01: ZigBee.General.OnOff.On,
        0x02: ZigBee.General.OnOff.Toggle,
        0x40: ZigBee.General.OnOff.OffWithEffect,
        0x41: ZigBee.General.OnOff.OnWithRecallGlobalScene,
        0x42: ZigBee.General.OnOff.OnWithTimedOff,
    };
    ZigBee.General.OnOff.Client.Command = { 
    };
    ZigBee.General.OnOffSwitchConfiguration.Server.Attribute = { 
        0x0000: ZigBee.General.Types.SwitchType,
        0x0010: ZigBee.General.Types.SwitchActions,
    };
    ZigBee.General.OnOffSwitchConfiguration.Client.Attribute = { 
    };
    ZigBee.General.OnOffSwitchConfiguration.Server.Command = { 
    };
    ZigBee.General.OnOffSwitchConfiguration.Client.Command = { 
    };
    ZigBee.General.LevelControl.Server.Attribute = { 
        0x0000: ZigBee.General.Types.CurrentLevel,
        0x0001: ZigBee.General.Types.RemainingTime,
        0x0010: ZigBee.General.Types.OnOffTransistionTime,
        0x0011: ZigBee.General.Types.OnLevel,
        0x0012: ZigBee.General.Types.OnTransitionTime,
        0x0013: ZigBee.General.Types.OffTransitionTime,
        0x0014: ZigBee.General.Types.DefaultMoveRate,
        0x000F: ZigBee.General.Types.LevelControlOptions,
        0x4000: ZigBee.General.Types.PowerOnLevel,
    };
    ZigBee.General.LevelControl.Client.Attribute = { 
    };
    ZigBee.General.LevelControl.Server.Command = { 
        0x00: ZigBee.General.LevelControl.MoveToLevel,
        0x01: ZigBee.General.LevelControl.Move,
        0x02: ZigBee.General.LevelControl.Step,
        0x03: ZigBee.General.LevelControl.Stop,
        0x04: ZigBee.General.LevelControl.MoveToLevelWithOnOff,
        0x05: ZigBee.General.LevelControl.MoveWithOnOff,
        0x06: ZigBee.General.LevelControl.StepWithOnOff,
        0x07: ZigBee.General.LevelControl.StopWithOnOff,
    };
    ZigBee.General.LevelControl.Client.Command = { 
    };
    ZigBee.General.Alarms.Server.Attribute = { 
        0x0000: ZigBee.General.Types.AlarmCount,
    };
    ZigBee.General.Alarms.Client.Attribute = { 
    };
    ZigBee.General.Alarms.Server.Command = { 
        0x00: ZigBee.General.Alarms.ResetAlarm,
        0x01: ZigBee.General.Alarms.ResetAllAlarms,
        0x02: ZigBee.General.Alarms.GetAlarm,
        0x03: ZigBee.General.Alarms.ResetAlarmLog,
    };
    ZigBee.General.Alarms.Client.Command = { 
        0x00: ZigBee.General.Alarms.Alarm,
        0x01: ZigBee.General.Alarms.GetAlarmResponse,
    };
    ZigBee.General.Time.Server.Attribute = { 
        0x0000: ZigBee.General.Types.Time,
        0x0001: ZigBee.General.Types.TimeStatus,
        0x0002: ZigBee.General.Types.TimeZone,
        0x0003: ZigBee.General.Types.DstStart,
        0x0004: ZigBee.General.Types.DstEnd,
        0x0005: ZigBee.General.Types.DstShift,
        0x0006: ZigBee.General.Types.StandardTime,
        0x0007: ZigBee.General.Types.LocalTime,
        0x0008: ZigBee.General.Types.LastSetTime,
        0x0009: ZigBee.General.Types.ValidUntilTime,
    };
    ZigBee.General.Time.Client.Attribute = { 
    };
    ZigBee.General.Time.Server.Command = { 
    };
    ZigBee.General.Time.Client.Command = { 
    };
    ZigBee.General.Location.Server.Attribute = { 
        0x0000: ZigBee.General.Types.LocationType,
        0x0001: ZigBee.General.Types.LocationMethod,
        0x0002: ZigBee.General.Types.LocationAge,
        0x0003: ZigBee.General.Types.QualityMeasure,
        0x0004: ZigBee.General.Types.NumberOfDevices,
        0x0010: ZigBee.General.Types.XCoordinate,
        0x0011: ZigBee.General.Types.YCoordinate,
        0x0012: ZigBee.General.Types.ZCoordinate,
        0x0013: ZigBee.General.Types.Power,
        0x0014: ZigBee.General.Types.PathLossExponent,
        0x0015: ZigBee.General.Types.ReportingPeriod,
        0x0016: ZigBee.General.Types.CalculationPeriod,
        0x0017: ZigBee.General.Types.NumberRssiMeasurements,
    };
    ZigBee.General.Location.Client.Attribute = { 
    };
    ZigBee.General.Location.Server.Command = { 
        0x00: ZigBee.General.Location.SetAbsoluteLocation,
        0x01: ZigBee.General.Location.SetDeviceConfiguration,
        0x02: ZigBee.General.Location.GetDeviceConfiguration,
        0x03: ZigBee.General.Location.GetLocationData,
        0x04: ZigBee.General.Location.RssiResponse,
        0x05: ZigBee.General.Location.SendPings,
        0x06: ZigBee.General.Location.AnchorNodeAnnounce,
        0x40: ZigBee.General.Location.DistanceMeasure,
    };
    ZigBee.General.Location.Client.Command = { 
        0x00: ZigBee.General.Location.DeviceConfigurationResponse,
        0x01: ZigBee.General.Location.LocationDataResponse,
        0x02: ZigBee.General.Location.LocationDataNotification,
        0x03: ZigBee.General.Location.CompactLocationDataNotification,
        0x04: ZigBee.General.Location.RssiPing,
        0x05: ZigBee.General.Location.RssiRequest,
        0x06: ZigBee.General.Location.ReportRssiMeasurements,
        0x07: ZigBee.General.Location.RequestOwnLocation,
        0x40: ZigBee.General.Location.DistanceMeasureResponse,
    };
    ZigBee.General.AnalogInput.Server.Attribute = { 
        0x001C: ZigBee.General.Types.IODescription,
        0x0041: ZigBee.General.Types.AnalogMaxPresentValue,
        0x0045: ZigBee.General.Types.AnalogMinPresentValue,
        0x0051: ZigBee.General.Types.IOOutOfService,
        0x0055: ZigBee.General.Types.AnalogPresentValue,
        0x0067: ZigBee.General.Types.IOReliability,
        0x006A: ZigBee.General.Types.AnalogResolution,
        0x006F: ZigBee.General.Types.IOStatusFlags,
        0x0075: ZigBee.General.Types.IOUnitType,
        0x0100: ZigBee.General.Types.IOApplicationType,
    };
    ZigBee.General.AnalogInput.Client.Attribute = { 
    };
    ZigBee.General.AnalogInput.Server.Command = { 
    };
    ZigBee.General.AnalogInput.Client.Command = { 
    };
    ZigBee.General.AnalogOutput.Server.Attribute = { 
        0x001C: ZigBee.General.Types.IODescription,
        0x0041: ZigBee.General.Types.AnalogMaxPresentValue,
        0x0045: ZigBee.General.Types.AnalogMinPresentValue,
        0x0051: ZigBee.General.Types.IOOutOfService,
        0x0055: ZigBee.General.Types.AnalogPresentValue,
        0x0057: ZigBee.General.Types.AnalogPriorityArray,
        0x0067: ZigBee.General.Types.IOReliability,
        0x0068: ZigBee.General.Types.AnalogRelinquishDefault,
        0x006A: ZigBee.General.Types.AnalogResolution,
        0x006F: ZigBee.General.Types.IOStatusFlags,
        0x0075: ZigBee.General.Types.IOUnitType,
        0x0100: ZigBee.General.Types.IOApplicationType,
    };
    ZigBee.General.AnalogOutput.Client.Attribute = { 
    };
    ZigBee.General.AnalogOutput.Server.Command = { 
    };
    ZigBee.General.AnalogOutput.Client.Command = { 
    };
    ZigBee.General.AnalogValue.Server.Attribute = { 
        0x001C: ZigBee.General.Types.IODescription,
        0x0051: ZigBee.General.Types.IOOutOfService,
        0x0055: ZigBee.General.Types.AnalogPresentValue,
        0x0057: ZigBee.General.Types.AnalogPriorityArray,
        0x0067: ZigBee.General.Types.IOReliability,
        0x0068: ZigBee.General.Types.AnalogRelinquishDefault,
        0x006A: ZigBee.General.Types.AnalogResolution,
        0x006F: ZigBee.General.Types.IOStatusFlags,
        0x0075: ZigBee.General.Types.IOUnitType,
        0x0100: ZigBee.General.Types.IOApplicationType,
    };
    ZigBee.General.AnalogValue.Client.Attribute = { 
    };
    ZigBee.General.AnalogValue.Server.Command = { 
    };
    ZigBee.General.AnalogValue.Client.Command = { 
    };
    ZigBee.General.BinaryInput.Server.Attribute = { 
        0x0004: ZigBee.General.Types.BinaryActiveText,
        0x001C: ZigBee.General.Types.IODescription,
        0x002E: ZigBee.General.Types.BinaryInactiveText,
        0x0051: ZigBee.General.Types.IOOutOfService,
        0x0054: ZigBee.General.Types.BinaryPolarity,
        0x0055: ZigBee.General.Types.BinaryPresentValue,
        0x0067: ZigBee.General.Types.IOReliability,
        0x006F: ZigBee.General.Types.IOStatusFlags,
        0x0100: ZigBee.General.Types.IOApplicationType,
    };
    ZigBee.General.BinaryInput.Client.Attribute = { 
    };
    ZigBee.General.BinaryInput.Server.Command = { 
    };
    ZigBee.General.BinaryInput.Client.Command = { 
    };
    ZigBee.General.BinaryOutput.Server.Attribute = { 
        0x0004: ZigBee.General.Types.BinaryActiveText,
        0x001C: ZigBee.General.Types.IODescription,
        0x002E: ZigBee.General.Types.BinaryInactiveText,
        0x0042: ZigBee.General.Types.BinaryMinOffTime,
        0x0043: ZigBee.General.Types.BinaryMaxOffTime,
        0x0051: ZigBee.General.Types.IOOutOfService,
        0x0054: ZigBee.General.Types.BinaryPolarity,
        0x0055: ZigBee.General.Types.BinaryPresentValue,
        0x0057: ZigBee.General.Types.BinaryPriorityArray,
        0x0067: ZigBee.General.Types.IOReliability,
        0x0068: ZigBee.General.Types.BinaryRelinquishDefault,
        0x006F: ZigBee.General.Types.IOStatusFlags,
        0x0100: ZigBee.General.Types.IOApplicationType,
    };
    ZigBee.General.BinaryOutput.Client.Attribute = { 
    };
    ZigBee.General.BinaryOutput.Server.Command = { 
    };
    ZigBee.General.BinaryOutput.Client.Command = { 
    };
    ZigBee.General.BinaryValue.Server.Attribute = { 
        0x0004: ZigBee.General.Types.BinaryActiveText,
        0x001C: ZigBee.General.Types.IODescription,
        0x002E: ZigBee.General.Types.BinaryInactiveText,
        0x0042: ZigBee.General.Types.BinaryMinOffTime,
        0x0043: ZigBee.General.Types.BinaryMaxOffTime,
        0x0051: ZigBee.General.Types.IOOutOfService,
        0x0055: ZigBee.General.Types.BinaryPresentValue,
        0x0057: ZigBee.General.Types.BinaryPriorityArray,
        0x0067: ZigBee.General.Types.IOReliability,
        0x0068: ZigBee.General.Types.BinaryRelinquishDefault,
        0x006F: ZigBee.General.Types.IOStatusFlags,
        0x0100: ZigBee.General.Types.IOApplicationType,
    };
    ZigBee.General.BinaryValue.Client.Attribute = { 
    };
    ZigBee.General.BinaryValue.Server.Command = { 
    };
    ZigBee.General.BinaryValue.Client.Command = { 
    };
    ZigBee.General.MultistateInput.Server.Attribute = { 
        0x000E: ZigBee.General.Types.MultistateText,
        0x001C: ZigBee.General.Types.IODescription,
        0x004A: ZigBee.General.Types.MultistateNumberOfStates,
        0x0051: ZigBee.General.Types.IOOutOfService,
        0x0055: ZigBee.General.Types.MultistatePresentValue,
        0x0067: ZigBee.General.Types.IOReliability,
        0x006F: ZigBee.General.Types.IOStatusFlags,
        0x0100: ZigBee.General.Types.IOApplicationType,
    };
    ZigBee.General.MultistateInput.Client.Attribute = { 
    };
    ZigBee.General.MultistateInput.Server.Command = { 
    };
    ZigBee.General.MultistateInput.Client.Command = { 
    };
    ZigBee.General.MultistateOutput.Server.Attribute = { 
        0x000E: ZigBee.General.Types.MultistateText,
        0x001C: ZigBee.General.Types.IODescription,
        0x004A: ZigBee.General.Types.MultistateNumberOfStates,
        0x0051: ZigBee.General.Types.IOOutOfService,
        0x0055: ZigBee.General.Types.MultistatePresentValue,
        0x0057: ZigBee.General.Types.MultistatePriorityArray,
        0x0067: ZigBee.General.Types.IOReliability,
        0x0068: ZigBee.General.Types.MultistateRelinquishDefault,
        0x006F: ZigBee.General.Types.IOStatusFlags,
        0x0100: ZigBee.General.Types.IOApplicationType,
    };
    ZigBee.General.MultistateOutput.Client.Attribute = { 
    };
    ZigBee.General.MultistateOutput.Server.Command = { 
    };
    ZigBee.General.MultistateOutput.Client.Command = { 
    };
    ZigBee.General.MultistateValue.Server.Attribute = { 
        0x000E: ZigBee.General.Types.MultistateText,
        0x001C: ZigBee.General.Types.IODescription,
        0x004A: ZigBee.General.Types.MultistateNumberOfStates,
        0x0051: ZigBee.General.Types.IOOutOfService,
        0x0055: ZigBee.General.Types.MultistatePresentValue,
        0x0057: ZigBee.General.Types.MultistatePriorityArray,
        0x0067: ZigBee.General.Types.IOReliability,
        0x0068: ZigBee.General.Types.MultistateRelinquishDefault,
        0x006F: ZigBee.General.Types.IOStatusFlags,
        0x0100: ZigBee.General.Types.IOApplicationType,
    };
    ZigBee.General.MultistateValue.Client.Attribute = { 
    };
    ZigBee.General.MultistateValue.Server.Command = { 
    };
    ZigBee.General.MultistateValue.Client.Command = { 
    };
    ZigBee.General.PollControl.Server.Attribute = { 
        0x0000: ZigBee.General.Types.CheckInInterval,
        0x0001: ZigBee.General.Types.LongPollInterval,
        0x0002: ZigBee.General.Types.ShortPollInterval,
        0x0003: ZigBee.General.Types.FastPollTimeout,
        0x0004: ZigBee.General.Types.CheckInIntervalMin,
        0x0005: ZigBee.General.Types.LongPollIntervalMin,
        0x0006: ZigBee.General.Types.FastPollTimeoutMax,
    };
    ZigBee.General.PollControl.Client.Attribute = { 
    };
    ZigBee.General.PollControl.Server.Command = { 
        0x00: ZigBee.General.PollControl.CheckIn,
    };
    ZigBee.General.PollControl.Client.Command = { 
    };
    ZigBee.General.MeterIdentification.Server.Attribute = { 
    };
    ZigBee.General.MeterIdentification.Client.Attribute = { 
    };
    ZigBee.General.MeterIdentification.Server.Command = { 
    };
    ZigBee.General.MeterIdentification.Client.Command = { 
    };
    ZigBee.General.Diagnostics.Server.Attribute = { 
        0x0000: ZigBee.General.Types.NumberOfResets,
        0x0001: ZigBee.General.Types.PersistensMemoryWrites,
        0x0100: ZigBee.General.Types.MacRxBcast,
        0x0101: ZigBee.General.Types.MacTxBcast,
        0x0102: ZigBee.General.Types.MacRxUcast,
        0x0103: ZigBee.General.Types.MacTxUcast,
        0x0104: ZigBee.General.Types.MacTxUcastRetry,
        0x0105: ZigBee.General.Types.MacTxUcastFail,
        0x0106: ZigBee.General.Types.ApsRxBcast,
        0x0107: ZigBee.General.Types.ApsTxBcast,
        0x0108: ZigBee.General.Types.ApsRxUcast,
        0x0109: ZigBee.General.Types.ApsTxUcastSuccess,
        0x010A: ZigBee.General.Types.ApsTxUcastRetry,
        0x010B: ZigBee.General.Types.ApsTxUcastFail,
        0x010C: ZigBee.General.Types.RouteDiscInitiated,
        0x010D: ZigBee.General.Types.NeighborAdded,
        0x010E: ZigBee.General.Types.NeighborRemoved,
        0x010F: ZigBee.General.Types.NeighborStale,
        0x0110: ZigBee.General.Types.JoinIndication,
        0x0111: ZigBee.General.Types.ChildMoved,
        0x0112: ZigBee.General.Types.NwkFcFailure,
        0x0113: ZigBee.General.Types.ApsFcFailure,
        0x0114: ZigBee.General.Types.ApsUnauthorizedKey,
        0x0115: ZigBee.General.Types.NwkDecryptFailures,
        0x0116: ZigBee.General.Types.ApsDecryptFailures,
        0x0117: ZigBee.General.Types.PacketBufferAllocFailures,
        0x0118: ZigBee.General.Types.RelayedUcast,
        0x0119: ZigBee.General.Types.PhyToMacQueueLimitReached,
        0x011A: ZigBee.General.Types.PacketValidateDropcount,
        0x011B: ZigBee.General.Types.AvgMacRetryPerApsMsgSent,
        0x011C: ZigBee.General.Types.LastMessageLqi,
        0x011D: ZigBee.General.Types.LastMessageRssi,
    };
    ZigBee.General.Diagnostics.Client.Attribute = { 
    };
    ZigBee.General.Diagnostics.Server.Command = { 
    };
    ZigBee.General.Diagnostics.Client.Command = { 
    };
    export const Hvac = {
        Types: { 
            AbsMaxCoolSetpointLimit: makeType<ZigBee.IHvac.IArgAbsMaxCoolSetpointLimit, ZigBee.IHvac.IArgAbsMaxCoolSetpointLimitPayload>(base.s16, ()=>({
                name: `Abs Max Cool Setpoint Limit`,
                description: `absolute maximum level that the cooling setpoint may be set to. This is
a limitation imposed by the manufacturer. The value is calculated as
described in the LocalTemperature attribute`,
                id: 0x0006,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.DegreesCelsius,
                scale: 100,
                
            })),
            AbsMaxHeatSetpointLimit: makeType<ZigBee.IHvac.IArgAbsMaxHeatSetpointLimit, ZigBee.IHvac.IArgAbsMaxHeatSetpointLimitPayload>(base.s16, ()=>({
                name: `Abs Max Heat Setpoint Limit`,
                description: `absolute maximum level that the heating setpoint may be set to. This is[[s]]
a limitation imposed by the manufacturer. The value is calculated as
described in the LocalTemperature attribute`,
                id: 0x0004,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.DegreesCelsius,
                scale: 100,
                
            })),
            AbsMinCoolSetpointLimit: makeType<ZigBee.IHvac.IArgAbsMinCoolSetpointLimit, ZigBee.IHvac.IArgAbsMinCoolSetpointLimitPayload>(base.s16, ()=>({
                name: `Abs Min Cool Setpoint Limit`,
                description: `absolute minimum level that the cooling setpoint may be set to. This is
a limitation imposed by the manufacturer. The value is calculated as
described in the LocalTemperature attribute`,
                id: 0x0005,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.DegreesCelsius,
                scale: 100,
                
            })),
            AbsMinHeatSetpointLimit: makeType<ZigBee.IHvac.IArgAbsMinHeatSetpointLimit, ZigBee.IHvac.IArgAbsMinHeatSetpointLimitPayload>(base.s16, ()=>({
                name: `Abs Min Heat Setpoint Limit`,
                description: `absolute minimum level that the heating setpoint may be set to. This is
a limitation imposed by the manufacturer. The value is calculated as
described in the LocalTemperature attribute`,
                id: 0x0003,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.DegreesCelsius,
                scale: 100,
                
            })),
            AcCapacity: makeType<ZigBee.IHvac.IArgAcCapacity, ZigBee.IHvac.IArgAcCapacityPayload>(base.u16, ()=>({
                name: `AC Capacity`,
                description: `capacity in terms of the format defined by the ACCapacityFormat`,
                id: 0x0041,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.BTUsPerHour,
                
            })),
            AcCapacityFormat: makeType<ZigBee.IHvac.IArgAcCapacityFormat, ZigBee.IHvac.IArgAcCapacityFormatPayload>(base.enum8, ()=>({
                name: `AC Capacity Format`,
                description: ``,
                id: 0x0047,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x00: `BTUh`,  },
                
            })),
            AcCoilTemperature: makeType<ZigBee.IHvac.IArgAcCoilTemperature, ZigBee.IHvac.IArgAcCoilTemperaturePayload>(base.s16, ()=>({
                name: `AC Coil Temperature`,
                description: ``,
                id: 0x0046,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.DegreesCelsius,
                scale: 100,
                
            })),
            AcCompressorType: makeType<ZigBee.IHvac.IArgAcCompressorType, ZigBee.IHvac.IArgAcCompressorTypePayload>(base.enum8, ()=>({
                name: `AC Compressor Type`,
                description: ``,
                id: 0x0043,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x00: `Unknown`, 
                0x01: `T1, Max working ambient 43ºC`, 
                0x02: `T2, Max working ambient 35ºC`, 
                0x03: `T3, Max working ambient 52ºC`,  },
                
            })),
            AcErrorCode: makeType<ZigBee.IHvac.IArgAcErrorCode, ZigBee.IHvac.IArgAcErrorCodePayload>(base.bmp32, ()=>({
                name: `AC Error Code`,
                description: ``,
                id: 0x0044,
                report: false,
                read: true,
                write: true,
                require: false,
                bits: { 
                0: `Compressor failure or refrigerant leakage`, 
                1: `Room temperature sensor failure`, 
                2: `Outdoor temperature sensor failure`, 
                3: `Indoor coil temperature sensor failure`, 
                4: `Fan failure`,  },
                
            })),
            AcLouverPosition: makeType<ZigBee.IHvac.IArgAcLouverPosition, ZigBee.IHvac.IArgAcLouverPositionPayload>(base.enum8, ()=>({
                name: `AC Louver Position`,
                description: ``,
                id: 0x0045,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x01: `Fully closed`, 
                0x02: `Fully open`, 
                0x03: `Quarter open`, 
                0x04: `Half open`, 
                0x05: `Three quarters open`,  },
                
            })),
            AcRefrigerantType: makeType<ZigBee.IHvac.IArgAcRefrigerantType, ZigBee.IHvac.IArgAcRefrigerantTypePayload>(base.enum8, ()=>({
                name: `AC Refrigerant Type`,
                description: `refrigerant used`,
                id: 0x0042,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x00: `Unknown`, 
                0x01: `R22`, 
                0x02: `R410a`, 
                0x03: `R407c`,  },
                
            })),
            AcType: makeType<ZigBee.IHvac.IArgAcType, ZigBee.IHvac.IArgAcTypePayload>(base.enum8, ()=>({
                name: `AC Type`,
                description: `type of Mini Split depending on how Cooling and Heating condition is
achieved`,
                id: 0x0040,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x00: `Unknown`, 
                0x01: `Cooling and fixed speed`, 
                0x02: `Heat pump and fixed speed`, 
                0x03: `Cooling and inverter`, 
                0x04: `Heat pump and inverter`,  },
                
            })),
            Capacity: makeType<ZigBee.IHvac.IArgCapacity, ZigBee.IHvac.IArgCapacityPayload>(base.s16, ()=>({
                name: `Capacity`,
                description: `actual capacity of the pump as a percentage of the effective maximum
setpoint value. It is updated dynamically as the speed of the pump
changes`,
                id: 0x0013,
                report: true,
                read: true,
                write: false,
                require: false,
                unit: units.Percent,
                scale: 200,
                
            })),
            ControlMode: makeType<ZigBee.IHvac.IArgControlMode, ZigBee.IHvac.IArgControlModePayload>(base.enum8, ()=>({
                name: `Control Mode`,
                description: `control mode of the pump`,
                id: 0x0021,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x00: `Constant speed`, 
                0x01: `Constant pressure`, 
                0x02: `Proportional pressure`, 
                0x03: `Constant flow`, 
                0x05: `Constant temperature`, 
                0x07: `Automatic`,  },
                
            })),
            ControlSequenceOfOperation: makeType<ZigBee.IHvac.IArgControlSequenceOfOperation, ZigBee.IHvac.IArgControlSequenceOfOperationPayload>(base.enum8, ()=>({
                name: `Control Sequence Of Operation`,
                description: `overall operating environment of the thermostat, and thus the possible
system modes that the thermostat can operate in`,
                id: 0x001B,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x00: `Cooling only`, 
                0x01: `Cooling with reheat`, 
                0x02: `Heating only`, 
                0x03: `Heating with reheat`, 
                0x04: `Cooling and heating 4-pipes`, 
                0x05: `Cooling and heating 4-pipes with reheat`,  },
                
            })),
            DehumidificationCooling: makeType<ZigBee.IHvac.IArgDehumidificationCooling, ZigBee.IHvac.IArgDehumidificationCoolingPayload>(base.u8, ()=>({
                name: `Dehumidification Cooling`,
                description: `current dehumidification cooling output`,
                id: 0x0001,
                report: true,
                read: true,
                write: false,
                require: false,
                
            })),
            DehumidificationHysteresis: makeType<ZigBee.IHvac.IArgDehumidificationHysteresis, ZigBee.IHvac.IArgDehumidificationHysteresisPayload>(base.u8, ()=>({
                name: `Dehumidification Hysteresis`,
                description: `hysteresis associated with RH`,
                id: 0x0013,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Percent,
                
            })),
            DehumidificationLockout: makeType<ZigBee.IHvac.IArgDehumidificationLockout, ZigBee.IHvac.IArgDehumidificationLockoutPayload>(base.enum8, ()=>({
                name: `Dehumidification Lockout`,
                description: `whether dehumidification is allowed or not`,
                id: 0x0012,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x00: `Denied`, 
                0x01: `Allowed`,  },
                
            })),
            DehumidificationMaxCool: makeType<ZigBee.IHvac.IArgDehumidificationMaxCool, ZigBee.IHvac.IArgDehumidificationMaxCoolPayload>(base.u8, ()=>({
                name: `Dehumidification Max Cool`,
                description: `maximum dehumidification cooling output`,
                id: 0x0014,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Percent,
                
            })),
            EffectiveControlMode: makeType<ZigBee.IHvac.IArgEffectiveControlMode, ZigBee.IHvac.IArgEffectiveControlModePayload>(base.enum8, ()=>({
                name: `Effective Control Mode`,
                description: `control mode that currently applies to the pump. It will have the value
of the ControlMode attribute, unless a remote sensor is used as the
sensor for regulation of the pump. In this case, EffectiveControlMode
will display Constant pressure, Constant flow or Constant temperature
if the remote sensor is a pressure sensor, a flow sensor or a
temperature sensor respectively, regardless of the value of the
ControlMode attribute`,
                id: 0x0012,
                report: false,
                read: true,
                write: false,
                require: false,
                values: { 
                0x00: `Constant speed`, 
                0x01: `Constant pressure`, 
                0x02: `Proportional pressure`, 
                0x03: `Constant flow`, 
                0x05: `Constant temperature`, 
                0x07: `Automatic`,  },
                
            })),
            EffectiveOperationMode: makeType<ZigBee.IHvac.IArgEffectiveOperationMode, ZigBee.IHvac.IArgEffectiveOperationModePayload>(base.enum8, ()=>({
                name: `Effective Operation Mode`,
                description: `current effective operation mode of the pump. The value of the
EffectiveOperationMode attribute is the same as the OperationMode
attribute of the Pump settings attribute set, except when it is
overridden locally`,
                id: 0x0011,
                report: false,
                read: true,
                write: false,
                require: false,
                values: { 
                0x00: `Normal`, 
                0x01: `Minimum`, 
                0x02: `Maximum`, 
                0x03: `Local`,  },
                
            })),
            EmergencyHeatDelta: makeType<ZigBee.IHvac.IArgEmergencyHeatDelta, ZigBee.IHvac.IArgEmergencyHeatDeltaPayload>(base.u8, ()=>({
                name: `Emergency Heat Delta`,
                description: `degrees between LocalTemperature and the OccupiedHeatingSetpoint or
UnoccupiedHeatingSetpoint attributes at which the Thermostat server
will operate in emergency heat mode`,
                id: 0x003A,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.DegreesCelsius,
                scale: 10,
                
            })),
            FanMode: makeType<ZigBee.IHvac.IArgFanMode, ZigBee.IHvac.IArgFanModePayload>(base.enum8, ()=>({
                name: `Fan Mode`,
                description: `current speed of the fan`,
                id: 0x0000,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x00: `Off`, 
                0x01: `Low`, 
                0x02: `Medium`, 
                0x03: `High`, 
                0x04: `On`, 
                0x05: `Auto`, 
                0x06: `Smart (based on occupancy)`,  },
                
            })),
            FanModeSequence: makeType<ZigBee.IHvac.IArgFanModeSequence, ZigBee.IHvac.IArgFanModeSequencePayload>(base.enum8, ()=>({
                name: `Fan Mode Sequence`,
                description: `possible fan speeds that the thermostat can set`,
                id: 0x0001,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x00: `Low/med/high`, 
                0x01: `Low/high`, 
                0x02: `Low/med/high/auto`, 
                0x03: `Low/high/auto`, 
                0x04: `On/auto`,  },
                
            })),
            GetWeeklyDaysToReturn: makeType<ZigBee.IHvac.IArgGetWeeklyDaysToReturn, ZigBee.IHvac.IArgGetWeeklyDaysToReturnPayload>(base.bmp8, ()=>({
                name: `Get Weekly Days To Return`,
                description: ``,
                bits: { 
                0: `Sunday`, 
                1: `Monday`, 
                2: `Tuesday`, 
                3: `Wednesday`, 
                4: `Thursday`, 
                5: `Friday`, 
                6: `Saturday`, 
                7: `Away or vacation`,  },
                
            })),
            GetWeeklyModeToReturn: makeType<ZigBee.IHvac.IArgGetWeeklyModeToReturn, ZigBee.IHvac.IArgGetWeeklyModeToReturnPayload>(base.bmp8, ()=>({
                name: `Get Weekly Mode To Return`,
                description: ``,
                bits: { 
                0: `Heat setpoint`, 
                1: `Cool setpoint`,  },
                
            })),
            HvacSystemTypeConfiguration: makeType<ZigBee.IHvac.IArgHvacSystemTypeConfiguration, ZigBee.IHvac.IArgHvacSystemTypeConfigurationPayload>(base.bmp8, ()=>({
                name: `HVAC System Type Configuration`,
                description: `HVAC system type controlled by the thermostat.
Bit | Description
0-1 | Cooling systemn stage
    | 00 Stage 1
    | 01 Stage 2
    | 10 Stage 3
2-3 | Heating system stage
    | 00 Stage 1
    | 01 Stage 2
    | 10 Stage 3
4   | Heating system type
    | 0 Conventional
    | 1 Heat pump
5   | Heating fuel source
    | 0 Electric
    | 1 Gas`,
                id: 0x0009,
                report: false,
                read: false,
                write: false,
                require: false,
                bits: { 
                1: `Cool stage 2`, 
                2: `Cool stage 3`, 
                3: `Heat stage 2`, 
                4: `Heat stage 3`, 
                5: `Heat pump`, 
                6: `Gas fuel source`,  },
                
            })),
            KeypadLockout: makeType<ZigBee.IHvac.IArgKeypadLockout, ZigBee.IHvac.IArgKeypadLockoutPayload>(base.enum8, ()=>({
                name: `Keypad Lockout`,
                description: `level of functionality that is available to the user via the keypad`,
                id: 0x0001,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x00: `No lockout`, 
                0x01: `Level 1 lockout`, 
                0x02: `Level 2 lockout`, 
                0x03: `Level 3 lockout`, 
                0x04: `Level 4 lockout`, 
                0x05: `Level 5 lockout`,  },
                
            })),
            LifetimeEnergyConsumed: makeType<ZigBee.IHvac.IArgLifetimeEnergyConsumed, ZigBee.IHvac.IArgLifetimeEnergyConsumedPayload>(base.u32, ()=>({
                name: `Lifetime Energy Consumed`,
                description: `accumulated energy consumption of the pump through the entire lifetime
of the pump in kWh. The value of the LifetimeEnergyConsumed attribute
is updated dynamically as the energy consumption of the pump increases.
If LifetimeEnergyConsumed rises above maximum value it rolls over and
starts at 0 (zero)`,
                id: 0x0017,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.KilowattHours,
                
            })),
            LifetimeRunningHours: makeType<ZigBee.IHvac.IArgLifetimeRunningHours, ZigBee.IHvac.IArgLifetimeRunningHoursPayload>(base.u24, ()=>({
                name: `Lifetime Running Hours`,
                description: `ccumulated number of hours that the pump has been powered and the motor
has been running. It is updated dynamically as it increases. It is
preserved over powercycles of the pump. if LifeTimeRunningHours rises
above maximum value it <rolls over= and starts at 0 (zero)`,
                id: 0x0015,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Hours,
                
            })),
            LocalTemperature: makeType<ZigBee.IHvac.IArgLocalTemperature, ZigBee.IHvac.IArgLocalTemperaturePayload>(base.s16, ()=>({
                name: `Local Temperature`,
                description: `temperature in degrees Celsius, as measured locally or remotely (over
the network), including any adjustments applied by
LocalTemperatureCalibration attribute (if any) as follows:
LocalTemperature = 100 x (temperature in degrees Celsius +
LocalTemperatureCalibration)`,
                id: 0x0000,
                report: true,
                read: true,
                write: false,
                require: false,
                unit: units.DegreesCelsius,
                scale: 100,
                
            })),
            LocalTemperatureCalibration: makeType<ZigBee.IHvac.IArgLocalTemperatureCalibration, ZigBee.IHvac.IArgLocalTemperatureCalibrationPayload>(base.s8, ()=>({
                name: `Local Temperature Calibration`,
                description: `offset the thermostat server shall make to the measured temperature
(locally or remotely) before calculating, displaying, or communicating
the LocalTemperature attribute`,
                id: 0x0010,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.DegreesCelsius,
                scale: 10,
                
            })),
            MaxCompPressure: makeType<ZigBee.IHvac.IArgMaxCompPressure, ZigBee.IHvac.IArgMaxCompPressurePayload>(base.s16, ()=>({
                name: `Max Comp Pressure`,
                description: `the maximum compensated pressure the pump can achieve when it is
running and working in control mode Proportional pressure (ControlMode
attribute of the Pump settings attribute set is set to Proportional
pressure)`,
                id: 0x0006,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.Kilopascals,
                scale: 10,
                
            })),
            MaxConstFlow: makeType<ZigBee.IHvac.IArgMaxConstFlow, ZigBee.IHvac.IArgMaxConstFlowPayload>(base.u16, ()=>({
                name: `Max Const Flow`,
                description: `the maximum flow the pump can achieve when it is running and working in
control mode Constant flow (ControlMode attribute of the Pump settings
attribute set is set to Constant flow)`,
                id: 0x000A,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.CubicMetersPerHour,
                scale: 10,
                
            })),
            MaxConstPressure: makeType<ZigBee.IHvac.IArgMaxConstPressure, ZigBee.IHvac.IArgMaxConstPressurePayload>(base.s16, ()=>({
                name: `Max Const Pressure`,
                description: `the maximum pressure the pump can achieve when it is running and
working in control mode constant pressure (ControlMode attribute of the
Pump settings attribute set is set to Constant pressure)`,
                id: 0x0004,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.Kilopascals,
                scale: 10,
                
            })),
            MaxConstSpeed: makeType<ZigBee.IHvac.IArgMaxConstSpeed, ZigBee.IHvac.IArgMaxConstSpeedPayload>(base.u16, ()=>({
                name: `Max Const Speed`,
                description: `the maximum speed the pump can achieve when it is running and working
in control mode Constant speed (ControlMode attribute of the Pump
settings attribute set is set to Constant speed)`,
                id: 0x0008,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.RevolutionsPerMinute,
                
            })),
            MaxConstTemp: makeType<ZigBee.IHvac.IArgMaxConstTemp, ZigBee.IHvac.IArgMaxConstTempPayload>(base.s16, ()=>({
                name: `Max Const Temp`,
                description: `the maximum temperature the pump can maintain in the system when it is
running and working in control mode Constant temperature (ControlMode
attribute of the Pump settings attribute set is set to Constant
temperature)`,
                id: 0x000C,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.DegreesCelsius,
                scale: 100,
                
            })),
            MaxCoolSetpointLimit: makeType<ZigBee.IHvac.IArgMaxCoolSetpointLimit, ZigBee.IHvac.IArgMaxCoolSetpointLimitPayload>(base.s16, ()=>({
                name: `Max Cool Setpoint Limit`,
                description: `maximum level that the cooling setpoint may be set to. It must be less
than or equal to AbsMaxCoolSetpointLimit. If this attribute is not
present, it shall be taken as equal to AbsMaxCoolSetpointLimit`,
                id: 0x0018,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.DegreesCelsius,
                scale: 10,
                
            })),
            MaxFlow: makeType<ZigBee.IHvac.IArgMaxFlow, ZigBee.IHvac.IArgMaxFlowPayload>(base.u16, ()=>({
                name: `Max Flow`,
                description: `the maximum flow the pump can achieve. It is a physical limit, and does
not apply to any specific control mode or operation mode`,
                id: 0x0002,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.CubicMetersPerHour,
                scale: 10,
                
            })),
            MaxHeatSetpointLimit: makeType<ZigBee.IHvac.IArgMaxHeatSetpointLimit, ZigBee.IHvac.IArgMaxHeatSetpointLimitPayload>(base.s16, ()=>({
                name: `Max Heat Setpoint Limit`,
                description: `maximum level that the heating setpoint MAY be set to. It must be less
than or equal to AbsMaxHeatSetpointLimit. If this attribute is not
present, it shall be taken as equal to AbsMaxHeatSetpointLimit`,
                id: 0x0016,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.DegreesCelsius,
                scale: 10,
                
            })),
            MaxPressure: makeType<ZigBee.IHvac.IArgMaxPressure, ZigBee.IHvac.IArgMaxPressurePayload>(base.s16, ()=>({
                name: `Max Pressure`,
                description: `the maximum pressure the pump can achieve. It is a physical limit,
and does not apply to any specific control mode or operation mode`,
                id: 0x0000,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.Kilopascals,
                scale: 10,
                
            })),
            MaxSpeed: makeType<ZigBee.IHvac.IArgMaxSpeed, ZigBee.IHvac.IArgMaxSpeedPayload>(base.u16, ()=>({
                name: `Max Speed`,
                description: `the maximum speed the pump can achieve. It is a physical limit, and
does not apply to any specific control mode or operation mode`,
                id: 0x0001,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.RevolutionsPerMinute,
                
            })),
            MinCompPressure: makeType<ZigBee.IHvac.IArgMinCompPressure, ZigBee.IHvac.IArgMinCompPressurePayload>(base.s16, ()=>({
                name: `Min Comp Pressure`,
                description: `the minimum compensated pressure the pump can achieve when it is
running and working in control mode Proportional pressure (ControlMode
attribute of the Pump settings attribute set is set to Proportional
pressure)`,
                id: 0x0005,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.Kilopascals,
                scale: 10,
                
            })),
            MinConstFlow: makeType<ZigBee.IHvac.IArgMinConstFlow, ZigBee.IHvac.IArgMinConstFlowPayload>(base.u16, ()=>({
                name: `Min Const Flow`,
                description: `the minimum flow the pump can achieve when it is running and working in
control mode Constant flow (ControlMode attribute of the Pump settings
attribute set is set to Constant flow)`,
                id: 0x0009,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.CubicMetersPerHour,
                scale: 10,
                
            })),
            MinConstPressure: makeType<ZigBee.IHvac.IArgMinConstPressure, ZigBee.IHvac.IArgMinConstPressurePayload>(base.s16, ()=>({
                name: `Min Const Pressure`,
                description: `the minimum pressure the pump can achieve when it is running and
working in control mode constant pressure (ControlMode attribute of the
Pump settings attribute set is set to Constant pressure)`,
                id: 0x0003,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.Kilopascals,
                scale: 10,
                
            })),
            MinConstSpeed: makeType<ZigBee.IHvac.IArgMinConstSpeed, ZigBee.IHvac.IArgMinConstSpeedPayload>(base.u16, ()=>({
                name: `Min Const Speed`,
                description: `the minimum speed the pump can achieve when it is running and working
in control mode Constant speed (ControlMode attribute of the Pump
settings attribute set is set to Constant speed)`,
                id: 0x0007,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.RevolutionsPerMinute,
                
            })),
            MinConstTemp: makeType<ZigBee.IHvac.IArgMinConstTemp, ZigBee.IHvac.IArgMinConstTempPayload>(base.s16, ()=>({
                name: `Min Const Temp`,
                description: `the minimum temperature the pump can maintain in the system when it is
running and working in control mode Constant temperature (ControlMode
attribute of the Pump settings attribute set is set to Constant
temperature)`,
                id: 0x000B,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.DegreesCelsius,
                scale: 100,
                
            })),
            MinCoolSetpointLimit: makeType<ZigBee.IHvac.IArgMinCoolSetpointLimit, ZigBee.IHvac.IArgMinCoolSetpointLimitPayload>(base.s16, ()=>({
                name: `Min Cool Setpoint Limit`,
                description: `minimum level that the cooling setpoint may be set to. It must be
greater than or equal to AbsMinCoolSetpointLimit. If this attribute is
not present, it shall be taken as equal to AbsMinCoolSetpointLimit`,
                id: 0x0017,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.DegreesCelsius,
                scale: 10,
                
            })),
            MinHeatSetpointLimit: makeType<ZigBee.IHvac.IArgMinHeatSetpointLimit, ZigBee.IHvac.IArgMinHeatSetpointLimitPayload>(base.s16, ()=>({
                name: `Min Heat Setpoint Limit`,
                description: `minimum level that the heating setpoint may be set to. If this
attribute is not present, it shall be taken as equal to
AbsMinHeatSetpointLimit`,
                id: 0x0015,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.DegreesCelsius,
                scale: 10,
                
            })),
            MinSetpointDeadBand: makeType<ZigBee.IHvac.IArgMinSetpointDeadBand, ZigBee.IHvac.IArgMinSetpointDeadBandPayload>(base.s8, ()=>({
                name: `Min Setpoint Dead Band`,
                description: `minimum difference between the Heat Setpoint and the Cool SetPoint`,
                id: 0x0019,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            NumberOfDailyTransitions: makeType<ZigBee.IHvac.IArgNumberOfDailyTransitions, ZigBee.IHvac.IArgNumberOfDailyTransitionsPayload>(base.u8, ()=>({
                name: `Number Of Daily Transitions`,
                description: `how many daily schedule transitions the thermostat is capable of
handling`,
                id: 0x0022,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            NumberOfWeeklyTransitions: makeType<ZigBee.IHvac.IArgNumberOfWeeklyTransitions, ZigBee.IHvac.IArgNumberOfWeeklyTransitionsPayload>(base.u8, ()=>({
                name: `Number Of Weekly Transitions`,
                description: `how many weekly schedule transitions the thermostat is capable of
handling`,
                id: 0x0021,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            Occupancy: makeType<ZigBee.IHvac.IArgOccupancy, ZigBee.IHvac.IArgOccupancyPayload>(base.bmp8, ()=>({
                name: `Occupancy`,
                description: `whether the heated/cooled space is occupied or not, as measured locally
or remotely (over the network)`,
                id: 0x0002,
                report: false,
                read: true,
                write: false,
                require: false,
                bits: { 
                1: `Occupied`,  },
                
            })),
            OccupiedCoolingSetpoint: makeType<ZigBee.IHvac.IArgOccupiedCoolingSetpoint, ZigBee.IHvac.IArgOccupiedCoolingSetpointPayload>(base.s16, ()=>({
                name: `Occupied Cooling Setpoint`,
                description: `cooling mode setpoint when the room is occupied. The
OccupiedHeatingSetpoint attribute shall always be below the value
specified by at least MinSetpointDeadband`,
                id: 0x0011,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.DegreesCelsius,
                scale: 10,
                
            })),
            OccupiedHeatingSetpoint: makeType<ZigBee.IHvac.IArgOccupiedHeatingSetpoint, ZigBee.IHvac.IArgOccupiedHeatingSetpointPayload>(base.s16, ()=>({
                name: `Occupied Heating Setpoint`,
                description: `heating mode setpoint when the room is occupied. The
OccupiedCoolingSetpoint attribute shall always be above the value
specified by at least MinSetpointDeadband`,
                id: 0x0012,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.DegreesCelsius,
                scale: 10,
                
            })),
            OccupiedSetback: makeType<ZigBee.IHvac.IArgOccupiedSetback, ZigBee.IHvac.IArgOccupiedSetbackPayload>(base.u8, ()=>({
                name: `Occupied Setback`,
                description: `degrees the thermostat will allow the LocalTemperature attribute to
float above the OccupiedCooling setpoint (i.e., OccupiedCooling +
OccupiedSetback) or below the OccupiedHeating setpoint (i.e.,
OccupiedHeating - OccupiedSetback) before initiating a state change to
bring the temperature back to the user's desired setpoint`,
                id: 0x0034,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.DegreesCelsius,
                scale: 10,
                
            })),
            OccupiedSetbackMax: makeType<ZigBee.IHvac.IArgOccupiedSetbackMax, ZigBee.IHvac.IArgOccupiedSetbackMaxPayload>(base.u8, ()=>({
                name: `Occupied Setback Max`,
                description: `degrees the thermostat will allow the OccupiedSetback attribute to be
configured by a user`,
                id: 0x0036,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.DegreesCelsius,
                scale: 10,
                
            })),
            OccupiedSetbackMin: makeType<ZigBee.IHvac.IArgOccupiedSetbackMin, ZigBee.IHvac.IArgOccupiedSetbackMinPayload>(base.u8, ()=>({
                name: `Occupied Setback Min`,
                description: `degrees the thermostat will allow the OccupiedSetback attribute to be
configured by a user`,
                id: 0x0035,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.DegreesCelsius,
                scale: 10,
                
            })),
            OperationMode: makeType<ZigBee.IHvac.IArgOperationMode, ZigBee.IHvac.IArgOperationModePayload>(base.enum8, ()=>({
                name: `Operation Mode`,
                description: `specifies the operation mode of the pump`,
                id: 0x0020,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x00: `Normal`, 
                0x01: `Minimum`, 
                0x02: `Maximum`, 
                0x03: `Local`,  },
                
            })),
            OutdoorTemperature: makeType<ZigBee.IHvac.IArgOutdoorTemperature, ZigBee.IHvac.IArgOutdoorTemperaturePayload>(base.s16, ()=>({
                name: `Outdoor Temperature`,
                description: `outdoor temperature in degrees Celsius, as measured locally or remotely
(over the network). It is measured as described for LocalTemperature`,
                id: 0x0001,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.DegreesCelsius,
                scale: 100,
                
            })),
            PiCoolingDemand: makeType<ZigBee.IHvac.IArgPiCoolingDemand, ZigBee.IHvac.IArgPiCoolingDemandPayload>(base.u8, ()=>({
                name: `PI Cooling Demand`,
                description: `specifies the level of cooling demanded by the PI (proportional
integral) control loop in use by the thermostat (if any), in
percent. This value is 0 when the thermostat is in off or heating mode`,
                id: 0x0007,
                report: true,
                read: true,
                write: false,
                require: false,
                
            })),
            PiHeatingDemand: makeType<ZigBee.IHvac.IArgPiHeatingDemand, ZigBee.IHvac.IArgPiHeatingDemandPayload>(base.u8, ()=>({
                name: `PI Heating Demand`,
                description: `specifies the level of heating demanded by the PI (proportional
integral) control loop in use by the thermostat (if any), in
percent. This value is 0 when the thermostat is in off or cooling mode`,
                id: 0x0008,
                report: true,
                read: true,
                write: false,
                require: false,
                
            })),
            Power: makeType<ZigBee.IHvac.IArgPower, ZigBee.IHvac.IArgPowerPayload>(base.u24, ()=>({
                name: `Power`,
                description: `power consumption of the pump in Watts`,
                id: 0x0016,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Watts,
                
            })),
            PumpAlarmMask: makeType<ZigBee.IHvac.IArgPumpAlarmMask, ZigBee.IHvac.IArgPumpAlarmMaskPayload>(base.bmp16, ()=>({
                name: `Pump Alarm Mask`,
                description: `whether each of the alarms listed is enabled. When the bit number
corresponding to the alarm code is set to 1, the alarm is enabled, else
it is disabled`,
                id: 0x0022,
                report: false,
                read: true,
                write: false,
                require: false,
                bits: { 
                0: `Supply voltage too low`, 
                1: `Supply voltage too high`, 
                10: `Sensor failure`, 
                11: `Electronic non-fatal failure`, 
                12: `Electronic fatal failure`, 
                13: `General fault`, 
                2: `Power missing phase`, 
                3: `System pressure too low`, 
                4: `System pressure too high`, 
                5: `Dry running`, 
                6: `Motor temperature too high`, 
                7: `Pump motor has fatal failure`, 
                8: `Electronic temperature too high`, 
                9: `Pump blocked`,  },
                
            })),
            PumpStatus: makeType<ZigBee.IHvac.IArgPumpStatus, ZigBee.IHvac.IArgPumpStatusPayload>(base.bmp16, ()=>({
                name: `Pump Status`,
                description: ``,
                id: 0x0010,
                report: true,
                read: true,
                write: false,
                require: false,
                bits: { 
                0: `Device fault`, 
                1: `Supply fault`, 
                2: `Speed low`, 
                3: `Speed high`, 
                4: `Local override`, 
                5: `Running`, 
                6: `Remote pressure`, 
                7: `Remote flow`, 
                8: `Remote temperature`,  },
                
            })),
            RelativeHumidity: makeType<ZigBee.IHvac.IArgRelativeHumidity, ZigBee.IHvac.IArgRelativeHumidityPayload>(base.u8, ()=>({
                name: `Relative Humidity`,
                description: `current relative humidity measured by a local or remote sensor`,
                id: 0x0000,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.Percent,
                
            })),
            RelativeHumidityDisplay: makeType<ZigBee.IHvac.IArgRelativeHumidityDisplay, ZigBee.IHvac.IArgRelativeHumidityDisplayPayload>(base.enum8, ()=>({
                name: `Relative Humidity Display`,
                description: `whether the RH value is displayed to the user or not`,
                id: 0x0015,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x00: `Not displayed`, 
                0x01: `Displayed`,  },
                
            })),
            RelativeHumidityMode: makeType<ZigBee.IHvac.IArgRelativeHumidityMode, ZigBee.IHvac.IArgRelativeHumidityModePayload>(base.enum8, ()=>({
                name: `Relative Humidity Mode`,
                description: `how the RelativeHumidity value is being updated`,
                id: 0x0011,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x00: `Locally`, 
                0x01: `Remotely`,  },
                
            })),
            RelayStatus: makeType<ZigBee.IHvac.IArgRelayStatus, ZigBee.IHvac.IArgRelayStatusPayload>(base.bmp8, ()=>({
                name: `Relay Status`,
                description: `status for thermostat when the log is captured. Each bit represents one
relay used by the thermostat. If the bit is on, the associated relay is
on and active. Each thermostat manufacturer can create its own mapping
between the bitmask and the associated relay`,
                
            })),
            RelayStatusHumidity: makeType<ZigBee.IHvac.IArgRelayStatusHumidity, ZigBee.IHvac.IArgRelayStatusHumidityPayload>(base.u8, ()=>({
                name: `Relay Status Humidity`,
                description: `humidity when the log was captured`,
                unit: units.Percent,
                
            })),
            RelayStatusLocalTemperature: makeType<ZigBee.IHvac.IArgRelayStatusLocalTemperature, ZigBee.IHvac.IArgRelayStatusLocalTemperaturePayload>(base.s16, ()=>({
                name: `Relay Status Local Temperature`,
                description: `temperature when the log is captured`,
                unit: units.DegreesCelsius,
                scale: 100,
                
            })),
            RelayStatusLogTimeOfDay: makeType<ZigBee.IHvac.IArgRelayStatusLogTimeOfDay, ZigBee.IHvac.IArgRelayStatusLogTimeOfDayPayload>(base.u16, ()=>({
                name: `Relay Status Log Time of Day`,
                description: `minutes since midnight when the relay status was captured for this
associated log entry`,
                unit: units.Minutes,
                
            })),
            RelayStatusSetpoint: makeType<ZigBee.IHvac.IArgRelayStatusSetpoint, ZigBee.IHvac.IArgRelayStatusSetpointPayload>(base.s16, ()=>({
                name: `Relay Status Setpoint`,
                description: `target setpoint temperature when the log is captured`,
                unit: units.DegreesCelsius,
                scale: 100,
                
            })),
            RelayStatusUnreadEntries: makeType<ZigBee.IHvac.IArgRelayStatusUnreadEntries, ZigBee.IHvac.IArgRelayStatusUnreadEntriesPayload>(base.u16, ()=>({
                name: `Relay status Unread Entries`,
                description: `umber of unread entries within the thermostat log system`,
                
            })),
            RemoteSensing: makeType<ZigBee.IHvac.IArgRemoteSensing, ZigBee.IHvac.IArgRemoteSensingPayload>(base.bmp8, ()=>({
                name: `Remote Sensing`,
                description: `specifies whether the local temperature, outdoor temperature and
occupancy are being sensed by internal sensors or remote networked
sensors. When the bit is 0 it means internal, 1 means remote`,
                id: 0x001A,
                report: false,
                read: true,
                write: true,
                require: false,
                bits: { 
                0: `Local temperature`, 
                1: `Outdoor temperature`, 
                2: `Occupancy`,  },
                
            })),
            RhDehumidificationSetpoint: makeType<ZigBee.IHvac.IArgRhDehumidificationSetpoint, ZigBee.IHvac.IArgRhDehumidificationSetpointPayload>(base.u8, ()=>({
                name: `RH Dehumidification Setpoint`,
                description: `relative humidity at which dehumidification occurs`,
                id: 0x0010,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Percent,
                
            })),
            ScheduleProgrammingVisibility: makeType<ZigBee.IHvac.IArgScheduleProgrammingVisibility, ZigBee.IHvac.IArgScheduleProgrammingVisibilityPayload>(base.enum8, ()=>({
                name: `Schedule Programming Visibility`,
                description: `hide the weekly schedule programming functionality or menu on a
thermostat from a user to prevent local user programming of the weekly
schedule`,
                id: 0x0002,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x00: `Not hidden`, 
                0x01: `Hidden`,  },
                
            })),
            SetWeeklyCoolSetpoint1: makeType<ZigBee.IHvac.IArgSetWeeklyCoolSetpoint1, ZigBee.IHvac.IArgSetWeeklyCoolSetpoint1Payload>(base.s16, ()=>({
                name: `Set Weekly Cool Setpoint 1`,
                description: ``,
                unit: units.DegreesCelsius,
                scale: 100,
                
            })),
            SetWeeklyCoolSetpoint10: makeType<ZigBee.IHvac.IArgSetWeeklyCoolSetpoint10, ZigBee.IHvac.IArgSetWeeklyCoolSetpoint10Payload>(base.s16, ()=>({
                name: `Set Weekly Cool Setpoint 10`,
                description: ``,
                unit: units.DegreesCelsius,
                scale: 100,
                
            })),
            SetWeeklyDayOfWeek: makeType<ZigBee.IHvac.IArgSetWeeklyDayOfWeek, ZigBee.IHvac.IArgSetWeeklyDayOfWeekPayload>(base.bmp8, ()=>({
                name: `Set Weekly Day Of Week`,
                description: `day of the week at which all the transitions within the payload of the
command should be associated to`,
                bits: { 
                0: `Sunday`, 
                1: `Monday`, 
                2: `Tuesday`, 
                3: `Wednesday`, 
                4: `Thursday`, 
                5: `Friday`, 
                6: `Saturday`, 
                7: `Away or vacation`,  },
                
            })),
            SetWeeklyHeatSetpoint1: makeType<ZigBee.IHvac.IArgSetWeeklyHeatSetpoint1, ZigBee.IHvac.IArgSetWeeklyHeatSetpoint1Payload>(base.s16, ()=>({
                name: `Set Weekly Heat Setpoint 1`,
                description: ``,
                unit: units.DegreesCelsius,
                scale: 100,
                
            })),
            SetWeeklyHeatSetpoint10: makeType<ZigBee.IHvac.IArgSetWeeklyHeatSetpoint10, ZigBee.IHvac.IArgSetWeeklyHeatSetpoint10Payload>(base.s16, ()=>({
                name: `Set Weekly Heat Setpoint 10`,
                description: ``,
                unit: units.DegreesCelsius,
                scale: 100,
                
            })),
            SetWeeklyMode: makeType<ZigBee.IHvac.IArgSetWeeklyMode, ZigBee.IHvac.IArgSetWeeklyModePayload>(base.bmp8, ()=>({
                name: `Set Weekly Mode`,
                description: `which type of setpoint transition is present in the rest of the
command`,
                bits: { 
                0: `Heat setpoint`, 
                1: `Cool setpoint`,  },
                
            })),
            SetWeeklyNumberOfTransitions: makeType<ZigBee.IHvac.IArgSetWeeklyNumberOfTransitions, ZigBee.IHvac.IArgSetWeeklyNumberOfTransitionsPayload>(base.u8, ()=>({
                name: `Set Weekly Number Of Transitions`,
                description: `how many individual transitions to expect for this sequence of
commands. If a device supports more than 10 transitions in its
schedule they can send this by sending more than 1 Set Weekly Schedule
command, each containing the separate information that the device
needs to set`,
                
            })),
            SetWeeklyTransitionTime1: makeType<ZigBee.IHvac.IArgSetWeeklyTransitionTime1, ZigBee.IHvac.IArgSetWeeklyTransitionTime1Payload>(base.u16, ()=>({
                name: `Set Weekly Transition Time 1`,
                description: ``,
                unit: units.Minutes,
                
            })),
            SetWeeklyTransitionTime10: makeType<ZigBee.IHvac.IArgSetWeeklyTransitionTime10, ZigBee.IHvac.IArgSetWeeklyTransitionTime10Payload>(base.u16, ()=>({
                name: `Set Weekly Transition Time 10`,
                description: ``,
                unit: units.Minutes,
                
            })),
            SetpointAmount: makeType<ZigBee.IHvac.IArgSetpointAmount, ZigBee.IHvac.IArgSetpointAmountPayload>(base.s8, ()=>({
                name: `Setpoint Amount`,
                description: ``,
                unit: units.DegreesCelsius,
                scale: 10,
                
            })),
            SetpointChangeAmount: makeType<ZigBee.IHvac.IArgSetpointChangeAmount, ZigBee.IHvac.IArgSetpointChangeAmountPayload>(base.s16, ()=>({
                name: `Setpoint Change Amount`,
                description: `delta between the current active OccupiedCoolingSetpoint or
OccupiedHeatingSetpoint and the previous active setpoint`,
                id: 0x0031,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.DegreesCelsius,
                scale: 100,
                
            })),
            SetpointChangeSource: makeType<ZigBee.IHvac.IArgSetpointChangeSource, ZigBee.IHvac.IArgSetpointChangeSourcePayload>(base.enum8, ()=>({
                name: `Setpoint Change Source`,
                description: `source of the current active OccupiedCoolingSetpoint or
OccupiedHeatingSetpoint (i.e., who or what determined the current
setpoint)`,
                id: 0x0030,
                report: false,
                read: true,
                write: false,
                require: false,
                values: { 
                0x00: `Manual/user initiated`, 
                0x01: `Scheduling/programming initiated`, 
                0x02: `Externally initiated by cluster command or attribute write`,  },
                
            })),
            SetpointChangeSourceTimestamp: makeType<ZigBee.IHvac.IArgSetpointChangeSourceTimestamp, ZigBee.IHvac.IArgSetpointChangeSourceTimestampPayload>(base.utc, ()=>({
                name: `Setpoint Change Source Timestamp`,
                description: `time in UTC at which the SetpointChangeSourceAmount attribute change
was recorded`,
                id: 0x0032,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            SetpointMode: makeType<ZigBee.IHvac.IArgSetpointMode, ZigBee.IHvac.IArgSetpointModePayload>(base.enum8, ()=>({
                name: `Setpoint Mode`,
                description: ``,
                values: { 
                0x00: `Adjust heat setpoint`, 
                0x01: `Adjust cool setpoint`, 
                0x02: `Adjust heat and cool setpoint`,  },
                
            })),
            Speed: makeType<ZigBee.IHvac.IArgSpeed, ZigBee.IHvac.IArgSpeedPayload>(base.u16, ()=>({
                name: `Speed`,
                description: `actual speed of the pump measured in RPM. It is updated dynamically as
the speed of the pump changes`,
                id: 0x0014,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.RevolutionsPerMinute,
                
            })),
            StartOfWeek: makeType<ZigBee.IHvac.IArgStartOfWeek, ZigBee.IHvac.IArgStartOfWeekPayload>(base.enum8, ()=>({
                name: `Start Of Week`,
                description: `day of the week that this thermostat considers to be the start of week
for weekly set point scheduling. This attribute may be able to be used
as the base to determine if the device supports weekly scheduling by
reading the attribute. Successful response means that the weekly
scheduling is supported`,
                id: 0x0020,
                report: false,
                read: true,
                write: false,
                require: false,
                values: { 
                0x00: `Sunday`, 
                0x01: `Monday`, 
                0x02: `Tuesday`, 
                0x03: `Wednesday`, 
                0x04: `Thursday`, 
                0x05: `Friday`, 
                0x06: `Saturday`,  },
                
            })),
            SystemMode: makeType<ZigBee.IHvac.IArgSystemMode, ZigBee.IHvac.IArgSystemModePayload>(base.enum8, ()=>({
                name: `System Mode`,
                description: `specifies the current operating mode of the thermostat`,
                id: 0x001C,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x00: `Off`, 
                0x01: `Auto`, 
                0x03: `Cool`, 
                0x04: `Heat`, 
                0x05: `Emergency heating`, 
                0x06: `Precooling`, 
                0x07: `Fan only`, 
                0x08: `Dry`, 
                0x09: `Sleep`,  },
                
            })),
            TemperatureDisplayMode: makeType<ZigBee.IHvac.IArgTemperatureDisplayMode, ZigBee.IHvac.IArgTemperatureDisplayModePayload>(base.enum8, ()=>({
                name: `Temperature Display Mode`,
                description: `units of the temperature displayed on the thermostat screen`,
                id: 0x0000,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x00: `Temperature in Celsius`, 
                0x01: `Temperature in Fahrenheit`,  },
                
            })),
            TemperatureSetpointHold: makeType<ZigBee.IHvac.IArgTemperatureSetpointHold, ZigBee.IHvac.IArgTemperatureSetpointHoldPayload>(base.enum8, ()=>({
                name: `Temperature Setpoint Hold`,
                description: `temperature hold status on the thermostat. If hold status is on, the
thermostat should maintain the temperature set point for the current
mode until a system mode change. If hold status is off, the thermostat
should follow the setpoint transitions specified by its internal
scheduling program. If the thermostat supports setpoint hold for a
specific duration, it should also implement the
TemperatureSetpointHoldDuration attribute`,
                id: 0x0023,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x00: `Off`, 
                0x01: `On`,  },
                
            })),
            TemperatureSetpointHoldDuration: makeType<ZigBee.IHvac.IArgTemperatureSetpointHoldDuration, ZigBee.IHvac.IArgTemperatureSetpointHoldDurationPayload>(base.u16, ()=>({
                name: `Temperature Setpoint Hold Duration`,
                description: `period in minutes for which a setpoint hold is active`,
                id: 0x0024,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Minutes,
                
            })),
            ThermostatAlarmMask: makeType<ZigBee.IHvac.IArgThermostatAlarmMask, ZigBee.IHvac.IArgThermostatAlarmMaskPayload>(base.bmp8, ()=>({
                name: `Thermostat Alarm Mask`,
                description: `specifies whether each of the alarms listed is enabled. When the bit
number corresponding to the alarm code is set to 1, the alarm is
enabled, else it is disabled`,
                id: 0x001D,
                report: false,
                read: true,
                write: false,
                require: false,
                bits: { 
                0: `Initialisation failure`, 
                1: `Hardware failure`, 
                2: `Self-calibration failure`,  },
                
            })),
            ThermostatProgrammingOperationMode: makeType<ZigBee.IHvac.IArgThermostatProgrammingOperationMode, ZigBee.IHvac.IArgThermostatProgrammingOperationModePayload>(base.bmp8, ()=>({
                name: `Thermostat Programming Operation Mode`,
                description: `operational state of the thermostat's programming. The thermostat shall
modify its programming operation when this attribute is modified by a
client and update this attribute when its programming operation is
modified locally by a user. When a bit is 0 it means off, 1 means on.
For the scheduling mode bit, 0 means the thermostate is manually
controlled, whereas 1 means it is following a programmed weekly
schedule`,
                id: 0x0025,
                report: true,
                read: true,
                write: true,
                require: false,
                bits: { 
                0: `Scheduling mode`, 
                1: `Auto/recovery mode`, 
                2: `Economy mode`,  },
                
            })),
            ThermostatRunningMode: makeType<ZigBee.IHvac.IArgThermostatRunningMode, ZigBee.IHvac.IArgThermostatRunningModePayload>(base.enum8, ()=>({
                name: `Thermostat Running Mode`,
                description: `represents the running mode of the thermostat`,
                id: 0x001E,
                report: false,
                read: true,
                write: false,
                require: false,
                values: { 
                0x00: `Off`, 
                0x03: `Cool`, 
                0x04: `Heat`,  },
                
            })),
            ThermostatRunningState: makeType<ZigBee.IHvac.IArgThermostatRunningState, ZigBee.IHvac.IArgThermostatRunningStatePayload>(base.bmp16, ()=>({
                name: `Thermostat Running State`,
                description: ``,
                id: 0x0029,
                report: false,
                read: true,
                write: false,
                require: false,
                bits: { 
                0: `Heat stage on`, 
                1: `Cool stage on`, 
                2: `Fan stage on`, 
                3: `Heat second stage on`, 
                4: `Cool second stage on`, 
                5: `Fan second stage on`, 
                6: `Fan third stage on`,  },
                
            })),
            UnoccupiedCoolingSetpoint: makeType<ZigBee.IHvac.IArgUnoccupiedCoolingSetpoint, ZigBee.IHvac.IArgUnoccupiedCoolingSetpointPayload>(base.s16, ()=>({
                name: `Unoccupied Cooling Setpoint`,
                description: `cooling mode setpoint when the room is unoccupied. The
UnoccupiedHeatingSetpoint attribute shall always be below the value
specified by at least MinSetpointDeadband`,
                id: 0x0013,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.DegreesCelsius,
                scale: 10,
                
            })),
            UnoccupiedHeatingSetpoint: makeType<ZigBee.IHvac.IArgUnoccupiedHeatingSetpoint, ZigBee.IHvac.IArgUnoccupiedHeatingSetpointPayload>(base.s16, ()=>({
                name: `Unoccupied Heating Setpoint`,
                description: `heating mode setpoint when the room is unoccupied. The
UnoccupiedCoolingSetpoint attribute shall always be above the value
specified by at least MinSetpointDeadband`,
                id: 0x0014,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.DegreesCelsius,
                scale: 10,
                
            })),
            UnoccupiedSetback: makeType<ZigBee.IHvac.IArgUnoccupiedSetback, ZigBee.IHvac.IArgUnoccupiedSetbackPayload>(base.u8, ()=>({
                name: `Unoccupied Setback`,
                description: `degrees the thermostat will allow the LocalTemperature attribute to
float above the UnoccupiedCooling setpoint (i.e., UnoccupiedCooling +
UnoccupiedSetback) or below the UnoccupiedHeating setpoint (i.e.,
UnoccupiedHeating - UnoccupiedSetback) before initiating a state change
to bring the temperature back to the user's desired setpoint`,
                id: 0x0037,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.DegreesCelsius,
                scale: 10,
                
            })),
            UnoccupiedSetbackMax: makeType<ZigBee.IHvac.IArgUnoccupiedSetbackMax, ZigBee.IHvac.IArgUnoccupiedSetbackMaxPayload>(base.u8, ()=>({
                name: `Unoccupied Setback Max`,
                description: `degrees the thermostat will allow the UnoccupiedSetback attribute to be
configured by a user`,
                id: 0x0039,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.DegreesCelsius,
                scale: 10,
                
            })),
            UnoccupiedSetbackMin: makeType<ZigBee.IHvac.IArgUnoccupiedSetbackMin, ZigBee.IHvac.IArgUnoccupiedSetbackMinPayload>(base.u8, ()=>({
                name: `Unoccupied Setback Min`,
                description: `degrees the thermostat will allow the UnoccupiedSetback attribute to be
configured by a user`,
                id: 0x0038,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.DegreesCelsius,
                scale: 10,
                
            })), },
        PumpConfigurationAndControl: {
            ID: 0x0200,
            Name: `Pump Configuration and Control`,
            Desc: `provides attributes and commands for configuring and
controlling pumps`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        Thermostat: {
            ID: 0x0201,
            Name: `Thermostat`,
            Desc: `provides attributes and commands for configuring and
controlling thermostats`,
            
            SetpointRaiseLower: makeType<ZigBee.IHvac.Thermostat.ICmdSetpointRaiseLower, ZigBee.IHvac.Thermostat.ICmdSetpointRaiseLowerPayload>(command, () => ({
                name: `Setpoint raise/lower`,
                description: ``,
                id: 0x0000,
                payload: { 
                    SetpointMode: ZigBee.Hvac.Types.SetpointMode,
                    SetpointAmount: ZigBee.Hvac.Types.SetpointAmount,
                }
            })),

            SetWeeklySchedule: makeType<ZigBee.IHvac.Thermostat.ICmdSetWeeklySchedule, ZigBee.IHvac.Thermostat.ICmdSetWeeklySchedulePayload>(command, () => ({
                name: `Set weekly schedule`,
                description: ``,
                id: 0x0001,
                payload: { 
                    SetWeeklyNumberOfTransitions: ZigBee.Hvac.Types.SetWeeklyNumberOfTransitions,
                    SetWeeklyDayOfWeek: ZigBee.Hvac.Types.SetWeeklyDayOfWeek,
                    SetWeeklyMode: ZigBee.Hvac.Types.SetWeeklyMode,
                    SetWeeklyTransitionTime1: ZigBee.Hvac.Types.SetWeeklyTransitionTime1,
                    SetWeeklyHeatSetpoint1: ZigBee.Hvac.Types.SetWeeklyHeatSetpoint1,
                    SetWeeklyCoolSetpoint1: ZigBee.Hvac.Types.SetWeeklyCoolSetpoint1,
                    SetWeeklyTransitionTime10: ZigBee.Hvac.Types.SetWeeklyTransitionTime10,
                    SetWeeklyHeatSetpoint10: ZigBee.Hvac.Types.SetWeeklyHeatSetpoint10,
                    SetWeeklyCoolSetpoint10: ZigBee.Hvac.Types.SetWeeklyCoolSetpoint10,
                }
            })),

            GetWeeklySchedule: makeType<ZigBee.IHvac.Thermostat.ICmdGetWeeklySchedule, ZigBee.IHvac.Thermostat.ICmdGetWeeklySchedulePayload>(command, () => ({
                name: `Get weekly schedule`,
                description: ``,
                id: 0x0002,
                payload: { 
                    GetWeeklyDaysToReturn: ZigBee.Hvac.Types.GetWeeklyDaysToReturn,
                    GetWeeklyModeToReturn: ZigBee.Hvac.Types.GetWeeklyModeToReturn,
                }
            })),

            ClearWeeklySchedule: makeType<ZigBee.IHvac.Thermostat.ICmdClearWeeklySchedule, ZigBee.IHvac.Thermostat.ICmdClearWeeklySchedulePayload>(command, () => ({
                name: `Clear weekly schedule`,
                description: ``,
                id: 0x0003,
                payload: {}
            })),

            GetRelayStatusLog: makeType<ZigBee.IHvac.Thermostat.ICmdGetRelayStatusLog, ZigBee.IHvac.Thermostat.ICmdGetRelayStatusLogPayload>(command, () => ({
                name: `Get relay status log`,
                description: ``,
                id: 0x0004,
                payload: {}
            })),

            
            GetWeeklyScheduleResponse: makeType<ZigBee.IHvac.Thermostat.ICmdGetWeeklyScheduleResponse, ZigBee.IHvac.Thermostat.ICmdGetWeeklyScheduleResponsePayload>(command, () => ({
                name: `Get weekly schedule response`,
                description: ``,
                id: 0x0000,
                payload: { 
                    SetWeeklyNumberOfTransitions: ZigBee.Hvac.Types.SetWeeklyNumberOfTransitions,
                    SetWeeklyDayOfWeek: ZigBee.Hvac.Types.SetWeeklyDayOfWeek,
                    SetWeeklyMode: ZigBee.Hvac.Types.SetWeeklyMode,
                    SetWeeklyTransitionTime1: ZigBee.Hvac.Types.SetWeeklyTransitionTime1,
                    SetWeeklyHeatSetpoint1: ZigBee.Hvac.Types.SetWeeklyHeatSetpoint1,
                    SetWeeklyCoolSetpoint1: ZigBee.Hvac.Types.SetWeeklyCoolSetpoint1,
                    SetWeeklyTransitionTime10: ZigBee.Hvac.Types.SetWeeklyTransitionTime10,
                    SetWeeklyHeatSetpoint10: ZigBee.Hvac.Types.SetWeeklyHeatSetpoint10,
                    SetWeeklyCoolSetpoint10: ZigBee.Hvac.Types.SetWeeklyCoolSetpoint10,
                }
            })),

            GetRelayStatusLogResponse: makeType<ZigBee.IHvac.Thermostat.ICmdGetRelayStatusLogResponse, ZigBee.IHvac.Thermostat.ICmdGetRelayStatusLogResponsePayload>(command, () => ({
                name: `Get relay status log response`,
                description: ``,
                id: 0x0001,
                payload: { 
                    RelayStatusLogTimeOfDay: ZigBee.Hvac.Types.RelayStatusLogTimeOfDay,
                    RelayStatus: ZigBee.Hvac.Types.RelayStatus,
                    RelayStatusLocalTemperature: ZigBee.Hvac.Types.RelayStatusLocalTemperature,
                    RelayStatusHumidity: ZigBee.Hvac.Types.RelayStatusHumidity,
                    RelayStatusSetpoint: ZigBee.Hvac.Types.RelayStatusSetpoint,
                    RelayStatusUnreadEntries: ZigBee.Hvac.Types.RelayStatusUnreadEntries,
                }
            })),

            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        FanControl: {
            ID: 0x0202,
            Name: `Fan Control`,
            Desc: `provides attributes and commands for configuring fans
in heating/cooling systems`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        DehumidificationControl: {
            ID: 0x0203,
            Name: `Dehumidification Control`,
            Desc: `provides attributes and commands for configuring
dehumidification appliances`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        ThermostatUiConfiguration: {
            ID: 0x0204,
            Name: `Thermostat UI configuration`,
            Desc: `provides attributes and commands for configuring
the UI of a (remote) thermostat`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        }
    };
    
    ZigBee.Hvac.PumpConfigurationAndControl.Server.Attribute = { 
        0x0000: ZigBee.Hvac.Types.MaxPressure,
        0x0001: ZigBee.Hvac.Types.MaxSpeed,
        0x0002: ZigBee.Hvac.Types.MaxFlow,
        0x0003: ZigBee.Hvac.Types.MinConstPressure,
        0x0004: ZigBee.Hvac.Types.MaxConstPressure,
        0x0005: ZigBee.Hvac.Types.MinCompPressure,
        0x0006: ZigBee.Hvac.Types.MaxCompPressure,
        0x0007: ZigBee.Hvac.Types.MinConstSpeed,
        0x0008: ZigBee.Hvac.Types.MaxConstSpeed,
        0x0009: ZigBee.Hvac.Types.MinConstFlow,
        0x000A: ZigBee.Hvac.Types.MaxConstFlow,
        0x000B: ZigBee.Hvac.Types.MinConstTemp,
        0x000C: ZigBee.Hvac.Types.MaxConstTemp,
        0x0010: ZigBee.Hvac.Types.PumpStatus,
        0x0011: ZigBee.Hvac.Types.EffectiveOperationMode,
        0x0012: ZigBee.Hvac.Types.EffectiveControlMode,
        0x0013: ZigBee.Hvac.Types.Capacity,
        0x0014: ZigBee.Hvac.Types.Speed,
        0x0015: ZigBee.Hvac.Types.LifetimeRunningHours,
        0x0016: ZigBee.Hvac.Types.Power,
        0x0017: ZigBee.Hvac.Types.LifetimeEnergyConsumed,
        0x0020: ZigBee.Hvac.Types.OperationMode,
        0x0021: ZigBee.Hvac.Types.ControlMode,
        0x0022: ZigBee.Hvac.Types.PumpAlarmMask,
    };
    ZigBee.Hvac.PumpConfigurationAndControl.Client.Attribute = { 
    };
    ZigBee.Hvac.PumpConfigurationAndControl.Server.Command = { 
    };
    ZigBee.Hvac.PumpConfigurationAndControl.Client.Command = { 
    };
    ZigBee.Hvac.Thermostat.Server.Attribute = { 
        0x0000: ZigBee.Hvac.Types.LocalTemperature,
        0x0001: ZigBee.Hvac.Types.OutdoorTemperature,
        0x0002: ZigBee.Hvac.Types.Occupancy,
        0x0003: ZigBee.Hvac.Types.AbsMinHeatSetpointLimit,
        0x0004: ZigBee.Hvac.Types.AbsMaxHeatSetpointLimit,
        0x0005: ZigBee.Hvac.Types.AbsMinCoolSetpointLimit,
        0x0006: ZigBee.Hvac.Types.AbsMaxCoolSetpointLimit,
        0x0007: ZigBee.Hvac.Types.PiCoolingDemand,
        0x0008: ZigBee.Hvac.Types.PiHeatingDemand,
        0x0009: ZigBee.Hvac.Types.HvacSystemTypeConfiguration,
        0x0010: ZigBee.Hvac.Types.LocalTemperatureCalibration,
        0x0011: ZigBee.Hvac.Types.OccupiedCoolingSetpoint,
        0x0012: ZigBee.Hvac.Types.OccupiedHeatingSetpoint,
        0x0013: ZigBee.Hvac.Types.UnoccupiedCoolingSetpoint,
        0x0014: ZigBee.Hvac.Types.UnoccupiedHeatingSetpoint,
        0x0015: ZigBee.Hvac.Types.MinHeatSetpointLimit,
        0x0016: ZigBee.Hvac.Types.MaxHeatSetpointLimit,
        0x0017: ZigBee.Hvac.Types.MinCoolSetpointLimit,
        0x0018: ZigBee.Hvac.Types.MaxCoolSetpointLimit,
        0x0019: ZigBee.Hvac.Types.MinSetpointDeadBand,
        0x001A: ZigBee.Hvac.Types.RemoteSensing,
        0x001B: ZigBee.Hvac.Types.ControlSequenceOfOperation,
        0x001C: ZigBee.Hvac.Types.SystemMode,
        0x001D: ZigBee.Hvac.Types.ThermostatAlarmMask,
        0x001E: ZigBee.Hvac.Types.ThermostatRunningMode,
        0x0020: ZigBee.Hvac.Types.StartOfWeek,
        0x0021: ZigBee.Hvac.Types.NumberOfWeeklyTransitions,
        0x0022: ZigBee.Hvac.Types.NumberOfDailyTransitions,
        0x0023: ZigBee.Hvac.Types.TemperatureSetpointHold,
        0x0024: ZigBee.Hvac.Types.TemperatureSetpointHoldDuration,
        0x0025: ZigBee.Hvac.Types.ThermostatProgrammingOperationMode,
        0x0029: ZigBee.Hvac.Types.ThermostatRunningState,
        0x0030: ZigBee.Hvac.Types.SetpointChangeSource,
        0x0031: ZigBee.Hvac.Types.SetpointChangeAmount,
        0x0032: ZigBee.Hvac.Types.SetpointChangeSourceTimestamp,
        0x0034: ZigBee.Hvac.Types.OccupiedSetback,
        0x0035: ZigBee.Hvac.Types.OccupiedSetbackMin,
        0x0036: ZigBee.Hvac.Types.OccupiedSetbackMax,
        0x0037: ZigBee.Hvac.Types.UnoccupiedSetback,
        0x0038: ZigBee.Hvac.Types.UnoccupiedSetbackMin,
        0x0039: ZigBee.Hvac.Types.UnoccupiedSetbackMax,
        0x003A: ZigBee.Hvac.Types.EmergencyHeatDelta,
        0x0040: ZigBee.Hvac.Types.AcType,
        0x0041: ZigBee.Hvac.Types.AcCapacity,
        0x0042: ZigBee.Hvac.Types.AcRefrigerantType,
        0x0043: ZigBee.Hvac.Types.AcCompressorType,
        0x0044: ZigBee.Hvac.Types.AcErrorCode,
        0x0045: ZigBee.Hvac.Types.AcLouverPosition,
        0x0046: ZigBee.Hvac.Types.AcCoilTemperature,
        0x0047: ZigBee.Hvac.Types.AcCapacityFormat,
    };
    ZigBee.Hvac.Thermostat.Client.Attribute = { 
    };
    ZigBee.Hvac.Thermostat.Server.Command = { 
        0x00: ZigBee.Hvac.Thermostat.SetpointRaiseLower,
        0x01: ZigBee.Hvac.Thermostat.SetWeeklySchedule,
        0x02: ZigBee.Hvac.Thermostat.GetWeeklySchedule,
        0x03: ZigBee.Hvac.Thermostat.ClearWeeklySchedule,
        0x04: ZigBee.Hvac.Thermostat.GetRelayStatusLog,
    };
    ZigBee.Hvac.Thermostat.Client.Command = { 
        0x00: ZigBee.Hvac.Thermostat.GetWeeklyScheduleResponse,
        0x01: ZigBee.Hvac.Thermostat.GetRelayStatusLogResponse,
    };
    ZigBee.Hvac.FanControl.Server.Attribute = { 
        0x0000: ZigBee.Hvac.Types.FanMode,
        0x0001: ZigBee.Hvac.Types.FanModeSequence,
    };
    ZigBee.Hvac.FanControl.Client.Attribute = { 
    };
    ZigBee.Hvac.FanControl.Server.Command = { 
    };
    ZigBee.Hvac.FanControl.Client.Command = { 
    };
    ZigBee.Hvac.DehumidificationControl.Server.Attribute = { 
        0x0000: ZigBee.Hvac.Types.RelativeHumidity,
        0x0001: ZigBee.Hvac.Types.DehumidificationCooling,
        0x0010: ZigBee.Hvac.Types.RhDehumidificationSetpoint,
        0x0011: ZigBee.Hvac.Types.RelativeHumidityMode,
        0x0012: ZigBee.Hvac.Types.DehumidificationLockout,
        0x0013: ZigBee.Hvac.Types.DehumidificationHysteresis,
        0x0014: ZigBee.Hvac.Types.DehumidificationMaxCool,
        0x0015: ZigBee.Hvac.Types.RelativeHumidityDisplay,
    };
    ZigBee.Hvac.DehumidificationControl.Client.Attribute = { 
    };
    ZigBee.Hvac.DehumidificationControl.Server.Command = { 
    };
    ZigBee.Hvac.DehumidificationControl.Client.Command = { 
    };
    ZigBee.Hvac.ThermostatUiConfiguration.Server.Attribute = { 
        0x0000: ZigBee.Hvac.Types.TemperatureDisplayMode,
        0x0001: ZigBee.Hvac.Types.KeypadLockout,
        0x0002: ZigBee.Hvac.Types.ScheduleProgrammingVisibility,
    };
    ZigBee.Hvac.ThermostatUiConfiguration.Client.Attribute = { 
    };
    ZigBee.Hvac.ThermostatUiConfiguration.Server.Command = { 
    };
    ZigBee.Hvac.ThermostatUiConfiguration.Client.Command = { 
    };
    export const SecurityAndSafety = {
        Types: { 
            CurrentZoneSensitivityLevel: makeType<ZigBee.ISecurityAndSafety.IArgCurrentZoneSensitivityLevel, ZigBee.ISecurityAndSafety.IArgCurrentZoneSensitivityLevelPayload>(base.u8, ()=>({
                name: `Current Zone Sensitivity Level`,
                description: ``,
                id: 0x0013,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            Delay: makeType<ZigBee.ISecurityAndSafety.IArgDelay, ZigBee.ISecurityAndSafety.IArgDelayPayload>(base.u16, ()=>({
                name: `Delay`,
                description: ``,
                unit: units.Seconds,
                scale: 4,
                
            })),
            EnrollResponseCode: makeType<ZigBee.ISecurityAndSafety.IArgEnrollResponseCode, ZigBee.ISecurityAndSafety.IArgEnrollResponseCodePayload>(base.enum8, ()=>({
                name: `Enroll response code`,
                description: ``,
                values: { 
                0x00: `Success`, 
                0x01: `Not Supported`, 
                0x02: `Not permitted`, 
                0x03: `Too many zones`,  },
                
            })),
            ExtendedStatus: makeType<ZigBee.ISecurityAndSafety.IArgExtendedStatus, ZigBee.ISecurityAndSafety.IArgExtendedStatusPayload>(base.bmp8, ()=>({
                name: `Extended Status`,
                description: ``,
                
            })),
            IasCieAddressMac: makeType<ZigBee.ISecurityAndSafety.IArgIasCieAddressMac, ZigBee.ISecurityAndSafety.IArgIasCieAddressMacPayload>(base.uid, ()=>({
                name: `IAS CIE Address (MAC)`,
                description: ``,
                id: 0x0010,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            ManufacturerCode: makeType<ZigBee.ISecurityAndSafety.IArgManufacturerCode, ZigBee.ISecurityAndSafety.IArgManufacturerCodePayload>(base.u16, ()=>({
                name: `Manufacturer Code`,
                description: ``,
                
            })),
            NumberOfZoneSensitivityLevelsSupported: makeType<ZigBee.ISecurityAndSafety.IArgNumberOfZoneSensitivityLevelsSupported, ZigBee.ISecurityAndSafety.IArgNumberOfZoneSensitivityLevelsSupportedPayload>(base.u8, ()=>({
                name: `Number Of Zone Sensitivity Levels Supported`,
                description: ``,
                id: 0x0012,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            TestModeDuration: makeType<ZigBee.ISecurityAndSafety.IArgTestModeDuration, ZigBee.ISecurityAndSafety.IArgTestModeDurationPayload>(base.u8, ()=>({
                name: `Test mode duration`,
                description: ``,
                unit: units.Seconds,
                
            })),
            ZoneId: makeType<ZigBee.ISecurityAndSafety.IArgZoneId, ZigBee.ISecurityAndSafety.IArgZoneIdPayload>(base.u8, ()=>({
                name: `Zone ID`,
                description: ``,
                id: 0x0011,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            ZoneState: makeType<ZigBee.ISecurityAndSafety.IArgZoneState, ZigBee.ISecurityAndSafety.IArgZoneStatePayload>(base.enum8, ()=>({
                name: `Zone State`,
                description: ``,
                id: 0x0000,
                report: true,
                read: false,
                write: false,
                require: false,
                values: { 
                0x00: `Not Enrolled`, 
                0x01: `Enrolled`,  },
                
            })),
            ZoneStatus: makeType<ZigBee.ISecurityAndSafety.IArgZoneStatus, ZigBee.ISecurityAndSafety.IArgZoneStatusPayload>(base.bmp16, ()=>({
                name: `Zone Status`,
                description: ``,
                id: 0x0002,
                report: true,
                read: false,
                write: false,
                require: false,
                bits: { 
                0: `Alarm1`, 
                1: `Alarm2`, 
                2: `Tamper`, 
                3: `Battery`, 
                4: `Supervision notify`, 
                5: `Restore notify`, 
                6: `Trouble`, 
                7: `AC (mains)`, 
                8: `Test`, 
                9: `Battery defect`,  },
                
            })),
            ZoneType: makeType<ZigBee.ISecurityAndSafety.IArgZoneType, ZigBee.ISecurityAndSafety.IArgZoneTypePayload>(base.enum16, ()=>({
                name: `Zone Type`,
                description: ``,
                id: 0x0001,
                report: false,
                read: false,
                write: false,
                require: false,
                values: { 
                0x0000: `Standard CIE`, 
                0x000d: `Motion sensor`, 
                0x0015: `Contact Switch`, 
                0x0016: `Door/Window handle`, 
                0x0028: `Fire sensor`, 
                0x002a: `Water sensor`, 
                0x002b: `Carbon Monoxide sensor`, 
                0x002c: `Personal emergency device`, 
                0x002d: `Viburation/Movement sensor`, 
                0x010f: `Remote control`, 
                0x0115: `Key fob`, 
                0x021d: `Key pad`, 
                0x0225: `Standard warning device`, 
                0x0226: `Glass break sensor`, 
                0x0229: `Security repeater`, 
                0xffff: `Invalid zone type`,  },
                
            })), },
        IasZone: {
            ID: 0x0500,
            Name: `IAS Zone`,
            Desc: `provides attributes and commands for controlling IAS security
zone devices`,
            
            ZoneEnrollResponse: makeType<ZigBee.ISecurityAndSafety.IasZone.ICmdZoneEnrollResponse, ZigBee.ISecurityAndSafety.IasZone.ICmdZoneEnrollResponsePayload>(command, () => ({
                name: `Zone Enroll Response`,
                description: ``,
                id: 0x0000,
                payload: { 
                    EnrollResponseCode: ZigBee.SecurityAndSafety.Types.EnrollResponseCode,
                    ZoneId: ZigBee.SecurityAndSafety.Types.ZoneId,
                }
            })),

            InitiateNormalOperationMode: makeType<ZigBee.ISecurityAndSafety.IasZone.ICmdInitiateNormalOperationMode, ZigBee.ISecurityAndSafety.IasZone.ICmdInitiateNormalOperationModePayload>(command, () => ({
                name: `Initiate Normal Operation Mode`,
                description: ``,
                id: 0x0001,
                payload: {}
            })),

            InitiateTestMode: makeType<ZigBee.ISecurityAndSafety.IasZone.ICmdInitiateTestMode, ZigBee.ISecurityAndSafety.IasZone.ICmdInitiateTestModePayload>(command, () => ({
                name: `Initiate Test Mode`,
                description: ``,
                id: 0x0002,
                payload: { 
                    TestModeDuration: ZigBee.SecurityAndSafety.Types.TestModeDuration,
                    CurrentZoneSensitivityLevel: ZigBee.SecurityAndSafety.Types.CurrentZoneSensitivityLevel,
                }
            })),

            
            ZoneStateChangeNotification: makeType<ZigBee.ISecurityAndSafety.IasZone.ICmdZoneStateChangeNotification, ZigBee.ISecurityAndSafety.IasZone.ICmdZoneStateChangeNotificationPayload>(command, () => ({
                name: `Zone State Change Notification`,
                description: ``,
                id: 0x0000,
                payload: { 
                    ZoneStatus: ZigBee.SecurityAndSafety.Types.ZoneStatus,
                    ExtendedStatus: ZigBee.SecurityAndSafety.Types.ExtendedStatus,
                    ZoneId: ZigBee.SecurityAndSafety.Types.ZoneId,
                    Delay: ZigBee.SecurityAndSafety.Types.Delay,
                }
            })),

            ZoneEnrollRequest: makeType<ZigBee.ISecurityAndSafety.IasZone.ICmdZoneEnrollRequest, ZigBee.ISecurityAndSafety.IasZone.ICmdZoneEnrollRequestPayload>(command, () => ({
                name: `Zone Enroll Request`,
                description: ``,
                id: 0x0001,
                payload: { 
                    ZoneType: ZigBee.SecurityAndSafety.Types.ZoneType,
                    ManufacturerCode: ZigBee.SecurityAndSafety.Types.ManufacturerCode,
                }
            })),

            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        }
    };
    
    ZigBee.SecurityAndSafety.IasZone.Server.Attribute = { 
        0x0000: ZigBee.SecurityAndSafety.Types.ZoneState,
        0x0001: ZigBee.SecurityAndSafety.Types.ZoneType,
        0x0002: ZigBee.SecurityAndSafety.Types.ZoneStatus,
        0x0010: ZigBee.SecurityAndSafety.Types.IasCieAddressMac,
        0x0011: ZigBee.SecurityAndSafety.Types.ZoneId,
        0x0012: ZigBee.SecurityAndSafety.Types.NumberOfZoneSensitivityLevelsSupported,
        0x0013: ZigBee.SecurityAndSafety.Types.CurrentZoneSensitivityLevel,
    };
    ZigBee.SecurityAndSafety.IasZone.Client.Attribute = { 
    };
    ZigBee.SecurityAndSafety.IasZone.Server.Command = { 
        0x00: ZigBee.SecurityAndSafety.IasZone.ZoneEnrollResponse,
        0x01: ZigBee.SecurityAndSafety.IasZone.InitiateNormalOperationMode,
        0x02: ZigBee.SecurityAndSafety.IasZone.InitiateTestMode,
    };
    ZigBee.SecurityAndSafety.IasZone.Client.Command = { 
        0x00: ZigBee.SecurityAndSafety.IasZone.ZoneStateChangeNotification,
        0x01: ZigBee.SecurityAndSafety.IasZone.ZoneEnrollRequest,
    };
    export const Ikea = {
        Types: { 
            VocIndex: makeType<ZigBee.IIkea.IArgVocIndex, ZigBee.IIkea.IArgVocIndexPayload>(base.float, ()=>({
                name: `VOC Index`,
                description: `represents the Sensirion AG measurement of VOC status relative to the sensors recent history.
The VOC Index uses a moving average over the past 24 hours (called the "learning time") as offset.
The VOC Index mimics the human nose’s perception of odors with a relative intensity compared to recent history.
A VOC Index above 100 means that there are more VOCs compared to the average (e.g., induced by a VOC event from
cooking, cleaning, breathing, etc.) while a VOC Index below 100 means that there are fewer VOCs compared to the
average (e.g., induced by fresh air from an open window, using an air purifier, etc.).
The VOC Index ranges from 1 to 500.`,
                id: 0x0000,
                report: true,
                read: true,
                write: false,
                require: false,
                
            })), },
        IkeaAirQuality: {
            ID: 0xFC7E,
            Name: `IKEA Air Quality`,
            Desc: ``,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        }
    };
    
    ZigBee.Ikea.IkeaAirQuality.Server.Attribute = { 
        0x0000: ZigBee.Ikea.Types.VocIndex,
    };
    ZigBee.Ikea.IkeaAirQuality.Client.Attribute = { 
    };
    ZigBee.Ikea.IkeaAirQuality.Server.Command = { 
    };
    ZigBee.Ikea.IkeaAirQuality.Client.Command = { 
    };
    export const Lighting = {
        Types: { 
            Action: makeType<ZigBee.ILighting.IArgAction, ZigBee.ILighting.IArgActionPayload>(base.enum8, ()=>({
                name: `Action`,
                description: ``,
                values: { 
                0x00: `De-activate color loop`, 
                0x01: `Activate from ColorLoopStartEnhancedHue`, 
                0x02: `Activate from EnhancedCurrentHue`,  },
                
            })),
            BallastFactorAdjustment: makeType<ZigBee.ILighting.IArgBallastFactorAdjustment, ZigBee.ILighting.IArgBallastFactorAdjustmentPayload>(base.u8, ()=>({
                name: `Ballast Factor Adjustment`,
                description: `specifies the multiplication factor, as a percentage, to be applied
to the configured light output of the lamps. A typical usage of this
mechanism is to compensate for reduction in efficiency over the lifetime
of a lamp. The light output is given by
Actual light output = configured light output x BallastFactorAdjustment / 100%
The range for this attribute is manufacturer dependent`,
                id: 0x0015,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Percent,
                
            })),
            BallastStatus: makeType<ZigBee.ILighting.IArgBallastStatus, ZigBee.ILighting.IArgBallastStatusPayload>(base.bmp8, ()=>({
                name: `Ballast Status`,
                description: `It specifies the activity status of the ballast functions. Where a
function is active, the corresponding bit is 1. Where a function is
not active, the corresponding bit is 0. This means that bit 0 with
a value of 0 means the ballast is operational and bit 1 with a
value of 0 that the lamp(s) is/are in their socket(s)`,
                id: 0x0002,
                report: false,
                read: true,
                write: false,
                require: false,
                bits: { 
                0: `Non-operational`, 
                1: `Lamp not in socket`,  },
                
            })),
            ColorCapabilities: makeType<ZigBee.ILighting.IArgColorCapabilities, ZigBee.ILighting.IArgColorCapabilitiesPayload>(base.bmp16, ()=>({
                name: `Color capabilities`,
                description: `specifies the color capabilities of the device supporting the color
control cluster`,
                id: 0x400A,
                report: false,
                read: true,
                write: false,
                require: false,
                bits: { 
                0: `Hue saturation`, 
                1: `Enhanced Hue saturation`, 
                2: `Color loop`, 
                3: `CIE 1931 XY`, 
                4: `Color temperature`,  },
                
            })),
            ColorControlOptions: makeType<ZigBee.ILighting.IArgColorControlOptions, ZigBee.ILighting.IArgColorControlOptionsPayload>(base.bmp8, ()=>({
                name: `Color control options`,
                description: `is a bitmap that determines the default behavior of some cluster commands`,
                id: 0x000F,
                report: false,
                read: true,
                write: true,
                require: false,
                bits: { 
                0x00: `Execute if off`,  },
                
            })),
            ColorLoopActive: makeType<ZigBee.ILighting.IArgColorLoopActive, ZigBee.ILighting.IArgColorLoopActivePayload>(base.u8, ()=>({
                name: `Color loop active`,
                description: `specifies the current active status of the color loop. 0x00 means
inactive, 0x01 means active`,
                id: 0x4002,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            ColorLoopDirection: makeType<ZigBee.ILighting.IArgColorLoopDirection, ZigBee.ILighting.IArgColorLoopDirectionPayload>(base.u8, ()=>({
                name: `Color loop direction`,
                description: `specifies the current direction of the color loop. If this attribute
has the value 0x00, the EnhancedCurrentHue is be decremented. If this
attribute has the value 0x01, the EnhancedCurrentHue is incremented`,
                id: 0x4003,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            ColorLoopStartEnhancedHue: makeType<ZigBee.ILighting.IArgColorLoopStartEnhancedHue, ZigBee.ILighting.IArgColorLoopStartEnhancedHuePayload>(base.u16, ()=>({
                name: `Color loop start enhanced hue`,
                description: `specifies the value of the EnhancedCurrentHue attribute from which
the color loop starts`,
                id: 0x4005,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            ColorLoopStoredEnhancedHue: makeType<ZigBee.ILighting.IArgColorLoopStoredEnhancedHue, ZigBee.ILighting.IArgColorLoopStoredEnhancedHuePayload>(base.u16, ()=>({
                name: `Color loop stored enhanced hue`,
                description: `specifies the value of the EnhancedCurrentHue attribute before the
color loop was started. Once the color loop is complete, It is restored
to this value`,
                id: 0x4006,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            ColorLoopTime: makeType<ZigBee.ILighting.IArgColorLoopTime, ZigBee.ILighting.IArgColorLoopTimePayload>(base.u16, ()=>({
                name: `Color loop time`,
                description: `specifies the number of seconds it takes to perform a full color
loop, i.e., to cycle all values of EnhancedCurrentHue`,
                id: 0x4004,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            ColorMode: makeType<ZigBee.ILighting.IArgColorMode, ZigBee.ILighting.IArgColorModePayload>(base.enum8, ()=>({
                name: `Color Mode`,
                description: `indicates which attributes are currently determining the color of
the device. This attribute is optional if the device does not implement
CurrentHue and CurrentSaturation`,
                id: 0x0008,
                report: false,
                read: true,
                write: false,
                require: false,
                values: { 
                0x00: `Current hue and current saturation`, 
                0x01: `Current X and Current Y`, 
                0x02: `Color temperature`,  },
                
            })),
            ColorPointBlueIntensity: makeType<ZigBee.ILighting.IArgColorPointBlueIntensity, ZigBee.ILighting.IArgColorPointBlueIntensityPayload>(base.u8, ()=>({
                name: `Color Point Blue intensity`,
                description: `contains a representation of the maximum intensity of this attribute as
defined in the Dimming Light Curve in the Ballast Configuration cluster,
normalized such that the attribute with the highest maximum intensity
contains the value 0xfe. A value of 0xff indicates that this attribute is
not available`,
                id: 0x003C,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            ColorPointBlueX: makeType<ZigBee.ILighting.IArgColorPointBlueX, ZigBee.ILighting.IArgColorPointBlueXPayload>(base.u16, ()=>({
                name: `Color Point Blue X`,
                description: `contains the normalized chromaticity value x for this attribute, as
defined in the CIE xyY Color Space. x = value / 65536 (value
in the range 0 to 65279 inclusive)`,
                id: 0x003A,
                report: false,
                read: true,
                write: true,
                require: false,
                scale: 65536,
                
            })),
            ColorPointBlueY: makeType<ZigBee.ILighting.IArgColorPointBlueY, ZigBee.ILighting.IArgColorPointBlueYPayload>(base.u16, ()=>({
                name: `Color Point Blue Y`,
                description: `contains the normalized chromaticity value y for this attribute, as
defined in the CIE xyY Color Space. y = value / 65536 (value
in the range 0 to 65279 inclusive)`,
                id: 0x003B,
                report: false,
                read: true,
                write: true,
                require: false,
                scale: 65536,
                
            })),
            ColorPointGreenIntensity: makeType<ZigBee.ILighting.IArgColorPointGreenIntensity, ZigBee.ILighting.IArgColorPointGreenIntensityPayload>(base.u8, ()=>({
                name: `Color Point Green intensity`,
                description: `contains a representation of the maximum intensity of this attribute as
defined in the Dimming Light Curve in the Ballast Configuration cluster,
normalized such that the attribute with the highest maximum intensity
contains the value 0xfe. A value of 0xff indicates that this attribute is
not available`,
                id: 0x0038,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            ColorPointGreenX: makeType<ZigBee.ILighting.IArgColorPointGreenX, ZigBee.ILighting.IArgColorPointGreenXPayload>(base.u16, ()=>({
                name: `Color Point Green X`,
                description: `contains the normalized chromaticity value x for this attribute, as
defined in the CIE xyY Color Space. x = value / 65536 (value
in the range 0 to 65279 inclusive)`,
                id: 0x0036,
                report: false,
                read: true,
                write: true,
                require: false,
                scale: 65536,
                
            })),
            ColorPointGreenY: makeType<ZigBee.ILighting.IArgColorPointGreenY, ZigBee.ILighting.IArgColorPointGreenYPayload>(base.u16, ()=>({
                name: `Color Point Green Y`,
                description: `contains the normalized chromaticity value y for this attribute, as
defined in the CIE xyY Color Space. y = value / 65536 (value
in the range 0 to 65279 inclusive)`,
                id: 0x0037,
                report: false,
                read: true,
                write: true,
                require: false,
                scale: 65536,
                
            })),
            ColorPointRedIntensity: makeType<ZigBee.ILighting.IArgColorPointRedIntensity, ZigBee.ILighting.IArgColorPointRedIntensityPayload>(base.u8, ()=>({
                name: `Color Point Red intensity`,
                description: `contains a representation of the maximum intensity of this attribute as
defined in the Dimming Light Curve in the Ballast Configuration cluster,
normalized such that the attribute with the highest maximum intensity
contains the value 0xfe. A value of 0xff indicates that this attribute is
not available`,
                id: 0x0034,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            ColorPointRedX: makeType<ZigBee.ILighting.IArgColorPointRedX, ZigBee.ILighting.IArgColorPointRedXPayload>(base.u16, ()=>({
                name: `Color Point Red X`,
                description: `contains the normalized chromaticity value x for this attribute, as
defined in the CIE xyY Color Space. x = value / 65536 (value
in the range 0 to 65279 inclusive)`,
                id: 0x0032,
                report: false,
                read: true,
                write: true,
                require: false,
                scale: 65536,
                
            })),
            ColorPointRedY: makeType<ZigBee.ILighting.IArgColorPointRedY, ZigBee.ILighting.IArgColorPointRedYPayload>(base.u16, ()=>({
                name: `Color Point Red Y`,
                description: `contains the normalized chromaticity value y for this attribute, as
defined in the CIE xyY Color Space. y = value / 65536 (value
in the range 0 to 65279 inclusive)`,
                id: 0x0033,
                report: false,
                read: true,
                write: true,
                require: false,
                scale: 65536,
                
            })),
            ColorTemperature: makeType<ZigBee.ILighting.IArgColorTemperature, ZigBee.ILighting.IArgColorTemperaturePayload>(base.u16, ()=>({
                name: `Color Temperature`,
                description: ``,
                
            })),
            ColorTemperatureMax: makeType<ZigBee.ILighting.IArgColorTemperatureMax, ZigBee.ILighting.IArgColorTemperatureMaxPayload>(base.u16, ()=>({
                name: `Color Temperature max`,
                description: ``,
                
            })),
            ColorTemperatureMaxMireds: makeType<ZigBee.ILighting.IArgColorTemperatureMaxMireds, ZigBee.ILighting.IArgColorTemperatureMaxMiredsPayload>(base.u16, ()=>({
                name: `Color Temperature max Mireds`,
                description: ``,
                unit: units.Mired,
                
            })),
            ColorTemperatureMin: makeType<ZigBee.ILighting.IArgColorTemperatureMin, ZigBee.ILighting.IArgColorTemperatureMinPayload>(base.u16, ()=>({
                name: `Color Temperature min`,
                description: ``,
                
            })),
            ColorTemperatureMinMireds: makeType<ZigBee.ILighting.IArgColorTemperatureMinMireds, ZigBee.ILighting.IArgColorTemperatureMinMiredsPayload>(base.u16, ()=>({
                name: `Color Temperature min Mireds`,
                description: ``,
                unit: units.Mired,
                
            })),
            ColorTemperatureMireds: makeType<ZigBee.ILighting.IArgColorTemperatureMireds, ZigBee.ILighting.IArgColorTemperatureMiredsPayload>(base.u16, ()=>({
                name: `Color temperature Mireds`,
                description: `contains a scaled inverse of the current value of the color
temperature. The unit of ColorTemperatureMireds is the mired
(micro reciprocal degree), a.k.a mirek (micro reciprocal
kelvin). Color temperature in kelvins = 1,000,000 / ColorTemperatureMireds,
where ColorTemperatureMireds is in the range 1 to 65279 mireds inclusive,
giving a color temperature range from 1,000,000 kelvins to 15.32 kelvins`,
                id: 0x0007,
                report: true,
                read: true,
                write: false,
                require: false,
                unit: units.Mired,
                
            })),
            ColorTemperaturePhysicalMaxMireds: makeType<ZigBee.ILighting.IArgColorTemperaturePhysicalMaxMireds, ZigBee.ILighting.IArgColorTemperaturePhysicalMaxMiredsPayload>(base.u16, ()=>({
                name: `Color temperature physical max Mireds`,
                description: `indicates the maximum mired value supported by the hardware.
ColorTempPhysicalMaxMireds corresponds to the minimum color
temperature in Kelvins supported by the hardware.
ColorTemperatureMireds ≤ ColorTempPhysicalMaxMireds`,
                id: 0x400C,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.Mired,
                
            })),
            ColorTemperaturePhysicalMinMireds: makeType<ZigBee.ILighting.IArgColorTemperaturePhysicalMinMireds, ZigBee.ILighting.IArgColorTemperaturePhysicalMinMiredsPayload>(base.u16, ()=>({
                name: `Color temperature physical min Mireds`,
                description: `indicates the minimum mired value supported by the hardware.
ColorTempPhysicalMinMireds corresponds to the maximum color
temperature in Kelvins supported by the hardware.
ColorTempPhysicalMinMireds ≤ ColorTemperatureMireds`,
                id: 0x400B,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.Mired,
                
            })),
            ColorX: makeType<ZigBee.ILighting.IArgColorX, ZigBee.ILighting.IArgColorXPayload>(base.u16, ()=>({
                name: `Color X`,
                description: `contains the normalized chromaticity value x for this attribute, as
defined in the CIE xyY Color Space. x = value / 65536 (value
in the range 0 to 65279 inclusive)`,
                scale: 65536,
                
            })),
            ColorY: makeType<ZigBee.ILighting.IArgColorY, ZigBee.ILighting.IArgColorYPayload>(base.u16, ()=>({
                name: `Color Y`,
                description: `contains the normalized chromaticity value y for this attribute, as
defined in the CIE xyY Color Space. y = value / 65536 (value
in the range 0 to 65279 inclusive)`,
                scale: 65536,
                
            })),
            CompensationText: makeType<ZigBee.ILighting.IArgCompensationText, ZigBee.ILighting.IArgCompensationTextPayload>(base.cstring, ()=>({
                name: `Compensation Text`,
                description: `holds a textual indication of what mechanism, if any, is in use to
compensate for color/intensity drift over time`,
                id: 0x0006,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            CoupleColorTempToLevelMinMireds: makeType<ZigBee.ILighting.IArgCoupleColorTempToLevelMinMireds, ZigBee.ILighting.IArgCoupleColorTempToLevelMinMiredsPayload>(base.u16, ()=>({
                name: `Couple Color Temp to Level Min Mireds`,
                description: `specifies a lower bound on the value of the ColorTemperatureMireds attribute for the purposes of coupling the
ColorTemperatureMireds attribute to the CurrentLevel attribute when the CoupleColorTempToLevel bit of the
Options attribute of the Level Control cluster is equal to 1`,
                id: 0x400D,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.Mired,
                
            })),
            CurrentHue: makeType<ZigBee.ILighting.IArgCurrentHue, ZigBee.ILighting.IArgCurrentHuePayload>(base.u8, ()=>({
                name: `Current hue`,
                description: `contains the current hue value of the light. Hue = CurrentHue x 360 / 254
(CurrentHue in the range 0 - 254 inclusive)`,
                id: 0x0000,
                report: true,
                read: true,
                write: false,
                require: false,
                unit: units.DegreesAngular,
                scale: 0.70556,
                
            })),
            CurrentSaturation: makeType<ZigBee.ILighting.IArgCurrentSaturation, ZigBee.ILighting.IArgCurrentSaturationPayload>(base.u8, ()=>({
                name: `Current saturation`,
                description: `holds the current saturation value of the light.
Saturation = CurrentSaturation/254 (CurrentSaturation in the range
0 - 254 inclusive)`,
                id: 0x0001,
                report: true,
                read: true,
                write: false,
                require: false,
                scale: 254,
                
            })),
            CurrentX: makeType<ZigBee.ILighting.IArgCurrentX, ZigBee.ILighting.IArgCurrentXPayload>(base.u16, ()=>({
                name: `Current X`,
                description: `contains the normalized chromaticity value x for this attribute, as
defined in the CIE xyY Color Space. x = value / 65536 (value
in the range 0 to 65279 inclusive)`,
                id: 0x0003,
                report: true,
                read: true,
                write: false,
                require: false,
                scale: 65536,
                
            })),
            CurrentY: makeType<ZigBee.ILighting.IArgCurrentY, ZigBee.ILighting.IArgCurrentYPayload>(base.u16, ()=>({
                name: `Current Y`,
                description: `contains the normalized chromaticity value y for this attribute, as
defined in the CIE xyY Color Space. y = value / 65536 (value
in the range 0 to 65279 inclusive)`,
                id: 0x0004,
                report: true,
                read: true,
                write: false,
                require: false,
                scale: 65536,
                
            })),
            DriftCompensation: makeType<ZigBee.ILighting.IArgDriftCompensation, ZigBee.ILighting.IArgDriftCompensationPayload>(base.enum8, ()=>({
                name: `Drift Compensation`,
                description: `indicates what mechanism, if any, is in use for compensation for
color/intensity drift over time`,
                id: 0x0005,
                report: false,
                read: true,
                write: false,
                require: false,
                values: { 
                0x00: `None`, 
                0x01: `Other / Unknown`, 
                0x02: `Temperature monitoring`, 
                0x03: `Optical luminance monitoring and feedback`, 
                0x04: `Optical color monitoring and feedback`,  },
                
            })),
            EnhancedColorMode: makeType<ZigBee.ILighting.IArgEnhancedColorMode, ZigBee.ILighting.IArgEnhancedColorModePayload>(base.enum8, ()=>({
                name: `Enhanced color mode`,
                description: `specifies which attributes are currently determining the color of
the device`,
                id: 0x4001,
                report: false,
                read: true,
                write: false,
                require: false,
                values: { 
                0x00: `Current hue and current saturation`, 
                0x01: `Current X and Current Y`, 
                0x02: `Color temperature`, 
                0x03: `Enhanced current hue and current saturation`,  },
                
            })),
            EnhancedCurrentHue: makeType<ZigBee.ILighting.IArgEnhancedCurrentHue, ZigBee.ILighting.IArgEnhancedCurrentHuePayload>(base.u16, ()=>({
                name: `Enhanced current hue`,
                description: `represents non-equidistant steps along the CIE 1931 color triangle,
and it provides 16-bits precision. The upper 8 bits of this attribute
are used as an index in the implementation specific XY lookup table to
provide the non-equidistance steps. The lower 8 bits are used to
interpolate between these steps in a linear way in order to provide color
zoom for the user. To provide compatibility with standard ZCL, the
CurrentHue attribute contains a hue value in the range 0 to 254,
calculated from the EnhancedCurrentHue attribute`,
                id: 0x4000,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            EnhancedHue: makeType<ZigBee.ILighting.IArgEnhancedHue, ZigBee.ILighting.IArgEnhancedHuePayload>(base.u16, ()=>({
                name: `Enhanced Hue`,
                description: ``,
                unit: units.DegreesAngular,
                scale: 0.002756094,
                
            })),
            Hue: makeType<ZigBee.ILighting.IArgHue, ZigBee.ILighting.IArgHuePayload>(base.u8, ()=>({
                name: `Hue`,
                description: ``,
                unit: units.DegreesAngular,
                scale: 0.70556,
                
            })),
            HueDirection: makeType<ZigBee.ILighting.IArgHueDirection, ZigBee.ILighting.IArgHueDirectionPayload>(base.enum8, ()=>({
                name: `Hue Direction`,
                description: ``,
                values: { 
                0x00: `Decrement hue`, 
                0x01: `Increment hue`,  },
                
            })),
            IntrinsicBallastFactor: makeType<ZigBee.ILighting.IArgIntrinsicBallastFactor, ZigBee.ILighting.IArgIntrinsicBallastFactorPayload>(base.u8, ()=>({
                name: `Intrinsic Ballast Factor`,
                description: `specifies, as a percentage, the ballast factor of the ballast/lamp
combination, prior to any adjustment.`,
                id: 0x0014,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Percent,
                
            })),
            LampAlarmMode: makeType<ZigBee.ILighting.IArgLampAlarmMode, ZigBee.ILighting.IArgLampAlarmModePayload>(base.bmp8, ()=>({
                name: `Lamp Alarm Mode`,
                description: `specifies which attributes can cause an alarm notification to be
generated`,
                id: 0x0034,
                report: false,
                read: true,
                write: true,
                require: false,
                bits: { 
                0: `Lamp Burn Hours`,  },
                
            })),
            LampBurnHours: makeType<ZigBee.ILighting.IArgLampBurnHours, ZigBee.ILighting.IArgLampBurnHoursPayload>(base.u24, ()=>({
                name: `Lamp Burn Hours`,
                description: `specifies the length of time, in hours, the currently connected
lamps have been operated, cumulative since the last re-lamping. Burn
hours are not accumulated if the lamps are off and are reset to 0
when a lamp is changed`,
                id: 0x0033,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Hours,
                
            })),
            LampBurnHoursTripPoint: makeType<ZigBee.ILighting.IArgLampBurnHoursTripPoint, ZigBee.ILighting.IArgLampBurnHoursTripPointPayload>(base.u24, ()=>({
                name: `Lamp Burn Hours Trip Point`,
                description: `specifies the number of hours the LampBurnHours attribute can
reach before an alarm is generated`,
                id: 0x0035,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Hours,
                
            })),
            LampManufacturer: makeType<ZigBee.ILighting.IArgLampManufacturer, ZigBee.ILighting.IArgLampManufacturerPayload>(base.cstring, ()=>({
                name: `Lamp Manufacturer`,
                description: `specifies the name of the manufacturer of the currently connected
lamps.`,
                id: 0x0031,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            LampQuantity: makeType<ZigBee.ILighting.IArgLampQuantity, ZigBee.ILighting.IArgLampQuantityPayload>(base.u8, ()=>({
                name: `Lamp Quantity`,
                description: `specifies the number of lamps connected to this ballast`,
                id: 0x0020,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            LampRatedHours: makeType<ZigBee.ILighting.IArgLampRatedHours, ZigBee.ILighting.IArgLampRatedHoursPayload>(base.u24, ()=>({
                name: `Lamp Rated Hours`,
                description: `specifies the number of hours of use the lamps are rated for by
the manufacturer`,
                id: 0x0032,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Hours,
                
            })),
            LampType: makeType<ZigBee.ILighting.IArgLampType, ZigBee.ILighting.IArgLampTypePayload>(base.cstring, ()=>({
                name: `Lamp Type`,
                description: ``,
                id: 0x0030,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            MaxLevel: makeType<ZigBee.ILighting.IArgMaxLevel, ZigBee.ILighting.IArgMaxLevelPayload>(base.u8, ()=>({
                name: `Max Level`,
                description: `specifies the maximum light level the ballast is permitted to use.
It is specified in the range 0x01 to 0xfe, and specifies the light
output of the ballast according to the dimming light curve. The value
is both less than or equal to PhysicalMaxLevel and greater than or equal
to MinLevel`,
                id: 0x0011,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            MinLevel: makeType<ZigBee.ILighting.IArgMinLevel, ZigBee.ILighting.IArgMinLevelPayload>(base.u8, ()=>({
                name: `Min Level`,
                description: ``,
                id: 0x0010,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            MoveDirection: makeType<ZigBee.ILighting.IArgMoveDirection, ZigBee.ILighting.IArgMoveDirectionPayload>(base.enum8, ()=>({
                name: `Move Direction`,
                description: ``,
                values: { 
                0x00: `Shortest distance`, 
                0x01: `Longest Distance`, 
                0x02: `Up`, 
                0x03: `Down`,  },
                
            })),
            MoveMode: makeType<ZigBee.ILighting.IArgMoveMode, ZigBee.ILighting.IArgMoveModePayload>(base.enum8, ()=>({
                name: `Move mode`,
                description: ``,
                values: { 
                0x00: `Stop`, 
                0x01: `Up`, 
                0x03: `Down`,  },
                
            })),
            NumberOfPrimaries: makeType<ZigBee.ILighting.IArgNumberOfPrimaries, ZigBee.ILighting.IArgNumberOfPrimariesPayload>(base.u8, ()=>({
                name: `Number of primaries`,
                description: `contains the number of color primaries implemented on this device.
A value of 0xff indicates that the number of primaries is unknown`,
                id: 0x0010,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            PhysicalMaxLevel: makeType<ZigBee.ILighting.IArgPhysicalMaxLevel, ZigBee.ILighting.IArgPhysicalMaxLevelPayload>(base.u8, ()=>({
                name: `Physical Max Level`,
                description: `specifies the maximum light level the ballast can achieve. It is
specified in the range 0x01 to 0xfe, and specifies the light output
of the ballast according to the dimming light curve`,
                id: 0x0001,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            PhysicalMinLevel: makeType<ZigBee.ILighting.IArgPhysicalMinLevel, ZigBee.ILighting.IArgPhysicalMinLevelPayload>(base.u8, ()=>({
                name: `Physical Min Level`,
                description: `specifies the minimum light level the ballast can achieve. It is
specified in the range 0x01 to 0xfe, and specifies the light output
of the ballast according to the dimming light curve`,
                id: 0x0000,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            PowerOnColorTemperature: makeType<ZigBee.ILighting.IArgPowerOnColorTemperature, ZigBee.ILighting.IArgPowerOnColorTemperaturePayload>(base.u16, ()=>({
                name: `Power-on color temperature`,
                description: ``,
                id: 0x4010,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            PowerOnLevel: makeType<ZigBee.ILighting.IArgPowerOnLevel, ZigBee.ILighting.IArgPowerOnLevelPayload>(base.u8, ()=>({
                name: `Power On level`,
                description: `It specifies the light level to which the ballast will go when power
is applied (e.g., when mains power is re-established after a power
failure). It can have a value between 0x00 and 0xfe to set it to
a specific light level, according to the dimming light curve, or
0xff to set it to the value it was before the power failure`,
                id: 0x0012,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            PowerOnTime: makeType<ZigBee.ILighting.IArgPowerOnTime, ZigBee.ILighting.IArgPowerOnTimePayload>(base.u16, ()=>({
                name: `Power-On Time`,
                description: `The transition time in 1/10ths of a second.`,
                id: 0x0013,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Seconds,
                scale: 10,
                
            })),
            Primary1Intensity: makeType<ZigBee.ILighting.IArgPrimary1Intensity, ZigBee.ILighting.IArgPrimary1IntensityPayload>(base.u8, ()=>({
                name: `Primary1 intensity`,
                description: `contains a representation of the maximum intensity of this attribute as
defined in the Dimming Light Curve in the Ballast Configuration cluster,
normalized such that the attribute with the highest maximum intensity
contains the value 0xfe. A value of 0xff indicates that this attribute is
not available`,
                id: 0x0013,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            Primary1X: makeType<ZigBee.ILighting.IArgPrimary1X, ZigBee.ILighting.IArgPrimary1XPayload>(base.u16, ()=>({
                name: `Primary1 X`,
                description: `contains the normalized chromaticity value x for this attribute, as
defined in the CIE xyY Color Space. x = value / 65536 (value
in the range 0 to 65279 inclusive)`,
                id: 0x0011,
                report: false,
                read: true,
                write: false,
                require: false,
                scale: 65536,
                
            })),
            Primary1Y: makeType<ZigBee.ILighting.IArgPrimary1Y, ZigBee.ILighting.IArgPrimary1YPayload>(base.u16, ()=>({
                name: `Primary1 Y`,
                description: `contains the normalized chromaticity value y for this attribute, as
defined in the CIE xyY Color Space. y = value / 65536 (value
in the range 0 to 65279 inclusive)`,
                id: 0x0012,
                report: false,
                read: true,
                write: false,
                require: false,
                scale: 65536,
                
            })),
            Primary2Intensity: makeType<ZigBee.ILighting.IArgPrimary2Intensity, ZigBee.ILighting.IArgPrimary2IntensityPayload>(base.u8, ()=>({
                name: `Primary2 intensity`,
                description: `contains a representation of the maximum intensity of this attribute as
defined in the Dimming Light Curve in the Ballast Configuration cluster,
normalized such that the attribute with the highest maximum intensity
contains the value 0xfe. A value of 0xff indicates that this attribute is
not available`,
                id: 0x0017,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            Primary2X: makeType<ZigBee.ILighting.IArgPrimary2X, ZigBee.ILighting.IArgPrimary2XPayload>(base.u16, ()=>({
                name: `Primary2 X`,
                description: `contains the normalized chromaticity value x for this attribute, as
defined in the CIE xyY Color Space. x = value / 65536 (value
in the range 0 to 65279 inclusive)`,
                id: 0x0015,
                report: false,
                read: true,
                write: false,
                require: false,
                scale: 65536,
                
            })),
            Primary2Y: makeType<ZigBee.ILighting.IArgPrimary2Y, ZigBee.ILighting.IArgPrimary2YPayload>(base.u16, ()=>({
                name: `Primary2 Y`,
                description: `contains the normalized chromaticity value y for this attribute, as
defined in the CIE xyY Color Space. y = value / 65536 (value
in the range 0 to 65279 inclusive)`,
                id: 0x0016,
                report: false,
                read: true,
                write: false,
                require: false,
                scale: 65536,
                
            })),
            Primary3Intensity: makeType<ZigBee.ILighting.IArgPrimary3Intensity, ZigBee.ILighting.IArgPrimary3IntensityPayload>(base.u8, ()=>({
                name: `Primary3 intensity`,
                description: `contains a representation of the maximum intensity of this attribute as
defined in the Dimming Light Curve in the Ballast Configuration cluster,
normalized such that the attribute with the highest maximum intensity
contains the value 0xfe. A value of 0xff indicates that this attribute is
not available`,
                id: 0x001B,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            Primary3X: makeType<ZigBee.ILighting.IArgPrimary3X, ZigBee.ILighting.IArgPrimary3XPayload>(base.u16, ()=>({
                name: `Primary3 X`,
                description: `contains the normalized chromaticity value x for this attribute, as
defined in the CIE xyY Color Space. x = value / 65536 (value
in the range 0 to 65279 inclusive)`,
                id: 0x0019,
                report: false,
                read: true,
                write: false,
                require: false,
                scale: 65536,
                
            })),
            Primary3Y: makeType<ZigBee.ILighting.IArgPrimary3Y, ZigBee.ILighting.IArgPrimary3YPayload>(base.u16, ()=>({
                name: `Primary3 Y`,
                description: `contains the normalized chromaticity value y for this attribute, as
defined in the CIE xyY Color Space. y = value / 65536 (value
in the range 0 to 65279 inclusive)`,
                id: 0x001A,
                report: false,
                read: true,
                write: false,
                require: false,
                scale: 65536,
                
            })),
            Primary4Intensity: makeType<ZigBee.ILighting.IArgPrimary4Intensity, ZigBee.ILighting.IArgPrimary4IntensityPayload>(base.u8, ()=>({
                name: `Primary4 intensity`,
                description: `contains a representation of the maximum intensity of this attribute as
defined in the Dimming Light Curve in the Ballast Configuration cluster,
normalized such that the attribute with the highest maximum intensity
contains the value 0xfe. A value of 0xff indicates that this attribute is
not available`,
                id: 0x0022,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            Primary4X: makeType<ZigBee.ILighting.IArgPrimary4X, ZigBee.ILighting.IArgPrimary4XPayload>(base.u16, ()=>({
                name: `Primary4 X`,
                description: `contains the normalized chromaticity value x for this attribute, as
defined in the CIE xyY Color Space. x = value / 65536 (value
in the range 0 to 65279 inclusive)`,
                id: 0x0020,
                report: false,
                read: true,
                write: false,
                require: false,
                scale: 65536,
                
            })),
            Primary4Y: makeType<ZigBee.ILighting.IArgPrimary4Y, ZigBee.ILighting.IArgPrimary4YPayload>(base.u16, ()=>({
                name: `Primary4 Y`,
                description: `contains the normalized chromaticity value y for this attribute, as
defined in the CIE xyY Color Space. y = value / 65536 (value
in the range 0 to 65279 inclusive)`,
                id: 0x0021,
                report: false,
                read: true,
                write: false,
                require: false,
                scale: 65536,
                
            })),
            Primary5Intensity: makeType<ZigBee.ILighting.IArgPrimary5Intensity, ZigBee.ILighting.IArgPrimary5IntensityPayload>(base.u8, ()=>({
                name: `Primary5 intensity`,
                description: `contains a representation of the maximum intensity of this attribute as
defined in the Dimming Light Curve in the Ballast Configuration cluster,
normalized such that the attribute with the highest maximum intensity
contains the value 0xfe. A value of 0xff indicates that this attribute is
not available`,
                id: 0x0026,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            Primary5X: makeType<ZigBee.ILighting.IArgPrimary5X, ZigBee.ILighting.IArgPrimary5XPayload>(base.u16, ()=>({
                name: `Primary5 X`,
                description: `contains the normalized chromaticity value x for this attribute, as
defined in the CIE xyY Color Space. x = value / 65536 (value
in the range 0 to 65279 inclusive)`,
                id: 0x0024,
                report: false,
                read: true,
                write: false,
                require: false,
                scale: 65536,
                
            })),
            Primary5Y: makeType<ZigBee.ILighting.IArgPrimary5Y, ZigBee.ILighting.IArgPrimary5YPayload>(base.u16, ()=>({
                name: `Primary5 Y`,
                description: `contains the normalized chromaticity value y for this attribute, as
defined in the CIE xyY Color Space. y = value / 65536 (value
in the range 0 to 65279 inclusive)`,
                id: 0x0025,
                report: false,
                read: true,
                write: false,
                require: false,
                scale: 65536,
                
            })),
            Primary6Intensity: makeType<ZigBee.ILighting.IArgPrimary6Intensity, ZigBee.ILighting.IArgPrimary6IntensityPayload>(base.u8, ()=>({
                name: `Primary6 intensity`,
                description: `contains a representation of the maximum intensity of this attribute as
defined in the Dimming Light Curve in the Ballast Configuration cluster,
normalized such that the attribute with the highest maximum intensity
contains the value 0xfe. A value of 0xff indicates that this attribute is
not available`,
                id: 0x002A,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            Primary6X: makeType<ZigBee.ILighting.IArgPrimary6X, ZigBee.ILighting.IArgPrimary6XPayload>(base.u16, ()=>({
                name: `Primary6 X`,
                description: `contains the normalized chromaticity value x for this attribute, as
defined in the CIE xyY Color Space. x = value / 65536 (value
in the range 0 to 65279 inclusive)`,
                id: 0x0028,
                report: false,
                read: true,
                write: false,
                require: false,
                scale: 65536,
                
            })),
            Primary6Y: makeType<ZigBee.ILighting.IArgPrimary6Y, ZigBee.ILighting.IArgPrimary6YPayload>(base.u16, ()=>({
                name: `Primary6 Y`,
                description: `contains the normalized chromaticity value y for this attribute, as
defined in the CIE xyY Color Space. y = value / 65536 (value
in the range 0 to 65279 inclusive)`,
                id: 0x0029,
                report: false,
                read: true,
                write: false,
                require: false,
                scale: 65536,
                
            })),
            Rate: makeType<ZigBee.ILighting.IArgRate, ZigBee.ILighting.IArgRatePayload>(base.u8, ()=>({
                name: `Rate`,
                description: `increment/steps per second`,
                
            })),
            RateX: makeType<ZigBee.ILighting.IArgRateX, ZigBee.ILighting.IArgRateXPayload>(base.s16, ()=>({
                name: `Rate X`,
                description: `increment/steps per second`,
                
            })),
            RateY: makeType<ZigBee.ILighting.IArgRateY, ZigBee.ILighting.IArgRateYPayload>(base.s16, ()=>({
                name: `Rate Y`,
                description: `increment/steps per second`,
                
            })),
            RemainingTime: makeType<ZigBee.ILighting.IArgRemainingTime, ZigBee.ILighting.IArgRemainingTimePayload>(base.u16, ()=>({
                name: `Remaining time`,
                description: `holds the time remaining, in 1/10ths of a second, until the currently
active command will be complete`,
                id: 0x0002,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.Seconds,
                scale: 10,
                
            })),
            Saturation: makeType<ZigBee.ILighting.IArgSaturation, ZigBee.ILighting.IArgSaturationPayload>(base.u8, ()=>({
                name: `Saturation`,
                description: ``,
                
            })),
            StepMode: makeType<ZigBee.ILighting.IArgStepMode, ZigBee.ILighting.IArgStepModePayload>(base.enum8, ()=>({
                name: `Step mode`,
                description: ``,
                values: { 
                0x01: `Up`, 
                0x03: `Down`,  },
                
            })),
            StepSize: makeType<ZigBee.ILighting.IArgStepSize, ZigBee.ILighting.IArgStepSizePayload>(base.u8, ()=>({
                name: `Step size`,
                description: ``,
                
            })),
            StepX: makeType<ZigBee.ILighting.IArgStepX, ZigBee.ILighting.IArgStepXPayload>(base.s16, ()=>({
                name: `Step X`,
                description: ``,
                
            })),
            StepY: makeType<ZigBee.ILighting.IArgStepY, ZigBee.ILighting.IArgStepYPayload>(base.s16, ()=>({
                name: `Step Y`,
                description: ``,
                
            })),
            Time: makeType<ZigBee.ILighting.IArgTime, ZigBee.ILighting.IArgTimePayload>(base.u16, ()=>({
                name: `Time`,
                description: `Time in seconds used for a whole color loop.`,
                unit: units.Seconds,
                
            })),
            TransitionTime: makeType<ZigBee.ILighting.IArgTransitionTime, ZigBee.ILighting.IArgTransitionTimePayload>(base.u16, ()=>({
                name: `Transition time`,
                description: `The transition time in 1/10ths of a second.`,
                unit: units.Seconds,
                scale: 10,
                
            })),
            UpdateFlags: makeType<ZigBee.ILighting.IArgUpdateFlags, ZigBee.ILighting.IArgUpdateFlagsPayload>(base.bmp8, ()=>({
                name: `Update flags`,
                description: ``,
                bits: { 
                0: `Update action`, 
                1: `Update direction`, 
                2: `Update time`, 
                3: `Update start hue`,  },
                
            })),
            WhitePointX: makeType<ZigBee.ILighting.IArgWhitePointX, ZigBee.ILighting.IArgWhitePointXPayload>(base.u16, ()=>({
                name: `White Point X`,
                description: `contains the normalized chromaticity value x for this attribute, as
defined in the CIE xyY Color Space. x = value / 65536 (value
in the range 0 to 65279 inclusive)`,
                id: 0x0030,
                report: false,
                read: true,
                write: true,
                require: false,
                scale: 65536,
                
            })),
            WhitePointY: makeType<ZigBee.ILighting.IArgWhitePointY, ZigBee.ILighting.IArgWhitePointYPayload>(base.u16, ()=>({
                name: `White Point Y`,
                description: `contains the normalized chromaticity value y for this attribute, as
defined in the CIE xyY Color Space. y = value / 65536 (value
in the range 0 to 65279 inclusive)`,
                id: 0x0031,
                report: false,
                read: true,
                write: true,
                require: false,
                scale: 65536,
                
            })), },
        ColorControl: {
            ID: 0x0300,
            Name: `Color control`,
            Desc: `provides attributes and commands for controlling the color properties
of a color-capable light`,
            
            MoveToHue: makeType<ZigBee.ILighting.ColorControl.ICmdMoveToHue, ZigBee.ILighting.ColorControl.ICmdMoveToHuePayload>(command, () => ({
                name: `Move to hue`,
                description: ``,
                id: 0x0000,
                payload: { 
                    Hue: ZigBee.Lighting.Types.Hue,
                    MoveDirection: ZigBee.Lighting.Types.MoveDirection,
                    TransitionTime: ZigBee.Lighting.Types.TransitionTime,
                }
            })),

            MoveHue: makeType<ZigBee.ILighting.ColorControl.ICmdMoveHue, ZigBee.ILighting.ColorControl.ICmdMoveHuePayload>(command, () => ({
                name: `Move hue`,
                description: ``,
                id: 0x0001,
                payload: { 
                    MoveMode: ZigBee.Lighting.Types.MoveMode,
                    Rate: ZigBee.Lighting.Types.Rate,
                }
            })),

            StepHue: makeType<ZigBee.ILighting.ColorControl.ICmdStepHue, ZigBee.ILighting.ColorControl.ICmdStepHuePayload>(command, () => ({
                name: `Step hue`,
                description: ``,
                id: 0x0002,
                payload: { 
                    StepMode: ZigBee.Lighting.Types.StepMode,
                    StepSize: ZigBee.Lighting.Types.StepSize,
                    TransitionTime: ZigBee.Lighting.Types.TransitionTime,
                }
            })),

            MoveToSaturation: makeType<ZigBee.ILighting.ColorControl.ICmdMoveToSaturation, ZigBee.ILighting.ColorControl.ICmdMoveToSaturationPayload>(command, () => ({
                name: `Move to saturation`,
                description: ``,
                id: 0x0003,
                payload: { 
                    Saturation: ZigBee.Lighting.Types.Saturation,
                    TransitionTime: ZigBee.Lighting.Types.TransitionTime,
                }
            })),

            MoveSaturation: makeType<ZigBee.ILighting.ColorControl.ICmdMoveSaturation, ZigBee.ILighting.ColorControl.ICmdMoveSaturationPayload>(command, () => ({
                name: `Move saturation`,
                description: ``,
                id: 0x0004,
                payload: { 
                    MoveMode: ZigBee.Lighting.Types.MoveMode,
                    Rate: ZigBee.Lighting.Types.Rate,
                }
            })),

            StepSaturation: makeType<ZigBee.ILighting.ColorControl.ICmdStepSaturation, ZigBee.ILighting.ColorControl.ICmdStepSaturationPayload>(command, () => ({
                name: `Step saturation`,
                description: ``,
                id: 0x0005,
                payload: { 
                    StepMode: ZigBee.Lighting.Types.StepMode,
                    StepSize: ZigBee.Lighting.Types.StepSize,
                    TransitionTime: ZigBee.Lighting.Types.TransitionTime,
                }
            })),

            MoveToHueAndSaturation: makeType<ZigBee.ILighting.ColorControl.ICmdMoveToHueAndSaturation, ZigBee.ILighting.ColorControl.ICmdMoveToHueAndSaturationPayload>(command, () => ({
                name: `Move to hue and saturation`,
                description: ``,
                id: 0x0006,
                payload: { 
                    Hue: ZigBee.Lighting.Types.Hue,
                    Saturation: ZigBee.Lighting.Types.Saturation,
                    TransitionTime: ZigBee.Lighting.Types.TransitionTime,
                }
            })),

            MoveToColor: makeType<ZigBee.ILighting.ColorControl.ICmdMoveToColor, ZigBee.ILighting.ColorControl.ICmdMoveToColorPayload>(command, () => ({
                name: `Move to color`,
                description: ``,
                id: 0x0007,
                payload: { 
                    ColorX: ZigBee.Lighting.Types.ColorX,
                    ColorY: ZigBee.Lighting.Types.ColorY,
                    TransitionTime: ZigBee.Lighting.Types.TransitionTime,
                }
            })),

            MoveColor: makeType<ZigBee.ILighting.ColorControl.ICmdMoveColor, ZigBee.ILighting.ColorControl.ICmdMoveColorPayload>(command, () => ({
                name: `Move color`,
                description: ``,
                id: 0x0008,
                payload: { 
                    RateX: ZigBee.Lighting.Types.RateX,
                    RateY: ZigBee.Lighting.Types.RateY,
                }
            })),

            StepColor: makeType<ZigBee.ILighting.ColorControl.ICmdStepColor, ZigBee.ILighting.ColorControl.ICmdStepColorPayload>(command, () => ({
                name: `Step color`,
                description: ``,
                id: 0x0009,
                payload: { 
                    StepX: ZigBee.Lighting.Types.StepX,
                    StepY: ZigBee.Lighting.Types.StepY,
                    TransitionTime: ZigBee.Lighting.Types.TransitionTime,
                }
            })),

            MoveToColorTemperature: makeType<ZigBee.ILighting.ColorControl.ICmdMoveToColorTemperature, ZigBee.ILighting.ColorControl.ICmdMoveToColorTemperaturePayload>(command, () => ({
                name: `Move to color temperature`,
                description: ``,
                id: 0x000A,
                payload: { 
                    ColorTemperature: ZigBee.Lighting.Types.ColorTemperature,
                    TransitionTime: ZigBee.Lighting.Types.TransitionTime,
                }
            })),

            EnhancedMoveToHue: makeType<ZigBee.ILighting.ColorControl.ICmdEnhancedMoveToHue, ZigBee.ILighting.ColorControl.ICmdEnhancedMoveToHuePayload>(command, () => ({
                name: `Enhanced move to hue`,
                description: ``,
                id: 0x0040,
                payload: { 
                    EnhancedHue: ZigBee.Lighting.Types.EnhancedHue,
                    MoveDirection: ZigBee.Lighting.Types.MoveDirection,
                    TransitionTime: ZigBee.Lighting.Types.TransitionTime,
                }
            })),

            EnhancedMoveHue: makeType<ZigBee.ILighting.ColorControl.ICmdEnhancedMoveHue, ZigBee.ILighting.ColorControl.ICmdEnhancedMoveHuePayload>(command, () => ({
                name: `Enhanced move hue`,
                description: ``,
                id: 0x0041,
                payload: { 
                    MoveMode: ZigBee.Lighting.Types.MoveMode,
                    Rate: ZigBee.Lighting.Types.Rate,
                }
            })),

            EnhancedStepHue: makeType<ZigBee.ILighting.ColorControl.ICmdEnhancedStepHue, ZigBee.ILighting.ColorControl.ICmdEnhancedStepHuePayload>(command, () => ({
                name: `Enhanced step hue`,
                description: ``,
                id: 0x0042,
                payload: { 
                    StepMode: ZigBee.Lighting.Types.StepMode,
                    StepSize: ZigBee.Lighting.Types.StepSize,
                    TransitionTime: ZigBee.Lighting.Types.TransitionTime,
                }
            })),

            EnhancedMoveToHueAndSaturation: makeType<ZigBee.ILighting.ColorControl.ICmdEnhancedMoveToHueAndSaturation, ZigBee.ILighting.ColorControl.ICmdEnhancedMoveToHueAndSaturationPayload>(command, () => ({
                name: `Enhanced move to hue and saturation`,
                description: ``,
                id: 0x0043,
                payload: { 
                    EnhancedHue: ZigBee.Lighting.Types.EnhancedHue,
                    Saturation: ZigBee.Lighting.Types.Saturation,
                    TransitionTime: ZigBee.Lighting.Types.TransitionTime,
                }
            })),

            ColorLoopSet: makeType<ZigBee.ILighting.ColorControl.ICmdColorLoopSet, ZigBee.ILighting.ColorControl.ICmdColorLoopSetPayload>(command, () => ({
                name: `Color loop set`,
                description: ``,
                id: 0x0044,
                payload: { 
                    UpdateFlags: ZigBee.Lighting.Types.UpdateFlags,
                    Action: ZigBee.Lighting.Types.Action,
                    HueDirection: ZigBee.Lighting.Types.HueDirection,
                    Time: ZigBee.Lighting.Types.Time,
                    EnhancedHue: ZigBee.Lighting.Types.EnhancedHue,
                }
            })),

            StopMoveStep: makeType<ZigBee.ILighting.ColorControl.ICmdStopMoveStep, ZigBee.ILighting.ColorControl.ICmdStopMoveStepPayload>(command, () => ({
                name: `Stop move step`,
                description: `Stops move to and step commands. It has no effect on a active
color loop.`,
                id: 0x0047,
                payload: {}
            })),

            MoveColorTemperature: makeType<ZigBee.ILighting.ColorControl.ICmdMoveColorTemperature, ZigBee.ILighting.ColorControl.ICmdMoveColorTemperaturePayload>(command, () => ({
                name: `Move color temperature`,
                description: ``,
                id: 0x004B,
                payload: { 
                    MoveMode: ZigBee.Lighting.Types.MoveMode,
                    Rate: ZigBee.Lighting.Types.Rate,
                    ColorTemperatureMin: ZigBee.Lighting.Types.ColorTemperatureMin,
                    ColorTemperatureMax: ZigBee.Lighting.Types.ColorTemperatureMax,
                }
            })),

            StepColorTemperature: makeType<ZigBee.ILighting.ColorControl.ICmdStepColorTemperature, ZigBee.ILighting.ColorControl.ICmdStepColorTemperaturePayload>(command, () => ({
                name: `Step color temperature`,
                description: ``,
                id: 0x004C,
                payload: { 
                    StepMode: ZigBee.Lighting.Types.StepMode,
                    StepSize: ZigBee.Lighting.Types.StepSize,
                    TransitionTime: ZigBee.Lighting.Types.TransitionTime,
                    ColorTemperatureMinMireds: ZigBee.Lighting.Types.ColorTemperatureMinMireds,
                    ColorTemperatureMaxMireds: ZigBee.Lighting.Types.ColorTemperatureMaxMireds,
                }
            })),

            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        BallastConfiguration: {
            ID: 0x0301,
            Name: `Ballast Configuration`,
            Desc: `provides attributes and commands to configure a ballast.`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        }
    };
    
    ZigBee.Lighting.ColorControl.Server.Attribute = { 
        0x0000: ZigBee.Lighting.Types.CurrentHue,
        0x0001: ZigBee.Lighting.Types.CurrentSaturation,
        0x0002: ZigBee.Lighting.Types.RemainingTime,
        0x0003: ZigBee.Lighting.Types.CurrentX,
        0x0004: ZigBee.Lighting.Types.CurrentY,
        0x0005: ZigBee.Lighting.Types.DriftCompensation,
        0x0006: ZigBee.Lighting.Types.CompensationText,
        0x0007: ZigBee.Lighting.Types.ColorTemperatureMireds,
        0x0008: ZigBee.Lighting.Types.ColorMode,
        0x000F: ZigBee.Lighting.Types.ColorControlOptions,
        0x4000: ZigBee.Lighting.Types.EnhancedCurrentHue,
        0x4001: ZigBee.Lighting.Types.EnhancedColorMode,
        0x4002: ZigBee.Lighting.Types.ColorLoopActive,
        0x4003: ZigBee.Lighting.Types.ColorLoopDirection,
        0x4004: ZigBee.Lighting.Types.ColorLoopTime,
        0x4005: ZigBee.Lighting.Types.ColorLoopStartEnhancedHue,
        0x4006: ZigBee.Lighting.Types.ColorLoopStoredEnhancedHue,
        0x400A: ZigBee.Lighting.Types.ColorCapabilities,
        0x400B: ZigBee.Lighting.Types.ColorTemperaturePhysicalMinMireds,
        0x400C: ZigBee.Lighting.Types.ColorTemperaturePhysicalMaxMireds,
        0x400D: ZigBee.Lighting.Types.CoupleColorTempToLevelMinMireds,
        0x4010: ZigBee.Lighting.Types.PowerOnColorTemperature,
        0x0010: ZigBee.Lighting.Types.NumberOfPrimaries,
        0x0011: ZigBee.Lighting.Types.Primary1X,
        0x0012: ZigBee.Lighting.Types.Primary1Y,
        0x0013: ZigBee.Lighting.Types.Primary1Intensity,
        0x0015: ZigBee.Lighting.Types.Primary2X,
        0x0016: ZigBee.Lighting.Types.Primary2Y,
        0x0017: ZigBee.Lighting.Types.Primary2Intensity,
        0x0019: ZigBee.Lighting.Types.Primary3X,
        0x001A: ZigBee.Lighting.Types.Primary3Y,
        0x001B: ZigBee.Lighting.Types.Primary3Intensity,
        0x0020: ZigBee.Lighting.Types.Primary4X,
        0x0021: ZigBee.Lighting.Types.Primary4Y,
        0x0022: ZigBee.Lighting.Types.Primary4Intensity,
        0x0024: ZigBee.Lighting.Types.Primary5X,
        0x0025: ZigBee.Lighting.Types.Primary5Y,
        0x0026: ZigBee.Lighting.Types.Primary5Intensity,
        0x0028: ZigBee.Lighting.Types.Primary6X,
        0x0029: ZigBee.Lighting.Types.Primary6Y,
        0x002A: ZigBee.Lighting.Types.Primary6Intensity,
        0x0030: ZigBee.Lighting.Types.WhitePointX,
        0x0031: ZigBee.Lighting.Types.WhitePointY,
        0x0032: ZigBee.Lighting.Types.ColorPointRedX,
        0x0033: ZigBee.Lighting.Types.ColorPointRedY,
        0x0034: ZigBee.Lighting.Types.ColorPointRedIntensity,
        0x0036: ZigBee.Lighting.Types.ColorPointGreenX,
        0x0037: ZigBee.Lighting.Types.ColorPointGreenY,
        0x0038: ZigBee.Lighting.Types.ColorPointGreenIntensity,
        0x003A: ZigBee.Lighting.Types.ColorPointBlueX,
        0x003B: ZigBee.Lighting.Types.ColorPointBlueY,
        0x003C: ZigBee.Lighting.Types.ColorPointBlueIntensity,
    };
    ZigBee.Lighting.ColorControl.Client.Attribute = { 
    };
    ZigBee.Lighting.ColorControl.Server.Command = { 
        0x00: ZigBee.Lighting.ColorControl.MoveToHue,
        0x01: ZigBee.Lighting.ColorControl.MoveHue,
        0x02: ZigBee.Lighting.ColorControl.StepHue,
        0x03: ZigBee.Lighting.ColorControl.MoveToSaturation,
        0x04: ZigBee.Lighting.ColorControl.MoveSaturation,
        0x05: ZigBee.Lighting.ColorControl.StepSaturation,
        0x06: ZigBee.Lighting.ColorControl.MoveToHueAndSaturation,
        0x07: ZigBee.Lighting.ColorControl.MoveToColor,
        0x08: ZigBee.Lighting.ColorControl.MoveColor,
        0x09: ZigBee.Lighting.ColorControl.StepColor,
        0x0A: ZigBee.Lighting.ColorControl.MoveToColorTemperature,
        0x40: ZigBee.Lighting.ColorControl.EnhancedMoveToHue,
        0x41: ZigBee.Lighting.ColorControl.EnhancedMoveHue,
        0x42: ZigBee.Lighting.ColorControl.EnhancedStepHue,
        0x43: ZigBee.Lighting.ColorControl.EnhancedMoveToHueAndSaturation,
        0x44: ZigBee.Lighting.ColorControl.ColorLoopSet,
        0x47: ZigBee.Lighting.ColorControl.StopMoveStep,
        0x4B: ZigBee.Lighting.ColorControl.MoveColorTemperature,
        0x4C: ZigBee.Lighting.ColorControl.StepColorTemperature,
    };
    ZigBee.Lighting.ColorControl.Client.Command = { 
    };
    ZigBee.Lighting.BallastConfiguration.Server.Attribute = { 
        0x0000: ZigBee.Lighting.Types.PhysicalMinLevel,
        0x0001: ZigBee.Lighting.Types.PhysicalMaxLevel,
        0x0002: ZigBee.Lighting.Types.BallastStatus,
        0x0010: ZigBee.Lighting.Types.MinLevel,
        0x0011: ZigBee.Lighting.Types.MaxLevel,
        0x0012: ZigBee.Lighting.Types.PowerOnLevel,
        0x0013: ZigBee.Lighting.Types.PowerOnTime,
        0x0014: ZigBee.Lighting.Types.IntrinsicBallastFactor,
        0x0015: ZigBee.Lighting.Types.BallastFactorAdjustment,
        0x0020: ZigBee.Lighting.Types.LampQuantity,
        0x0030: ZigBee.Lighting.Types.LampType,
        0x0031: ZigBee.Lighting.Types.LampManufacturer,
        0x0032: ZigBee.Lighting.Types.LampRatedHours,
        0x0033: ZigBee.Lighting.Types.LampBurnHours,
        0x0034: ZigBee.Lighting.Types.LampAlarmMode,
    };
    ZigBee.Lighting.BallastConfiguration.Client.Attribute = { 
    };
    ZigBee.Lighting.BallastConfiguration.Server.Command = { 
    };
    ZigBee.Lighting.BallastConfiguration.Client.Command = { 
    };
    export const MeasurementAndSensing = {
        Types: { 
            AcActivePowerOverload: makeType<ZigBee.IMeasurementAndSensing.IArgAcActivePowerOverload, ZigBee.IMeasurementAndSensing.IArgAcActivePowerOverloadPayload>(base.s16, ()=>({
                name: `AC Active Power Overload`,
                description: `specifies the alarm threshold, set by the manufacturer, for the maximum output active power supported by
device. The value is multiplied and divided by the ACPowerMultiplier and ACPowerDivisor, respectively.`,
                id: 0x0803,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Watts,
                
            })),
            AcAlarmsMask: makeType<ZigBee.IMeasurementAndSensing.IArgAcAlarmsMask, ZigBee.IMeasurementAndSensing.IArgAcAlarmsMaskPayload>(base.bmp16, ()=>({
                name: `AC Alarms Mask`,
                description: `specifies which configurable alarms may be generated`,
                id: 0x0800,
                report: false,
                read: true,
                write: true,
                require: false,
                bits: { 
                0: `Voltage Overload`, 
                1: `Current Overload`, 
                2: `Active Power Overload`, 
                3: `Reactive Power Overload`, 
                4: `Average RMS Over Voltage`, 
                5: `Average RMS Under Voltage`, 
                6: `RMS Extreme Over Voltage`, 
                7: `RMS Extreme Under Voltage`, 
                8: `RMS Voltage Sag`, 
                9: `RMS Voltage Swell`,  },
                
            })),
            AcCurrentDivisor: makeType<ZigBee.IMeasurementAndSensing.IArgAcCurrentDivisor, ZigBee.IMeasurementAndSensing.IArgAcCurrentDivisorPayload>(base.u16, ()=>({
                name: `AC Current Divisor`,
                description: `Provides a value to be divided against the ACCurrent, InstantaneousCurrent and RMSCurrent attributes.
This attribute must be used in conjunction with the ACCurrentMultiplier attribute. 0x0000 is an invalid value
for this attribute`,
                id: 0x0603,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            AcCurrentMultiplier: makeType<ZigBee.IMeasurementAndSensing.IArgAcCurrentMultiplier, ZigBee.IMeasurementAndSensing.IArgAcCurrentMultiplierPayload>(base.u16, ()=>({
                name: `AC Current Multiplier`,
                description: `Provides a value to be multiplied against the InstantaneousCurrent and RMSCurrent attributes. This attribute
must be used in conjunction with the ACCurrentDivisor attribute. 0x0000 is an invalid value for this attribute.`,
                id: 0x0602,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            AcCurrentOverload: makeType<ZigBee.IMeasurementAndSensing.IArgAcCurrentOverload, ZigBee.IMeasurementAndSensing.IArgAcCurrentOverloadPayload>(base.s16, ()=>({
                name: `AC Current Overload`,
                description: `specifies the alarm threshold, set by the manufacturer, for the maximum output current supported by device.
The value is multiplied and divided by the ACCurrentMultiplier and ACCurrentDivider, respectively. The
value is current RMS.`,
                id: 0x0802,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Amperes,
                
            })),
            AcFrequency: makeType<ZigBee.IMeasurementAndSensing.IArgAcFrequency, ZigBee.IMeasurementAndSensing.IArgAcFrequencyPayload>(base.u16, ()=>({
                name: `AC Frequency`,
                description: `represents the most recent AC Frequency reading in Hertz (Hz). If the frequency
cannot be measured, a value of 0xFFFF is returned.`,
                id: 0x0300,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Hertz,
                
            })),
            AcFrequencyDivisor: makeType<ZigBee.IMeasurementAndSensing.IArgAcFrequencyDivisor, ZigBee.IMeasurementAndSensing.IArgAcFrequencyDivisorPayload>(base.u16, ()=>({
                name: `AC Frequency Divisor`,
                description: `Provides a value to be divided against the ACFrequency attribute. This attribute must be used in conjunction
with the ACFrequencyMultiplier attribute. 0x0000 is an invalid value for this attribute.`,
                id: 0x0401,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            AcFrequencyMax: makeType<ZigBee.IMeasurementAndSensing.IArgAcFrequencyMax, ZigBee.IMeasurementAndSensing.IArgAcFrequencyMaxPayload>(base.u16, ()=>({
                name: `AC Frequency Max`,
                description: `attribute represents the highest AC Frequency value measured in Hertz (Hz). After
resetting, this attribute will return a value of 0xFFFF until a measurement is made.`,
                id: 0x0302,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Hertz,
                
            })),
            AcFrequencyMin: makeType<ZigBee.IMeasurementAndSensing.IArgAcFrequencyMin, ZigBee.IMeasurementAndSensing.IArgAcFrequencyMinPayload>(base.u16, ()=>({
                name: `AC Frequency Min`,
                description: `attribute represents the lowest AC Frequency value measured in Hertz (Hz). After
resetting, this attribute will return a value of 0xFFFF until a measurement is made.`,
                id: 0x0301,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Hertz,
                
            })),
            AcFrequencyMultiplier: makeType<ZigBee.IMeasurementAndSensing.IArgAcFrequencyMultiplier, ZigBee.IMeasurementAndSensing.IArgAcFrequencyMultiplierPayload>(base.u16, ()=>({
                name: `AC Frequency Multiplier`,
                description: `Provides a value to be multiplied against the ACFrequency attribute.
This attribute must be used in conjunction with the ACFrequencyDivisor attribute.
0x0000 is an invalid value for this attribute.`,
                id: 0x0400,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            AcPowerDivisor: makeType<ZigBee.IMeasurementAndSensing.IArgAcPowerDivisor, ZigBee.IMeasurementAndSensing.IArgAcPowerDivisorPayload>(base.u16, ()=>({
                name: `AC Power Divisor`,
                description: `Provides a value to be divided against the InstantaneousPower and ActivePower attributes. This attribute
must be used in conjunction with the ACPowerMultiplier attribute. 0x0000 is an invalid value for this attribute.`,
                id: 0x0605,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            AcPowerMultiplier: makeType<ZigBee.IMeasurementAndSensing.IArgAcPowerMultiplier, ZigBee.IMeasurementAndSensing.IArgAcPowerMultiplierPayload>(base.u16, ()=>({
                name: `AC Power Multiplier`,
                description: `Provides a value to be multiplied against the InstantaneousPower and ActivePower attributes. This attribute
must be used in conjunction with the ACPowerDivisor attribute. 0x0000 is an invalid value for this attribute.`,
                id: 0x0604,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            AcReactivePowerOverload: makeType<ZigBee.IMeasurementAndSensing.IArgAcReactivePowerOverload, ZigBee.IMeasurementAndSensing.IArgAcReactivePowerOverloadPayload>(base.s16, ()=>({
                name: `AC Reactive Power Overload`,
                description: `specifies the alarm threshold, set by the manufacturer, for the maximum output reactive power supported by
device. The value is multiplied and divided by the ACPowerMultiplier and ACPowerDivisor, respectively.`,
                id: 0x0804,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            AcVoltageDivisor: makeType<ZigBee.IMeasurementAndSensing.IArgAcVoltageDivisor, ZigBee.IMeasurementAndSensing.IArgAcVoltageDivisorPayload>(base.u16, ()=>({
                name: `AC Voltage Divisor`,
                description: `Provides a value to be divided against the InstantaneousVoltage and RMSVoltage attributes. This attribute
must be used in conjunction with the ACVoltageMultiplier attribute. 0x0000 is an invalid value for this
attribute.`,
                id: 0x0601,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            AcVoltageMultiplier: makeType<ZigBee.IMeasurementAndSensing.IArgAcVoltageMultiplier, ZigBee.IMeasurementAndSensing.IArgAcVoltageMultiplierPayload>(base.u16, ()=>({
                name: `AC Voltage Multiplier`,
                description: `Provides a value to be multiplied against the InstantaneousVoltage and RMSVoltage attributes. This attribute
must be used in conjunction with the ACVoltageDivisor attribute. 0x0000 is an invalid value for this attribute.`,
                id: 0x0600,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            AcVoltageOverload: makeType<ZigBee.IMeasurementAndSensing.IArgAcVoltageOverload, ZigBee.IMeasurementAndSensing.IArgAcVoltageOverloadPayload>(base.s16, ()=>({
                name: `AC Voltage Overload`,
                description: `specifies the alarm threshold, set by the manufacturer, for the maximum output voltage supported by device.
The value is multiplied and divided by the ACVoltageMultiplier the ACVoltageDivisor, respectively. The
value is voltage RMS.`,
                id: 0x0801,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Volts,
                
            })),
            ActiveCurrent: makeType<ZigBee.IMeasurementAndSensing.IArgActiveCurrent, ZigBee.IMeasurementAndSensing.IArgActiveCurrentPayload>(base.s16, ()=>({
                name: `Active Current`,
                description: `Represents the single phase or Phase A, AC active/resistive current value at the moment in time the attribute
is read, in Amps (A). Positive values indicate power delivered to the premises where negative values indicate
power received from the premises.`,
                id: 0x0502,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Amperes,
                
            })),
            ActiveCurrentPhB: makeType<ZigBee.IMeasurementAndSensing.IArgActiveCurrentPhB, ZigBee.IMeasurementAndSensing.IArgActiveCurrentPhBPayload>(base.s16, ()=>({
                name: `Active Current Ph B`,
                description: `represents the Phase B, AC active/resistive current value at the moment in time the attribute is read.
Positive values indicate power delivered to the premises where negative values indicate power received
from the premises.`,
                id: 0x0902,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Amperes,
                
            })),
            ActiveCurrentPhC: makeType<ZigBee.IMeasurementAndSensing.IArgActiveCurrentPhC, ZigBee.IMeasurementAndSensing.IArgActiveCurrentPhCPayload>(base.s16, ()=>({
                name: `Active Current Ph C`,
                description: ``,
                id: 0x0A02,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            ActivePower: makeType<ZigBee.IMeasurementAndSensing.IArgActivePower, ZigBee.IMeasurementAndSensing.IArgActivePowerPayload>(base.s16, ()=>({
                name: `Active Power`,
                description: `Represents the single phase or Phase A, current demand of active power delivered or received at the premises.
Positive values indicate power delivered to the premises where negative values indicate power received from the
premises.`,
                id: 0x050B,
                report: true,
                read: false,
                write: false,
                require: false,
                unit: units.Watts,
                
            })),
            ActivePowerMax: makeType<ZigBee.IMeasurementAndSensing.IArgActivePowerMax, ZigBee.IMeasurementAndSensing.IArgActivePowerMaxPayload>(base.s16, ()=>({
                name: `Active Power Max`,
                description: `Represents the highest AC power value`,
                id: 0x050D,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Watts,
                
            })),
            ActivePowerMaxPhB: makeType<ZigBee.IMeasurementAndSensing.IArgActivePowerMaxPhB, ZigBee.IMeasurementAndSensing.IArgActivePowerMaxPhBPayload>(base.s16, ()=>({
                name: `Active Power Max Ph B`,
                description: ``,
                id: 0x090D,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            ActivePowerMaxPhC: makeType<ZigBee.IMeasurementAndSensing.IArgActivePowerMaxPhC, ZigBee.IMeasurementAndSensing.IArgActivePowerMaxPhCPayload>(base.s16, ()=>({
                name: `Active Power Max Ph C`,
                description: ``,
                id: 0x0A0D,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            ActivePowerMin: makeType<ZigBee.IMeasurementAndSensing.IArgActivePowerMin, ZigBee.IMeasurementAndSensing.IArgActivePowerMinPayload>(base.s16, ()=>({
                name: `Active Power Min`,
                description: `Represents the lowest AC power value`,
                id: 0x050C,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Watts,
                
            })),
            ActivePowerMinPhB: makeType<ZigBee.IMeasurementAndSensing.IArgActivePowerMinPhB, ZigBee.IMeasurementAndSensing.IArgActivePowerMinPhBPayload>(base.s16, ()=>({
                name: `Active Power Min Ph B`,
                description: ``,
                id: 0x090C,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            ActivePowerMinPhC: makeType<ZigBee.IMeasurementAndSensing.IArgActivePowerMinPhC, ZigBee.IMeasurementAndSensing.IArgActivePowerMinPhCPayload>(base.s16, ()=>({
                name: `Active Power Min Ph C`,
                description: ``,
                id: 0x0A0C,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            ActivePowerPhB: makeType<ZigBee.IMeasurementAndSensing.IArgActivePowerPhB, ZigBee.IMeasurementAndSensing.IArgActivePowerPhBPayload>(base.s16, ()=>({
                name: `Active Power Ph B`,
                description: ``,
                id: 0x090B,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            ActivePowerPhC: makeType<ZigBee.IMeasurementAndSensing.IArgActivePowerPhC, ZigBee.IMeasurementAndSensing.IArgActivePowerPhCPayload>(base.s16, ()=>({
                name: `Active Power Ph C`,
                description: ``,
                id: 0x0A0B,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            ApparentPower: makeType<ZigBee.IMeasurementAndSensing.IArgApparentPower, ZigBee.IMeasurementAndSensing.IArgApparentPowerPayload>(base.u16, ()=>({
                name: `Apparent Power`,
                description: `Represents the single phase or Phase A, current demand of apparent (Square root of active and reactive
power) power, in VA.`,
                id: 0x050F,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.VoltAmperes,
                
            })),
            ApparentPowerPhB: makeType<ZigBee.IMeasurementAndSensing.IArgApparentPowerPhB, ZigBee.IMeasurementAndSensing.IArgApparentPowerPhBPayload>(base.u16, ()=>({
                name: `Apparent Power Ph B`,
                description: ``,
                id: 0x090F,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            ApparentPowerPhC: makeType<ZigBee.IMeasurementAndSensing.IArgApparentPowerPhC, ZigBee.IMeasurementAndSensing.IArgApparentPowerPhCPayload>(base.u16, ()=>({
                name: `Apparent Power Ph C`,
                description: ``,
                id: 0x0A0F,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            AttributeId: makeType<ZigBee.IMeasurementAndSensing.IArgAttributeId, ZigBee.IMeasurementAndSensing.IArgAttributeIdPayload>(base.u16, ()=>({
                name: `Attribute ID`,
                description: ``,
                
            })),
            AverageRmsOverVoltage: makeType<ZigBee.IMeasurementAndSensing.IArgAverageRmsOverVoltage, ZigBee.IMeasurementAndSensing.IArgAverageRmsOverVoltagePayload>(base.s16, ()=>({
                name: `Average RMS over-voltage`,
                description: `is the average RMS voltage above which an over voltage condition is reported.
The threshold shall be configurable within the specified operating range of the electricity meter.
The value is multiplied and divided by the ACVoltageMultiplier and ACVoltageDivisor, respectively.`,
                id: 0x0805,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Volts,
                
            })),
            AverageRmsOverVoltageCounter: makeType<ZigBee.IMeasurementAndSensing.IArgAverageRmsOverVoltageCounter, ZigBee.IMeasurementAndSensing.IArgAverageRmsOverVoltageCounterPayload>(base.u16, ()=>({
                name: `Average RMS Over Voltage Counter`,
                description: `is the number of times the average RMS voltage, has been above the AverageRMS OverVoltage threshold
since last reset. This counter may be reset by writing zero to the attribute.`,
                id: 0x0512,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            AverageRmsOverVoltageCounterPhB: makeType<ZigBee.IMeasurementAndSensing.IArgAverageRmsOverVoltageCounterPhB, ZigBee.IMeasurementAndSensing.IArgAverageRmsOverVoltageCounterPhBPayload>(base.u16, ()=>({
                name: `Average RMS Over Voltage Counter Ph B`,
                description: ``,
                id: 0x0912,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            AverageRmsOverVoltageCounterPhC: makeType<ZigBee.IMeasurementAndSensing.IArgAverageRmsOverVoltageCounterPhC, ZigBee.IMeasurementAndSensing.IArgAverageRmsOverVoltageCounterPhCPayload>(base.u16, ()=>({
                name: `Average RMS Over Voltage Counter Ph C`,
                description: ``,
                id: 0x0A12,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            AverageRmsUnderVoltage: makeType<ZigBee.IMeasurementAndSensing.IArgAverageRmsUnderVoltage, ZigBee.IMeasurementAndSensing.IArgAverageRmsUnderVoltagePayload>(base.s16, ()=>({
                name: `Average RMS under-voltage`,
                description: `is the average RMS voltage below which an under voltage condition is reported.
The threshold shall be configurable within the specified operating range of the electricity meter.
The value is multiplied and divided by the ACVoltageMultiplier and ACVoltageDivisor, respectively.`,
                id: 0x0806,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Volts,
                
            })),
            AverageRmsUnderVoltageCounter: makeType<ZigBee.IMeasurementAndSensing.IArgAverageRmsUnderVoltageCounter, ZigBee.IMeasurementAndSensing.IArgAverageRmsUnderVoltageCounterPayload>(base.u16, ()=>({
                name: `Average RMS Under Voltage Counter`,
                description: `is the number of times the average RMS voltage, has been below the AverageRMS underVoltage threshold
since last reset. This counter may be reset by writing zero to the attribute.`,
                id: 0x0513,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            AverageRmsUnderVoltageCounterPhB: makeType<ZigBee.IMeasurementAndSensing.IArgAverageRmsUnderVoltageCounterPhB, ZigBee.IMeasurementAndSensing.IArgAverageRmsUnderVoltageCounterPhBPayload>(base.u16, ()=>({
                name: `Average RMS Under Voltage Counter Ph B`,
                description: ``,
                id: 0x0913,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            AverageRmsUnderVoltageCounterPhC: makeType<ZigBee.IMeasurementAndSensing.IArgAverageRmsUnderVoltageCounterPhC, ZigBee.IMeasurementAndSensing.IArgAverageRmsUnderVoltageCounterPhCPayload>(base.u16, ()=>({
                name: `Average RMS Under Voltage Counter Ph C`,
                description: ``,
                id: 0x0A13,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            AverageRmsVoltageMeasurementPeriod: makeType<ZigBee.IMeasurementAndSensing.IArgAverageRmsVoltageMeasurementPeriod, ZigBee.IMeasurementAndSensing.IArgAverageRmsVoltageMeasurementPeriodPayload>(base.u16, ()=>({
                name: `Average RMS Voltage Measurement Period`,
                description: `is the period that the RMS voltage is averaged over.`,
                id: 0x0511,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Seconds,
                
            })),
            AverageRmsVoltageMeasurementPeriodPhB: makeType<ZigBee.IMeasurementAndSensing.IArgAverageRmsVoltageMeasurementPeriodPhB, ZigBee.IMeasurementAndSensing.IArgAverageRmsVoltageMeasurementPeriodPhBPayload>(base.u16, ()=>({
                name: `Average RMS Voltage Measurement Period Ph B`,
                description: ``,
                id: 0x0911,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            AverageRmsVoltageMeasurementPeriodPhC: makeType<ZigBee.IMeasurementAndSensing.IArgAverageRmsVoltageMeasurementPeriodPhC, ZigBee.IMeasurementAndSensing.IArgAverageRmsVoltageMeasurementPeriodPhCPayload>(base.u16, ()=>({
                name: `Average RMS Voltage Measurement Period Ph C`,
                description: ``,
                id: 0x0A11,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            ConcentrationTolerance: makeType<ZigBee.IMeasurementAndSensing.IArgConcentrationTolerance, ZigBee.IMeasurementAndSensing.IArgConcentrationTolerancePayload>(base.float, ()=>({
                name: `Concentration Tolerance`,
                description: `indicates the magnitude of the possible error that is associated with MeasuredValue.
The true value is located in the range (MeasuredValue – Tolerance) to
(MeasuredValue + Tolerance)`,
                id: 0x0003,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.Concentration,
                
            })),
            DcCurrent: makeType<ZigBee.IMeasurementAndSensing.IArgDcCurrent, ZigBee.IMeasurementAndSensing.IArgDcCurrentPayload>(base.s16, ()=>({
                name: `DC Current`,
                description: `represents the most recent DC current reading in Amps (A). If the current cannot be
measured, a value of 0x8000 is returned.`,
                id: 0x0103,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Amperes,
                
            })),
            DcCurrentDivisor: makeType<ZigBee.IMeasurementAndSensing.IArgDcCurrentDivisor, ZigBee.IMeasurementAndSensing.IArgDcCurrentDivisorPayload>(base.u16, ()=>({
                name: `DC Current Divisor`,
                description: `provides a value to be divided against the DCCurrent, DCCurrentMin, and DCCurrentMax attributes.
This attribute must be used in conjunction with the DCCurrentMultiplier attribute.
0x0000 is an invalid value for this attribute.`,
                id: 0x0203,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            DcCurrentMax: makeType<ZigBee.IMeasurementAndSensing.IArgDcCurrentMax, ZigBee.IMeasurementAndSensing.IArgDcCurrentMaxPayload>(base.s16, ()=>({
                name: `DC Current Max`,
                description: `represents the highest DC current value measured in Amps (A). After resetting,
this attribute will return a value of 0x8000 until a measurement is made.`,
                id: 0x0105,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Amperes,
                
            })),
            DcCurrentMin: makeType<ZigBee.IMeasurementAndSensing.IArgDcCurrentMin, ZigBee.IMeasurementAndSensing.IArgDcCurrentMinPayload>(base.s16, ()=>({
                name: `DC Current Min`,
                description: `represents the lowest DC current value measured in Amps (A). After resetting,
this attribute will return a value of 0x8000 until a measurement is made.`,
                id: 0x0104,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Amperes,
                
            })),
            DcCurrentMultiplier: makeType<ZigBee.IMeasurementAndSensing.IArgDcCurrentMultiplier, ZigBee.IMeasurementAndSensing.IArgDcCurrentMultiplierPayload>(base.u16, ()=>({
                name: `DC Current Multiplier`,
                description: `provides a value to be multiplied against the DCCurrent, DCCurrentMin, and DCCurrentMax attributes.
This attribute must be used in conjunction with the DCCurrentDivisor attribute.
0x0000 is an invalid value for this attribute.`,
                id: 0x0202,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            DcCurrentOverload: makeType<ZigBee.IMeasurementAndSensing.IArgDcCurrentOverload, ZigBee.IMeasurementAndSensing.IArgDcCurrentOverloadPayload>(base.s16, ()=>({
                name: `DC Current Overload`,
                description: `Specifies the alarm threshold, set by the manufacturer, for the maximum output current supported by device.
The value is multiplied and divided by the DCCurrentMultiplier and DCCurrentDivider respectively.`,
                id: 0x0702,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Amperes,
                
            })),
            DcOverloadAlarmsMask: makeType<ZigBee.IMeasurementAndSensing.IArgDcOverloadAlarmsMask, ZigBee.IMeasurementAndSensing.IArgDcOverloadAlarmsMaskPayload>(base.bmp8, ()=>({
                name: `DC Overload Alarms Mask`,
                description: `Specifies which configurable alarms may be generated`,
                id: 0x0700,
                report: false,
                read: true,
                write: true,
                require: false,
                bits: { 
                0: `Voltage Overload`, 
                1: `Current Overload`,  },
                
            })),
            DcPower: makeType<ZigBee.IMeasurementAndSensing.IArgDcPower, ZigBee.IMeasurementAndSensing.IArgDcPowerPayload>(base.s16, ()=>({
                name: `DC Power`,
                description: `represents the most recent DC power reading in Watts (W). If the power cannot be
measured, a value of 0x8000 is returned.`,
                id: 0x0106,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Watts,
                
            })),
            DcPowerDivisor: makeType<ZigBee.IMeasurementAndSensing.IArgDcPowerDivisor, ZigBee.IMeasurementAndSensing.IArgDcPowerDivisorPayload>(base.u16, ()=>({
                name: `DC Power Divisor`,
                description: `provides a value to be divided against the DCPower, DCPowerMin, and DCPowerMax attributes.
This attribute must be used in conjunction with the DCPowerMultiplier attribute.
0x0000 is an invalid value for this attribute`,
                id: 0x0205,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            DcPowerMax: makeType<ZigBee.IMeasurementAndSensing.IArgDcPowerMax, ZigBee.IMeasurementAndSensing.IArgDcPowerMaxPayload>(base.s16, ()=>({
                name: `DC Power Max`,
                description: `represents the highest DC power value measured in Watts (W). After resetting,
this attribute will return a value of 0x8000 until a measurement is made.`,
                id: 0x0108,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Watts,
                
            })),
            DcPowerMin: makeType<ZigBee.IMeasurementAndSensing.IArgDcPowerMin, ZigBee.IMeasurementAndSensing.IArgDcPowerMinPayload>(base.s16, ()=>({
                name: `DC Power Min`,
                description: `represents the lowest DC power value measured in Watts (W). After resetting,
this attribute will return a value of 0x8000 until a measurement is made.`,
                id: 0x0107,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Watts,
                
            })),
            DcPowerMultiplier: makeType<ZigBee.IMeasurementAndSensing.IArgDcPowerMultiplier, ZigBee.IMeasurementAndSensing.IArgDcPowerMultiplierPayload>(base.u16, ()=>({
                name: `DC Power Multiplier`,
                description: `provides a value to be multiplied against the DCPower, DCPowerMin, and DCPowerMax attributes.
This attribute must be used in conjunction with the DCPowerDivisor attribute.
0x0000 is an invalid value for this attribute.`,
                id: 0x0204,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            DcVoltage: makeType<ZigBee.IMeasurementAndSensing.IArgDcVoltage, ZigBee.IMeasurementAndSensing.IArgDcVoltagePayload>(base.s16, ()=>({
                name: `DC Voltage`,
                description: `represents the most recent DC voltage reading in Volts (V). If the voltage cannot be
measured, a value of 0x8000 is returned.`,
                id: 0x0100,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Volts,
                
            })),
            DcVoltageDivisor: makeType<ZigBee.IMeasurementAndSensing.IArgDcVoltageDivisor, ZigBee.IMeasurementAndSensing.IArgDcVoltageDivisorPayload>(base.u16, ()=>({
                name: `DC Voltage Divisor`,
                description: `provides a value to be divided against the DCVoltage, DCVoltageMin, and DCVoltageMax attributes.
This attribute must be used in conjunction with the DCVoltageMultiplier attribute.
0x0000 is an invalid value for this attribute.`,
                id: 0x0201,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            DcVoltageMax: makeType<ZigBee.IMeasurementAndSensing.IArgDcVoltageMax, ZigBee.IMeasurementAndSensing.IArgDcVoltageMaxPayload>(base.s16, ()=>({
                name: `DC Voltage Max`,
                description: `represents the highest DC voltage value measured in Volts (V). After resetting,
this attribute will return a value of 0x8000 until a measurement is made.`,
                id: 0x0102,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Volts,
                
            })),
            DcVoltageMin: makeType<ZigBee.IMeasurementAndSensing.IArgDcVoltageMin, ZigBee.IMeasurementAndSensing.IArgDcVoltageMinPayload>(base.s16, ()=>({
                name: `DC Voltage Min`,
                description: `represents the lowest DC voltage value measured in Volts (V). After resetting,
this attribute will return a value of 0x8000 until a measurement is made.`,
                id: 0x0101,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Volts,
                
            })),
            DcVoltageMultiplier: makeType<ZigBee.IMeasurementAndSensing.IArgDcVoltageMultiplier, ZigBee.IMeasurementAndSensing.IArgDcVoltageMultiplierPayload>(base.u16, ()=>({
                name: `DC Voltage Multiplier`,
                description: `provides a value to be multiplied against the DCVoltage, DCVoltageMin, and DCVoltageMax attributes.
This attribute must be used in conjunction with the DCVoltageDivisor attribute.
0x0000 is an invalid value for this attribute`,
                id: 0x0200,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            DcVoltageOverload: makeType<ZigBee.IMeasurementAndSensing.IArgDcVoltageOverload, ZigBee.IMeasurementAndSensing.IArgDcVoltageOverloadPayload>(base.s16, ()=>({
                name: `DC Voltage Overload`,
                description: `Specifies the alarm threshold, set by the manufacturer, for the maximum output voltage supported by device.
The value is multiplied and divided by the DCVoltageMultiplier the DCVoltageDivisor respectively.`,
                id: 0x0701,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Volts,
                
            })),
            ElectricalMeasurementType: makeType<ZigBee.IMeasurementAndSensing.IArgElectricalMeasurementType, ZigBee.IMeasurementAndSensing.IArgElectricalMeasurementTypePayload>(base.bmp32, ()=>({
                name: `Electrical Measurement Type`,
                description: ``,
                id: 0x0000,
                report: false,
                read: false,
                write: false,
                require: false,
                bits: { 
                0: `Active measurement (AC)`, 
                1: `Reactive measurement (AC)`, 
                2: `Apparent measurement (AC)`, 
                3: `Phase A measurement`, 
                4: `Phase B measurement`, 
                5: `Phase C measurement`, 
                6: `DC measurement`, 
                7: `Harmonics measurement`, 
                8: `Power quality measurement`,  },
                
            })),
            HarmonicCurrentMultiplier: makeType<ZigBee.IMeasurementAndSensing.IArgHarmonicCurrentMultiplier, ZigBee.IMeasurementAndSensing.IArgHarmonicCurrentMultiplierPayload>(base.s8, ()=>({
                name: `Harmonic Current Multiplier`,
                description: `Represents the unit value for the MeasuredNthHarmonicCurrent attribute in the format
MeasuredNthHarmonicCurrent * 10 ^ HarmonicCurrentMultiplier amperes.`,
                id: 0x0404,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            IlluminanceLightSensorType: makeType<ZigBee.IMeasurementAndSensing.IArgIlluminanceLightSensorType, ZigBee.IMeasurementAndSensing.IArgIlluminanceLightSensorTypePayload>(base.enum8, ()=>({
                name: `Illuminance Light Sensor Type`,
                description: `specifies the electronic type of the light sensor. The range 0x40-0xfe is reserved for
manufacturer specific light sensor types`,
                id: 0x0001,
                report: false,
                read: true,
                write: false,
                require: false,
                values: { 
                0x00: `Photodiode`, 
                0x01: `CMOS`, 
                0xff: `Unknown`,  },
                
            })),
            IlluminanceTargetLevel: makeType<ZigBee.IMeasurementAndSensing.IArgIlluminanceTargetLevel, ZigBee.IMeasurementAndSensing.IArgIlluminanceTargetLevelPayload>(base.u16, ()=>({
                name: `Illuminance Target Level`,
                description: `specifies the target Illuminance level. This target level is taken as the
centre of a 'dead band', which must be sufficient in width, with hysteresis
bands at both top and bottom, to provide reliable notifications without
'chatter'. IlluminanceTargetLevel represents Illuminance in Lux (symbol lx) as
follows: IlluminanceTargetLevel = 10,000 x log 10 Illuminance`,
                id: 0x0010,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Luxes,
                
            })),
            Intervals: makeType<ZigBee.IMeasurementAndSensing.IArgIntervals, ZigBee.IMeasurementAndSensing.IArgIntervalsPayload>(base.array, ()=>({
                name: `Intervals`,
                description: `is a series of interval data captured using the period specified by the ProfileIntervalPeriod field. The
content of the interval data depend of the type of information requested using the AttributeID field in the Get
Measurement Profile Command. Data is organized in a reverse chronological order, the oldest intervals are
transmitted first and the newest interval is transmitted last. Invalid intervals should be marked as 0xFFFF.
For scaling and data type use the respective attribute set as defined above in attribute sets.`,
                arrayType: base.u16,
                
            })),
            LevelStatus: makeType<ZigBee.IMeasurementAndSensing.IArgLevelStatus, ZigBee.IMeasurementAndSensing.IArgLevelStatusPayload>(base.enum8, ()=>({
                name: `Level Status`,
                description: `indicates whether the measured Flow is above, below, or within a band
around FlowTargetLevel`,
                id: 0x0000,
                report: true,
                read: true,
                write: false,
                require: false,
                values: { 
                0x00: `Flow on target`, 
                0x01: `Flow below target`, 
                0x03: `Flow above target`,  },
                
            })),
            LightSensorType: makeType<ZigBee.IMeasurementAndSensing.IArgLightSensorType, ZigBee.IMeasurementAndSensing.IArgLightSensorTypePayload>(base.enum8, ()=>({
                name: `Light Sensor Type`,
                description: `specifies the electronic type of the light sensor. The range 0x40-0xfe is reserved for
manufacturer specific light sensor types`,
                id: 0x0004,
                report: false,
                read: true,
                write: false,
                require: false,
                values: { 
                0x00: `Photodiode`, 
                0x01: `CMOS`, 
                0xff: `Unknown`,  },
                
            })),
            LineCurrent: makeType<ZigBee.IMeasurementAndSensing.IArgLineCurrent, ZigBee.IMeasurementAndSensing.IArgLineCurrentPayload>(base.u16, ()=>({
                name: `Line Current`,
                description: `Represents the single phase or Phase A, AC line current (Square root of active and reactive current) value at
the moment in time the attribute is read. If it cannot be measured, a value of 0x8000 is returned.`,
                id: 0x0501,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Amperes,
                
            })),
            LineCurrentPhB: makeType<ZigBee.IMeasurementAndSensing.IArgLineCurrentPhB, ZigBee.IMeasurementAndSensing.IArgLineCurrentPhBPayload>(base.u16, ()=>({
                name: `Line Current Ph B`,
                description: `represents the Phase B, AC line current (Square root sum of active and reactive currents) value at the moment
in time the attribute is read.
If the instantaneous current cannot be measured, a value of 0x8000 is returned.`,
                id: 0x0901,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Amperes,
                
            })),
            LineCurrentPhC: makeType<ZigBee.IMeasurementAndSensing.IArgLineCurrentPhC, ZigBee.IMeasurementAndSensing.IArgLineCurrentPhCPayload>(base.u16, ()=>({
                name: `Line Current Ph C`,
                description: ``,
                id: 0x0A01,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            ListOfAttributes: makeType<ZigBee.IMeasurementAndSensing.IArgListOfAttributes, ZigBee.IMeasurementAndSensing.IArgListOfAttributesPayload>(base.set, ()=>({
                name: `List Of Attributes`,
                description: `is the list of attributes being profiled`,
                arrayType: base.u16,
                
            })),
            MaxMeasuredConcentration: makeType<ZigBee.IMeasurementAndSensing.IArgMaxMeasuredConcentration, ZigBee.IMeasurementAndSensing.IArgMaxMeasuredConcentrationPayload>(base.float, ()=>({
                name: `Max Measured Concentration`,
                description: `indicates the maximum of MeasuredConcentration that is capable of being measured.
A value of NaN indicates that the maximum value is not defined`,
                id: 0x0002,
                report: true,
                read: true,
                write: false,
                require: false,
                unit: units.Concentration,
                
            })),
            MaxMeasuredFlow: makeType<ZigBee.IMeasurementAndSensing.IArgMaxMeasuredFlow, ZigBee.IMeasurementAndSensing.IArgMaxMeasuredFlowPayload>(base.u16, ()=>({
                name: `Max Measured Flow`,
                description: `indicates the maximum value of MeasuredFlow that can be measured. A
value of 0xffff indicates that this attribute is not defined`,
                id: 0x0002,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.CubicMetersPerHour,
                scale: 10,
                
            })),
            MaxMeasuredIlluminance: makeType<ZigBee.IMeasurementAndSensing.IArgMaxMeasuredIlluminance, ZigBee.IMeasurementAndSensing.IArgMaxMeasuredIlluminancePayload>(base.u16, ()=>({
                name: `Max Measured Illuminance`,
                description: `indicates the maximum value of MeasuredIlluminance that can be measured. A
value of 0xffff indicates that this attribute is not defined`,
                id: 0x0002,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.Luxes,
                
            })),
            MaxMeasuredPm25: makeType<ZigBee.IMeasurementAndSensing.IArgMaxMeasuredPm25, ZigBee.IMeasurementAndSensing.IArgMaxMeasuredPm25Payload>(base.float, ()=>({
                name: `Max Measured PM2.5`,
                description: ``,
                id: 0x0002,
                report: true,
                read: true,
                write: false,
                require: false,
                unit: units.MicrogramPerCubicMeter,
                
            })),
            MaxMeasuredPressure: makeType<ZigBee.IMeasurementAndSensing.IArgMaxMeasuredPressure, ZigBee.IMeasurementAndSensing.IArgMaxMeasuredPressurePayload>(base.s16, ()=>({
                name: `Max Measured Pressure`,
                description: `indicates the maximum value of MeasuredPressure that can be measured. A
value of 0x8000 indicates that this attribute is not defined`,
                id: 0x0002,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.Kilopascals,
                scale: 10,
                
            })),
            MaxMeasuredRelativeHumidity: makeType<ZigBee.IMeasurementAndSensing.IArgMaxMeasuredRelativeHumidity, ZigBee.IMeasurementAndSensing.IArgMaxMeasuredRelativeHumidityPayload>(base.u16, ()=>({
                name: `Max Measured Relative Humidity`,
                description: `indicates the maximum value of MeasuredRH that can be measured. A
value of 0xffff indicates that this attribute is not defined`,
                id: 0x0002,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.PercentRelativeHumidity,
                scale: 100,
                
            })),
            MaxMeasuredTemperature: makeType<ZigBee.IMeasurementAndSensing.IArgMaxMeasuredTemperature, ZigBee.IMeasurementAndSensing.IArgMaxMeasuredTemperaturePayload>(base.s16, ()=>({
                name: `Max Measured Temperature`,
                description: `indicates the maximum value of MeasuredTemperature that can be measured. A
value of 0x8000 indicates that this attribute is not defined`,
                id: 0x0002,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.DegreesCelsius,
                scale: 100,
                
            })),
            MaxNumberOfIntervals: makeType<ZigBee.IMeasurementAndSensing.IArgMaxNumberOfIntervals, ZigBee.IMeasurementAndSensing.IArgMaxNumberOfIntervalsPayload>(base.u8, ()=>({
                name: `Max Number Of Intervals`,
                description: `represents the maximum number of intervals the device is capable of returning in one command.
It is required MaxNumberofIntervals fit within the default Fragmentation ASDU size of 128 bytes,
or an optionally agreed upon larger Fragmentation ASDU size supported by both devices.`,
                
            })),
            Measured11ThHarmonicCurrent: makeType<ZigBee.IMeasurementAndSensing.IArgMeasured11ThHarmonicCurrent, ZigBee.IMeasurementAndSensing.IArgMeasured11ThHarmonicCurrentPayload>(base.s16, ()=>({
                name: `Measured 11th Harmonic Current`,
                description: `represent the most recent 11th harmonic current reading in an AC frequency.`,
                id: 0x030C,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Amperes,
                
            })),
            Measured1StHarmonicCurrent: makeType<ZigBee.IMeasurementAndSensing.IArgMeasured1StHarmonicCurrent, ZigBee.IMeasurementAndSensing.IArgMeasured1StHarmonicCurrentPayload>(base.s16, ()=>({
                name: `Measured 1st Harmonic Current`,
                description: `represent the most recent 1st harmonic current reading in an AC frequency.
The unit for this measurement is 10 ^ 1stHarmonicCurrentMultiplier amperes.
If 1stHarmonicCurrentMultiplier is not implemented the unit is in amperes.
If the 1st harmonic current cannot be measured a value of 0x8000 is returned.
A positive value indicates the measured 1st harmonic current is positive, and a
negative value indicates that the measured 1st harmonic current is negative.`,
                id: 0x0307,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Amperes,
                
            })),
            Measured3RdHarmonicCurrent: makeType<ZigBee.IMeasurementAndSensing.IArgMeasured3RdHarmonicCurrent, ZigBee.IMeasurementAndSensing.IArgMeasured3RdHarmonicCurrentPayload>(base.s16, ()=>({
                name: `Measured 3rd Harmonic Current`,
                description: `represent the most recent 3rd harmonic current reading in an AC frequency.`,
                id: 0x0308,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Amperes,
                
            })),
            Measured5ThHarmonicCurrent: makeType<ZigBee.IMeasurementAndSensing.IArgMeasured5ThHarmonicCurrent, ZigBee.IMeasurementAndSensing.IArgMeasured5ThHarmonicCurrentPayload>(base.s16, ()=>({
                name: `Measured 5th Harmonic Current`,
                description: `represent the most recent 5th harmonic current reading in an AC frequency.`,
                id: 0x0309,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Amperes,
                
            })),
            Measured7ThHarmonicCurrent: makeType<ZigBee.IMeasurementAndSensing.IArgMeasured7ThHarmonicCurrent, ZigBee.IMeasurementAndSensing.IArgMeasured7ThHarmonicCurrentPayload>(base.s16, ()=>({
                name: `Measured 7th Harmonic Current`,
                description: `represent the most recent 7th harmonic current reading in an AC frequency.`,
                id: 0x030A,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Amperes,
                
            })),
            Measured9ThHarmonicCurrent: makeType<ZigBee.IMeasurementAndSensing.IArgMeasured9ThHarmonicCurrent, ZigBee.IMeasurementAndSensing.IArgMeasured9ThHarmonicCurrentPayload>(base.s16, ()=>({
                name: `Measured 9th Harmonic Current`,
                description: `represent the most recent 9th harmonic current reading in an AC frequency.`,
                id: 0x030B,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Amperes,
                
            })),
            MeasuredConcentration: makeType<ZigBee.IMeasurementAndSensing.IArgMeasuredConcentration, ZigBee.IMeasurementAndSensing.IArgMeasuredConcentrationPayload>(base.float, ()=>({
                name: `Measured Concentration`,
                description: `represents the concentration as a fraction of 1 (one).
A value of NaN indicates that the concentration measurement is unknown or outside the valid range.
MinMeasuredConcentration and MaxMeasuredConcentration define the valid range for MeasuredConcentration.
MeasuredConcentration is updated continuously as new measurements are made`,
                id: 0x0000,
                report: true,
                read: true,
                write: false,
                require: false,
                unit: units.Concentration,
                
            })),
            MeasuredFlow: makeType<ZigBee.IMeasurementAndSensing.IArgMeasuredFlow, ZigBee.IMeasurementAndSensing.IArgMeasuredFlowPayload>(base.u16, ()=>({
                name: `Measured Flow`,
                description: `represents the flow in m^3/h`,
                id: 0x0000,
                report: true,
                read: true,
                write: false,
                require: false,
                unit: units.CubicMetersPerHour,
                scale: 10,
                
            })),
            MeasuredIlluminance: makeType<ZigBee.IMeasurementAndSensing.IArgMeasuredIlluminance, ZigBee.IMeasurementAndSensing.IArgMeasuredIlluminancePayload>(base.u16, ()=>({
                name: `Measured Illuminance`,
                description: `represents the Illuminance in Lux (symbol lx) as follows
MeasuredValue = 10,000 x log10 Flow + 1 where 1 lx
<= Flow <=3.576 Mlx, corresponding to a MeasuredValue
in the range 1 to 0xfffe`,
                id: 0x0000,
                report: true,
                read: true,
                write: false,
                require: false,
                unit: units.Luxes,
                
            })),
            MeasuredPhase11ThHarmonicCurrent: makeType<ZigBee.IMeasurementAndSensing.IArgMeasuredPhase11ThHarmonicCurrent, ZigBee.IMeasurementAndSensing.IArgMeasuredPhase11ThHarmonicCurrentPayload>(base.s16, ()=>({
                name: `Measured Phase 11th Harmonic Current`,
                description: `represent the most recent phase of the 11th harmonic current reading in an AC frequency.`,
                id: 0x0312,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.DegreesPhase,
                
            })),
            MeasuredPhase1StHarmonicCurrent: makeType<ZigBee.IMeasurementAndSensing.IArgMeasuredPhase1StHarmonicCurrent, ZigBee.IMeasurementAndSensing.IArgMeasuredPhase1StHarmonicCurrentPayload>(base.s16, ()=>({
                name: `Measured Phase 1st Harmonic Current`,
                description: `represent the most recent phase of the 1st harmonic current reading in an AC frequency.
The unit for this measurement is 10 ^ Phase1StHarmonicCurrentMultiplier degree.
If Phase1StHarmonicCurrentMultiplier is not implemented the unit is in degree.
If the phase of cannot be measured a value of 0x8000 is returned.
A positive value indicates the measured phase is prehurry,
and a negative value indicates that the measured phase is lagging.`,
                id: 0x030D,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.DegreesPhase,
                
            })),
            MeasuredPhase3RdHarmonicCurrent: makeType<ZigBee.IMeasurementAndSensing.IArgMeasuredPhase3RdHarmonicCurrent, ZigBee.IMeasurementAndSensing.IArgMeasuredPhase3RdHarmonicCurrentPayload>(base.s16, ()=>({
                name: `Measured Phase 3rd Harmonic Current`,
                description: `represent the most recent phase of the 3rd harmonic current reading in an AC frequency.`,
                id: 0x030E,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.DegreesPhase,
                
            })),
            MeasuredPhase5ThHarmonicCurrent: makeType<ZigBee.IMeasurementAndSensing.IArgMeasuredPhase5ThHarmonicCurrent, ZigBee.IMeasurementAndSensing.IArgMeasuredPhase5ThHarmonicCurrentPayload>(base.s16, ()=>({
                name: `Measured Phase 5th Harmonic Current`,
                description: `represent the most recent phase of the 5th harmonic current reading in an AC frequency.`,
                id: 0x030F,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.DegreesPhase,
                
            })),
            MeasuredPhase7ThHarmonicCurrent: makeType<ZigBee.IMeasurementAndSensing.IArgMeasuredPhase7ThHarmonicCurrent, ZigBee.IMeasurementAndSensing.IArgMeasuredPhase7ThHarmonicCurrentPayload>(base.s16, ()=>({
                name: `Measured Phase 7th Harmonic Current`,
                description: `represent the most recent phase of the 7th harmonic current reading in an AC frequency.`,
                id: 0x0310,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.DegreesPhase,
                
            })),
            MeasuredPhase9ThHarmonicCurrent: makeType<ZigBee.IMeasurementAndSensing.IArgMeasuredPhase9ThHarmonicCurrent, ZigBee.IMeasurementAndSensing.IArgMeasuredPhase9ThHarmonicCurrentPayload>(base.s16, ()=>({
                name: `Measured Phase 9th Harmonic Current`,
                description: `represent the most recent phase of the 9th harmonic current reading in an AC frequency.`,
                id: 0x0311,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.DegreesPhase,
                
            })),
            MeasuredPm25: makeType<ZigBee.IMeasurementAndSensing.IArgMeasuredPm25, ZigBee.IMeasurementAndSensing.IArgMeasuredPm25Payload>(base.float, ()=>({
                name: `Measured PM2.5`,
                description: ``,
                id: 0x0000,
                report: true,
                read: true,
                write: false,
                require: false,
                unit: units.MicrogramPerCubicMeter,
                
            })),
            MeasuredPressure: makeType<ZigBee.IMeasurementAndSensing.IArgMeasuredPressure, ZigBee.IMeasurementAndSensing.IArgMeasuredPressurePayload>(base.s16, ()=>({
                name: `Measured Pressure`,
                description: `represents the temperature in degrees DegreesCelsius`,
                id: 0x0000,
                report: true,
                read: true,
                write: false,
                require: false,
                unit: units.Kilopascals,
                scale: 10,
                
            })),
            MeasuredRelativeHumidity: makeType<ZigBee.IMeasurementAndSensing.IArgMeasuredRelativeHumidity, ZigBee.IMeasurementAndSensing.IArgMeasuredRelativeHumidityPayload>(base.u16, ()=>({
                name: `Measured Relative Humidity`,
                description: `represents the relative humidity in percent`,
                id: 0x0000,
                report: true,
                read: true,
                write: false,
                require: false,
                unit: units.PercentRelativeHumidity,
                scale: 100,
                
            })),
            MeasuredTemperature: makeType<ZigBee.IMeasurementAndSensing.IArgMeasuredTemperature, ZigBee.IMeasurementAndSensing.IArgMeasuredTemperaturePayload>(base.s16, ()=>({
                name: `Measured Temperature`,
                description: `represents the temperature in degrees DegreesCelsius`,
                id: 0x0000,
                report: true,
                read: true,
                write: false,
                require: false,
                unit: units.DegreesCelsius,
                scale: 100,
                
            })),
            MeasurementResponseStatus: makeType<ZigBee.IMeasurementAndSensing.IArgMeasurementResponseStatus, ZigBee.IMeasurementAndSensing.IArgMeasurementResponseStatusPayload>(base.enum8, ()=>({
                name: `Measurement Response Status`,
                description: ``,
                values: { 
                0x00: `Success`, 
                0x01: `Attribute Profile not supported`, 
                0x02: `Invalid Start Time`, 
                0x03: `More intervals requested than can be returned`, 
                0x04: `No intervals available for the requested time`,  },
                
            })),
            MinMeasuredConcentration: makeType<ZigBee.IMeasurementAndSensing.IArgMinMeasuredConcentration, ZigBee.IMeasurementAndSensing.IArgMinMeasuredConcentrationPayload>(base.float, ()=>({
                name: `Min Measured Concentration`,
                description: `indicates the minimum value of MeasuredConcentration that is capable of being measured.
A value of NaN indicates that the minimum value is not defined`,
                id: 0x0001,
                report: true,
                read: true,
                write: false,
                require: false,
                unit: units.Concentration,
                
            })),
            MinMeasuredFlow: makeType<ZigBee.IMeasurementAndSensing.IArgMinMeasuredFlow, ZigBee.IMeasurementAndSensing.IArgMinMeasuredFlowPayload>(base.u16, ()=>({
                name: `Min Measured Flow`,
                description: `indicates the minimum value of MeasuredFlow that can be measured. A
value of 0xffff indicates that this attribute is not defined`,
                id: 0x0001,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.CubicMetersPerHour,
                scale: 10,
                
            })),
            MinMeasuredIlluminance: makeType<ZigBee.IMeasurementAndSensing.IArgMinMeasuredIlluminance, ZigBee.IMeasurementAndSensing.IArgMinMeasuredIlluminancePayload>(base.u16, ()=>({
                name: `Min Measured Illuminance`,
                description: `indicates the minimum value of MeasuredIlluminance that can be measured. A
value of 0xffff indicates that this attribute is not defined`,
                id: 0x0001,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.Luxes,
                
            })),
            MinMeasuredPm25: makeType<ZigBee.IMeasurementAndSensing.IArgMinMeasuredPm25, ZigBee.IMeasurementAndSensing.IArgMinMeasuredPm25Payload>(base.float, ()=>({
                name: `Min Measured PM2.5`,
                description: ``,
                id: 0x0001,
                report: true,
                read: true,
                write: false,
                require: false,
                unit: units.MicrogramPerCubicMeter,
                
            })),
            MinMeasuredPressure: makeType<ZigBee.IMeasurementAndSensing.IArgMinMeasuredPressure, ZigBee.IMeasurementAndSensing.IArgMinMeasuredPressurePayload>(base.s16, ()=>({
                name: `Min Measured Pressure`,
                description: `indicates the minimum value of MeasuredPressure that can be measured. A
value of 0x8000 indicates that this attribute is not defined`,
                id: 0x0001,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.Kilopascals,
                scale: 10,
                
            })),
            MinMeasuredRelativeHumidity: makeType<ZigBee.IMeasurementAndSensing.IArgMinMeasuredRelativeHumidity, ZigBee.IMeasurementAndSensing.IArgMinMeasuredRelativeHumidityPayload>(base.u16, ()=>({
                name: `Min Measured Relative Humidity`,
                description: `indicates the minimum value of MeasuredRH that can be measured. A
value of 0xffff indicates that this attribute is not defined`,
                id: 0x0001,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.PercentRelativeHumidity,
                scale: 100,
                
            })),
            MinMeasuredTemperature: makeType<ZigBee.IMeasurementAndSensing.IArgMinMeasuredTemperature, ZigBee.IMeasurementAndSensing.IArgMinMeasuredTemperaturePayload>(base.s16, ()=>({
                name: `Min Measured Temperature`,
                description: `indicates the minimum value of MeasuredTemperature that can be measured. A
value of 0x8000 indicates that this attribute is not defined`,
                id: 0x0001,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.DegreesCelsius,
                scale: 100,
                
            })),
            NeutralCurrent: makeType<ZigBee.IMeasurementAndSensing.IArgNeutralCurrent, ZigBee.IMeasurementAndSensing.IArgNeutralCurrentPayload>(base.u16, ()=>({
                name: `Neutral Current`,
                description: `represents the AC neutral (Line-Out) current value at the moment in time the
attribute is read. If the instantaneous current cannot be measured, a value of 0xFFFF is returned`,
                id: 0x0303,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Amperes,
                
            })),
            NumberOfIntervals: makeType<ZigBee.IMeasurementAndSensing.IArgNumberOfIntervals, ZigBee.IMeasurementAndSensing.IArgNumberOfIntervalsPayload>(base.u8, ()=>({
                name: `Number Of Intervals`,
                description: `is the number of intervals requested or returned`,
                
            })),
            Occupancy: makeType<ZigBee.IMeasurementAndSensing.IArgOccupancy, ZigBee.IMeasurementAndSensing.IArgOccupancyPayload>(base.bmp8, ()=>({
                name: `Occupancy`,
                description: ``,
                id: 0x0000,
                report: true,
                read: true,
                write: false,
                require: false,
                bits: { 
                0: `Occupied`,  },
                
            })),
            OccupancyType: makeType<ZigBee.IMeasurementAndSensing.IArgOccupancyType, ZigBee.IMeasurementAndSensing.IArgOccupancyTypePayload>(base.enum8, ()=>({
                name: `Occupancy Type`,
                description: ``,
                id: 0x0001,
                report: false,
                read: true,
                write: false,
                require: false,
                values: { 
                0x00: `PIR`, 
                0x01: `Ultrasonic`, 
                0x03: `PIR and ultrasonic`,  },
                
            })),
            PhaseHarmonicCurrentMultiplier: makeType<ZigBee.IMeasurementAndSensing.IArgPhaseHarmonicCurrentMultiplier, ZigBee.IMeasurementAndSensing.IArgPhaseHarmonicCurrentMultiplierPayload>(base.s8, ()=>({
                name: `Phase Harmonic Current Multiplier`,
                description: `Represents the unit value for the MeasuredPhaseNthHarmonicCurrent attribute in the format
MeasuredPhaseNthHarmonicCurrent * 10 ^ PhaseHarmonicCurrentMultiplier degrees.`,
                id: 0x0405,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            PirOccupiedToUnoccupiedDelay: makeType<ZigBee.IMeasurementAndSensing.IArgPirOccupiedToUnoccupiedDelay, ZigBee.IMeasurementAndSensing.IArgPirOccupiedToUnoccupiedDelayPayload>(base.u16, ()=>({
                name: `PIR Occupied To Unoccupied Delay`,
                description: `is 16 bits in length and specifies the time delay, in seconds, before the PIR sensor
changes to its unoccupied state after the last detection of movement in the sensed area`,
                id: 0x0010,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            PirUnoccupiedToOccupiedDelay: makeType<ZigBee.IMeasurementAndSensing.IArgPirUnoccupiedToOccupiedDelay, ZigBee.IMeasurementAndSensing.IArgPirUnoccupiedToOccupiedDelayPayload>(base.u16, ()=>({
                name: `PIR Unoccupied To Occupied Delay`,
                description: `is 16 bits in length and specifies the time delay, in seconds, before the PIR sensor changes
to its occupied state after the detection of movement in the sensed area. This attribute is
mandatory if the PIRUnoccupiedToOccupiedThreshold attribute is implemented`,
                id: 0x0011,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            PirUnoccupiedToOccupiedThreshold: makeType<ZigBee.IMeasurementAndSensing.IArgPirUnoccupiedToOccupiedThreshold, ZigBee.IMeasurementAndSensing.IArgPirUnoccupiedToOccupiedThresholdPayload>(base.u8, ()=>({
                name: `PIR Unoccupied To Occupied Threshold`,
                description: `is 8 bits in length and specifies the number of movement detection events that must occur in the period PIRUnoccupiedToOccupiedDelay, before the PIR sensor changes to its occupied state. This attribute is mandatory if the PIRUnoccupiedToOccupiedDelay attribute is implemented`,
                id: 0x0012,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            Pm25Tolerance: makeType<ZigBee.IMeasurementAndSensing.IArgPm25Tolerance, ZigBee.IMeasurementAndSensing.IArgPm25TolerancePayload>(base.float, ()=>({
                name: `PM2.5 Tolerance`,
                description: ``,
                id: 0x0003,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.MicrogramPerCubicMeter,
                
            })),
            PowerDivisor: makeType<ZigBee.IMeasurementAndSensing.IArgPowerDivisor, ZigBee.IMeasurementAndSensing.IArgPowerDivisorPayload>(base.u32, ()=>({
                name: `Power Divisor`,
                description: `Provides a value to divide against the results of applying the Multiplier attribute against a raw or
uncompensated sensor count of power being measured by the metering device. If present, this attribute must
be applied against all demand/power values to derive the delivered and received values expressed in the
specified units. This attribute must be used in conjunction with the PowerMultiplier attribute.`,
                id: 0x0403,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            PowerFactor: makeType<ZigBee.IMeasurementAndSensing.IArgPowerFactor, ZigBee.IMeasurementAndSensing.IArgPowerFactorPayload>(base.s8, ()=>({
                name: `Power Factor`,
                description: `contains the single phase or PhaseA, Power Factor ratio in 1/100ths`,
                id: 0x0510,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            PowerFactorPhB: makeType<ZigBee.IMeasurementAndSensing.IArgPowerFactorPhB, ZigBee.IMeasurementAndSensing.IArgPowerFactorPhBPayload>(base.s8, ()=>({
                name: `Power Factor Ph B`,
                description: ``,
                id: 0x0910,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            PowerFactorPhC: makeType<ZigBee.IMeasurementAndSensing.IArgPowerFactorPhC, ZigBee.IMeasurementAndSensing.IArgPowerFactorPhCPayload>(base.s8, ()=>({
                name: `Power Factor Ph C`,
                description: ``,
                id: 0x0A10,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            PowerMultiplier: makeType<ZigBee.IMeasurementAndSensing.IArgPowerMultiplier, ZigBee.IMeasurementAndSensing.IArgPowerMultiplierPayload>(base.u32, ()=>({
                name: `Power Multiplier`,
                description: `Provides a value to be multiplied against a raw or uncompensated sensor count of power being measured by
the metering device. If present, this attribute must be applied against all power/demand values to derive the
delivered and received values expressed in the specified units. This attribute must be used in conjunction with
the PowerDivisor attribute.`,
                id: 0x0402,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            ProfileCount: makeType<ZigBee.IMeasurementAndSensing.IArgProfileCount, ZigBee.IMeasurementAndSensing.IArgProfileCountPayload>(base.u8, ()=>({
                name: `Profile Count`,
                description: `is the total number of supported profiles`,
                
            })),
            ProfileIntervalPeriod: makeType<ZigBee.IMeasurementAndSensing.IArgProfileIntervalPeriod, ZigBee.IMeasurementAndSensing.IArgProfileIntervalPeriodPayload>(base.enum8, ()=>({
                name: `Profile Interval Period`,
                description: `represents the interval or time frame used to capture parameter for profiling purposes`,
                values: { 
                0x00: `Daily`, 
                0x01: `60 minutes`, 
                0x02: `30 minutes`, 
                0x03: `15 minutes`, 
                0x04: `10 minutes`, 
                0x05: `7.5 minutes`, 
                0x06: `5 minutes`, 
                0x07: `2.5 minutes`,  },
                
            })),
            ReactiveCurrent: makeType<ZigBee.IMeasurementAndSensing.IArgReactiveCurrent, ZigBee.IMeasurementAndSensing.IArgReactiveCurrentPayload>(base.s16, ()=>({
                name: `Reactive Current`,
                description: `Represents the single phase or Phase A, AC reactive current value at the moment in time the attribute is read,
in Amps (A). Positive values indicate power delivered to the premises where negative values indicate power
received from the premises.`,
                id: 0x0503,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Amperes,
                
            })),
            ReactiveCurrentPhB: makeType<ZigBee.IMeasurementAndSensing.IArgReactiveCurrentPhB, ZigBee.IMeasurementAndSensing.IArgReactiveCurrentPhBPayload>(base.s16, ()=>({
                name: `Reactive Current Ph B`,
                description: `represents the Phase B, AC reactive current value at the moment in time the attribute is read.
Positive values indicate power delivered to the premises where negative values indicate power received from
the premises.`,
                id: 0x0903,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Amperes,
                
            })),
            ReactiveCurrentPhC: makeType<ZigBee.IMeasurementAndSensing.IArgReactiveCurrentPhC, ZigBee.IMeasurementAndSensing.IArgReactiveCurrentPhCPayload>(base.s16, ()=>({
                name: `Reactive Current Ph C`,
                description: ``,
                id: 0x0A03,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            ReactivePower: makeType<ZigBee.IMeasurementAndSensing.IArgReactivePower, ZigBee.IMeasurementAndSensing.IArgReactivePowerPayload>(base.s16, ()=>({
                name: `Reactive Power`,
                description: `Represents the single phase or Phase A, current demand of reactive power delivered or received at the
premises, in VAr. Positive values indicate power delivered to the premises where negative values indicate
power received from the premises.`,
                id: 0x050E,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.VoltAmperesReactive,
                
            })),
            ReactivePowerPhB: makeType<ZigBee.IMeasurementAndSensing.IArgReactivePowerPhB, ZigBee.IMeasurementAndSensing.IArgReactivePowerPhBPayload>(base.s16, ()=>({
                name: `Reactive Power Ph B`,
                description: ``,
                id: 0x090E,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            ReactivePowerPhC: makeType<ZigBee.IMeasurementAndSensing.IArgReactivePowerPhC, ZigBee.IMeasurementAndSensing.IArgReactivePowerPhCPayload>(base.s16, ()=>({
                name: `Reactive Power Ph C`,
                description: ``,
                id: 0x0A0E,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            RmsCurrent: makeType<ZigBee.IMeasurementAndSensing.IArgRmsCurrent, ZigBee.IMeasurementAndSensing.IArgRmsCurrentPayload>(base.u16, ()=>({
                name: `RMS Current`,
                description: `Represents the most recent RMS current reading.`,
                id: 0x0508,
                report: true,
                read: false,
                write: false,
                require: false,
                unit: units.Amperes,
                
            })),
            RmsCurrentMax: makeType<ZigBee.IMeasurementAndSensing.IArgRmsCurrentMax, ZigBee.IMeasurementAndSensing.IArgRmsCurrentMaxPayload>(base.u16, ()=>({
                name: `RMS Current Max`,
                description: `Represents the highest RMS current value.`,
                id: 0x050A,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Amperes,
                
            })),
            RmsCurrentMaxPhB: makeType<ZigBee.IMeasurementAndSensing.IArgRmsCurrentMaxPhB, ZigBee.IMeasurementAndSensing.IArgRmsCurrentMaxPhBPayload>(base.u16, ()=>({
                name: `RMS Current Max Ph B`,
                description: ``,
                id: 0x090A,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            RmsCurrentMaxPhC: makeType<ZigBee.IMeasurementAndSensing.IArgRmsCurrentMaxPhC, ZigBee.IMeasurementAndSensing.IArgRmsCurrentMaxPhCPayload>(base.u16, ()=>({
                name: `RMS Current Max Ph C`,
                description: ``,
                id: 0x0A0A,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            RmsCurrentMin: makeType<ZigBee.IMeasurementAndSensing.IArgRmsCurrentMin, ZigBee.IMeasurementAndSensing.IArgRmsCurrentMinPayload>(base.u16, ()=>({
                name: `RMS Current Min`,
                description: `Represents the lowest RMS current value.`,
                id: 0x0509,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Amperes,
                
            })),
            RmsCurrentMinPhB: makeType<ZigBee.IMeasurementAndSensing.IArgRmsCurrentMinPhB, ZigBee.IMeasurementAndSensing.IArgRmsCurrentMinPhBPayload>(base.u16, ()=>({
                name: `RMS Current Min Ph B`,
                description: ``,
                id: 0x0909,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            RmsCurrentMinPhC: makeType<ZigBee.IMeasurementAndSensing.IArgRmsCurrentMinPhC, ZigBee.IMeasurementAndSensing.IArgRmsCurrentMinPhCPayload>(base.u16, ()=>({
                name: `RMS Current Min Ph C`,
                description: ``,
                id: 0x0A09,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            RmsCurrentPhB: makeType<ZigBee.IMeasurementAndSensing.IArgRmsCurrentPhB, ZigBee.IMeasurementAndSensing.IArgRmsCurrentPhBPayload>(base.u16, ()=>({
                name: `RMS Current Ph B`,
                description: ``,
                id: 0x0908,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            RmsCurrentPhC: makeType<ZigBee.IMeasurementAndSensing.IArgRmsCurrentPhC, ZigBee.IMeasurementAndSensing.IArgRmsCurrentPhCPayload>(base.u16, ()=>({
                name: `RMS Current Ph C`,
                description: ``,
                id: 0x0A08,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            RmsExtremeOverVoltage: makeType<ZigBee.IMeasurementAndSensing.IArgRmsExtremeOverVoltage, ZigBee.IMeasurementAndSensing.IArgRmsExtremeOverVoltagePayload>(base.s16, ()=>({
                name: `RMS extreme over-voltage`,
                description: `is the RMS voltage above which an extreme under voltage condition is reported.
The threshold shall be configurable within the specified operating range of the electricity meter.
The value is multiplied and divided by the ACVoltageMultiplier and ACVoltageDivisor, respectively.`,
                id: 0x0807,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Volts,
                
            })),
            RmsExtremeOverVoltagePeriod: makeType<ZigBee.IMeasurementAndSensing.IArgRmsExtremeOverVoltagePeriod, ZigBee.IMeasurementAndSensing.IArgRmsExtremeOverVoltagePeriodPayload>(base.u16, ()=>({
                name: `RMS Extreme Over Voltage Period`,
                description: `is the durations used to measure an extreme over voltage condition.`,
                id: 0x0514,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Seconds,
                
            })),
            RmsExtremeOverVoltagePeriodPhB: makeType<ZigBee.IMeasurementAndSensing.IArgRmsExtremeOverVoltagePeriodPhB, ZigBee.IMeasurementAndSensing.IArgRmsExtremeOverVoltagePeriodPhBPayload>(base.u16, ()=>({
                name: `RMS Extreme Over Voltage Period Ph B`,
                description: ``,
                id: 0x0914,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            RmsExtremeOverVoltagePeriodPhC: makeType<ZigBee.IMeasurementAndSensing.IArgRmsExtremeOverVoltagePeriodPhC, ZigBee.IMeasurementAndSensing.IArgRmsExtremeOverVoltagePeriodPhCPayload>(base.u16, ()=>({
                name: `RMS Extreme Over-voltage Period Ph C`,
                description: ``,
                id: 0x0A14,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            RmsExtremeUnderVoltage: makeType<ZigBee.IMeasurementAndSensing.IArgRmsExtremeUnderVoltage, ZigBee.IMeasurementAndSensing.IArgRmsExtremeUnderVoltagePayload>(base.s16, ()=>({
                name: `RMS extreme under-voltage`,
                description: `is the RMS voltage below which an extreme under voltage condition is reported. The threshold shall be
configurable within the specified operating range of the electricity meter. The value is multiplied and divided
by the ACVoltageMultiplier and ACVoltageDivisor, respectively.`,
                id: 0x0808,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Volts,
                
            })),
            RmsExtremeUnderVoltagePeriod: makeType<ZigBee.IMeasurementAndSensing.IArgRmsExtremeUnderVoltagePeriod, ZigBee.IMeasurementAndSensing.IArgRmsExtremeUnderVoltagePeriodPayload>(base.u16, ()=>({
                name: `RMS Extreme Under Voltage Period`,
                description: `is the duration used to measure an extreme under voltage condition.`,
                id: 0x0515,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Seconds,
                
            })),
            RmsExtremeUnderVoltagePeriodPhB: makeType<ZigBee.IMeasurementAndSensing.IArgRmsExtremeUnderVoltagePeriodPhB, ZigBee.IMeasurementAndSensing.IArgRmsExtremeUnderVoltagePeriodPhBPayload>(base.u16, ()=>({
                name: `RMS Extreme Under Voltage Period Ph B`,
                description: ``,
                id: 0x0915,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            RmsExtremeUnderVoltagePeriodPhC: makeType<ZigBee.IMeasurementAndSensing.IArgRmsExtremeUnderVoltagePeriodPhC, ZigBee.IMeasurementAndSensing.IArgRmsExtremeUnderVoltagePeriodPhCPayload>(base.u16, ()=>({
                name: `RMS Extreme Under-voltage Period Ph C`,
                description: ``,
                id: 0x0A15,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            RmsVoltage: makeType<ZigBee.IMeasurementAndSensing.IArgRmsVoltage, ZigBee.IMeasurementAndSensing.IArgRmsVoltagePayload>(base.u16, ()=>({
                name: `RMS Voltage`,
                description: `represents the most recent RMS voltage reading.
If the RMS voltage cannot be measured, a value of 0xFFFF is returned`,
                id: 0x0505,
                report: true,
                read: false,
                write: false,
                require: false,
                unit: units.Volts,
                
            })),
            RmsVoltageMax: makeType<ZigBee.IMeasurementAndSensing.IArgRmsVoltageMax, ZigBee.IMeasurementAndSensing.IArgRmsVoltageMaxPayload>(base.u16, ()=>({
                name: `RMS Voltage Max`,
                description: `Represents the highest RMS voltage value.
After resetting, this attribute will return a value of 0xFFFF until a measurement is made.`,
                id: 0x0507,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Volts,
                
            })),
            RmsVoltageMaxPhB: makeType<ZigBee.IMeasurementAndSensing.IArgRmsVoltageMaxPhB, ZigBee.IMeasurementAndSensing.IArgRmsVoltageMaxPhBPayload>(base.u16, ()=>({
                name: `RMS Voltage Max Ph B`,
                description: `represents the highest RMS voltage value measured. After resetting, this attribute will return a
value of 0xFFFF until a measurement is made.`,
                id: 0x0907,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Volts,
                
            })),
            RmsVoltageMaxPhC: makeType<ZigBee.IMeasurementAndSensing.IArgRmsVoltageMaxPhC, ZigBee.IMeasurementAndSensing.IArgRmsVoltageMaxPhCPayload>(base.u16, ()=>({
                name: `RMS Voltage Max Ph C`,
                description: ``,
                id: 0x0A07,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            RmsVoltageMin: makeType<ZigBee.IMeasurementAndSensing.IArgRmsVoltageMin, ZigBee.IMeasurementAndSensing.IArgRmsVoltageMinPayload>(base.u16, ()=>({
                name: `RMS Voltage Min`,
                description: `Represents the lowest RMS voltage value.
After resetting, this attribute will return a value of 0xFFFF until a measurement is made.`,
                id: 0x0506,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Volts,
                
            })),
            RmsVoltageMinPhB: makeType<ZigBee.IMeasurementAndSensing.IArgRmsVoltageMinPhB, ZigBee.IMeasurementAndSensing.IArgRmsVoltageMinPhBPayload>(base.u16, ()=>({
                name: `RMS Voltage Min Ph B`,
                description: `represents the lowest RMS voltage value measured. After resetting, this attribute will return a
value of 0xFFFF until a measurement is made.`,
                id: 0x0906,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Volts,
                
            })),
            RmsVoltageMinPhC: makeType<ZigBee.IMeasurementAndSensing.IArgRmsVoltageMinPhC, ZigBee.IMeasurementAndSensing.IArgRmsVoltageMinPhCPayload>(base.u16, ()=>({
                name: `RMS Voltage Min Ph C`,
                description: ``,
                id: 0x0A06,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            RmsVoltagePhB: makeType<ZigBee.IMeasurementAndSensing.IArgRmsVoltagePhB, ZigBee.IMeasurementAndSensing.IArgRmsVoltagePhBPayload>(base.u16, ()=>({
                name: `RMS Voltage Ph B`,
                description: `represents the most recent RMS voltage reading. If the RMS voltage cannot be measured, a value of 0xFFFF
is returned.`,
                id: 0x0905,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Volts,
                
            })),
            RmsVoltagePhC: makeType<ZigBee.IMeasurementAndSensing.IArgRmsVoltagePhC, ZigBee.IMeasurementAndSensing.IArgRmsVoltagePhCPayload>(base.u16, ()=>({
                name: `RMS Voltage Ph C`,
                description: ``,
                id: 0x0A05,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            RmsVoltageSag: makeType<ZigBee.IMeasurementAndSensing.IArgRmsVoltageSag, ZigBee.IMeasurementAndSensing.IArgRmsVoltageSagPayload>(base.s16, ()=>({
                name: `RMS Voltage Sag`,
                description: `is the RMS voltage below which a sag condition is reported.
The threshold shall be configurable within the specified operating range of the electricity meter.
The value is multiplied and divided by the ACVoltageMultiplier and ACVoltageDivisor, respectively`,
                id: 0x0809,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Volts,
                
            })),
            RmsVoltageSagPeriod: makeType<ZigBee.IMeasurementAndSensing.IArgRmsVoltageSagPeriod, ZigBee.IMeasurementAndSensing.IArgRmsVoltageSagPeriodPayload>(base.u16, ()=>({
                name: `RMS Voltage Sag Period`,
                description: `is the duration used to measure a voltage sag condition.`,
                id: 0x0516,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Seconds,
                
            })),
            RmsVoltageSagPeriodPhB: makeType<ZigBee.IMeasurementAndSensing.IArgRmsVoltageSagPeriodPhB, ZigBee.IMeasurementAndSensing.IArgRmsVoltageSagPeriodPhBPayload>(base.u16, ()=>({
                name: `RMS Voltage Sag Period Ph B`,
                description: ``,
                id: 0x0916,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            RmsVoltageSagPeriodPhC: makeType<ZigBee.IMeasurementAndSensing.IArgRmsVoltageSagPeriodPhC, ZigBee.IMeasurementAndSensing.IArgRmsVoltageSagPeriodPhCPayload>(base.u16, ()=>({
                name: `RMS Voltage Sag Period Ph C`,
                description: ``,
                id: 0x0A16,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            RmsVoltageSwell: makeType<ZigBee.IMeasurementAndSensing.IArgRmsVoltageSwell, ZigBee.IMeasurementAndSensing.IArgRmsVoltageSwellPayload>(base.s16, ()=>({
                name: `RMS Voltage Swell`,
                description: `is the RMS voltage above which a swell condition is reported.
The threshold shall be configurable within the specified operating range of the electricity meter.
The value is multiplied and divided by the ACVoltageMultiplier and ACVoltageDivisor, respectively`,
                id: 0x080A,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Volts,
                
            })),
            RmsVoltageSwellPeriod: makeType<ZigBee.IMeasurementAndSensing.IArgRmsVoltageSwellPeriod, ZigBee.IMeasurementAndSensing.IArgRmsVoltageSwellPeriodPayload>(base.u16, ()=>({
                name: `RMS Voltage Swell Period`,
                description: `is the duration used to measure a voltage swell condition.`,
                id: 0x0517,
                report: false,
                read: true,
                write: true,
                require: false,
                unit: units.Seconds,
                
            })),
            RmsVoltageSwellPeriodPhB: makeType<ZigBee.IMeasurementAndSensing.IArgRmsVoltageSwellPeriodPhB, ZigBee.IMeasurementAndSensing.IArgRmsVoltageSwellPeriodPhBPayload>(base.u16, ()=>({
                name: `RMS Voltage Swell Period Ph B`,
                description: ``,
                id: 0x0917,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            RmsVoltageSwellPeriodPhC: makeType<ZigBee.IMeasurementAndSensing.IArgRmsVoltageSwellPeriodPhC, ZigBee.IMeasurementAndSensing.IArgRmsVoltageSwellPeriodPhCPayload>(base.u16, ()=>({
                name: `RMS Voltage Swell Period Ph C`,
                description: ``,
                id: 0x0A17,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            Scale: makeType<ZigBee.IMeasurementAndSensing.IArgScale, ZigBee.IMeasurementAndSensing.IArgScalePayload>(base.s8, ()=>({
                name: `Scale`,
                description: `indicates the base 10 exponent used to obtain ScaledPressure`,
                id: 0x0014,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            ScaledMaxPressure: makeType<ZigBee.IMeasurementAndSensing.IArgScaledMaxPressure, ZigBee.IMeasurementAndSensing.IArgScaledMaxPressurePayload>(base.s16, ()=>({
                name: `Scaled Max Pressure`,
                description: `indicates the maximum value of ScaledPressure that can be measured. A
value of 0x8000 indicates that this attribute is not defined`,
                id: 0x0012,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.Pascals,
                
            })),
            ScaledMinPressure: makeType<ZigBee.IMeasurementAndSensing.IArgScaledMinPressure, ZigBee.IMeasurementAndSensing.IArgScaledMinPressurePayload>(base.s16, ()=>({
                name: `Scaled Min Pressure`,
                description: `indicates the minimum value of ScaledPressure that can be measured. A
value of 0x8000 indicates that this attribute is not defined`,
                id: 0x0011,
                report: false,
                read: true,
                write: false,
                require: false,
                unit: units.Pascals,
                
            })),
            ScaledPressure: makeType<ZigBee.IMeasurementAndSensing.IArgScaledPressure, ZigBee.IMeasurementAndSensing.IArgScaledPressurePayload>(base.s16, ()=>({
                name: `Scaled Pressure`,
                description: `represents the pressure in Pascals as follows:
ScaledPressure = 10^Scale x Pressure in Pa`,
                id: 0x0010,
                report: true,
                read: true,
                write: false,
                require: false,
                unit: units.Pascals,
                
            })),
            ScaledTolerance: makeType<ZigBee.IMeasurementAndSensing.IArgScaledTolerance, ZigBee.IMeasurementAndSensing.IArgScaledTolerancePayload>(base.u16, ()=>({
                name: `Scaled Tolerance`,
                description: `indicates the magnitude of the possible error that is associated with ScaledPressure.
The true value is located in the range (ScaledPressure – ScaledTolerance) to
(ScaledPressure + ScaledTolerance)`,
                id: 0x0013,
                report: true,
                read: true,
                write: false,
                require: false,
                
            })),
            StartTime: makeType<ZigBee.IMeasurementAndSensing.IArgStartTime, ZigBee.IMeasurementAndSensing.IArgStartTimePayload>(base.utc, ()=>({
                name: `Start Time`,
                description: `is the end time of the most chronologically recent interval being requested.
Example: Data collected from 2:00 PM to 3:00 PM would be specified as a 3:00 PM interval (end time).`,
                
            })),
            Tolerance: makeType<ZigBee.IMeasurementAndSensing.IArgTolerance, ZigBee.IMeasurementAndSensing.IArgTolerancePayload>(base.u16, ()=>({
                name: `Tolerance`,
                description: `indicates the magnitude of the possible error that is associated with MeasuredValue.
The true value is located in the range (MeasuredValue – Tolerance) to
(MeasuredValue + Tolerance)`,
                id: 0x0003,
                report: true,
                read: true,
                write: false,
                require: false,
                
            })),
            TotalActivePower: makeType<ZigBee.IMeasurementAndSensing.IArgTotalActivePower, ZigBee.IMeasurementAndSensing.IArgTotalActivePowerPayload>(base.s32, ()=>({
                name: `Total Active Power`,
                description: `represents the current demand of active power delivered or received at the premises, in kW.
Positive values indicate power delivered to the premises where negative values indicate power received from
the premises. In case if device is capable of measuring multi elements or phases then this will be net active
power value.`,
                id: 0x0304,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.Kilowatts,
                
            })),
            TotalApparentPower: makeType<ZigBee.IMeasurementAndSensing.IArgTotalApparentPower, ZigBee.IMeasurementAndSensing.IArgTotalApparentPowerPayload>(base.u32, ()=>({
                name: `Total Apparent Power`,
                description: `represents the current demand of apparent power.`,
                id: 0x0306,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.KiloVoltAmperes,
                
            })),
            TotalReactivePower: makeType<ZigBee.IMeasurementAndSensing.IArgTotalReactivePower, ZigBee.IMeasurementAndSensing.IArgTotalReactivePowerPayload>(base.s32, ()=>({
                name: `Total Reactive Power`,
                description: `represents the current demand of reactive power delivered or received at the premises, in kVAr. Positive values
indicate power delivered to the premises where negative values indicate power received from the premises.
In case if device is capable of measuring multi elements or phases then this will be net reactive
power value.`,
                id: 0x0305,
                report: false,
                read: false,
                write: false,
                require: false,
                unit: units.KiloVoltAmperesReactive,
                
            })),
            UltrasonicOccupiedToUnoccupiedDelay: makeType<ZigBee.IMeasurementAndSensing.IArgUltrasonicOccupiedToUnoccupiedDelay, ZigBee.IMeasurementAndSensing.IArgUltrasonicOccupiedToUnoccupiedDelayPayload>(base.u16, ()=>({
                name: `Ultrasonic Occupied To Unoccupied Delay`,
                description: `is 16 bits in length and specifies the time delay, in seconds, before the Ultrasonic sensor
changes to its unoccupied state after the last detection of movement in the sensed area`,
                id: 0x0020,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            UltrasonicUnoccupiedToOccupiedDelay: makeType<ZigBee.IMeasurementAndSensing.IArgUltrasonicUnoccupiedToOccupiedDelay, ZigBee.IMeasurementAndSensing.IArgUltrasonicUnoccupiedToOccupiedDelayPayload>(base.u16, ()=>({
                name: `Ultrasonic Unoccupied To Occupied Delay`,
                description: `is 16 bits in length and specifies the time delay, in seconds, before the Ultrasonic sensor changes
to its occupied state after the detection of movement in the sensed area. This attribute is
mandatory if the UltrasonicUnoccupiedToOccupiedThreshold attribute is implemented`,
                id: 0x0021,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })),
            UltrasonicUnoccupiedToOccupiedThreshold: makeType<ZigBee.IMeasurementAndSensing.IArgUltrasonicUnoccupiedToOccupiedThreshold, ZigBee.IMeasurementAndSensing.IArgUltrasonicUnoccupiedToOccupiedThresholdPayload>(base.u8, ()=>({
                name: `Ultrasonic Unoccupied To Occupied Threshold`,
                description: `is 8 bits in length and specifies the number of movement detection events that must occur in the period UltrasonicUnoccupiedToOccupiedDelay, before the Ultrasonic sensor changes to its occupied state. This attribute is mandatory if the UltrasonicUnoccupiedToOccupiedDelay attribute is implemented`,
                id: 0x0022,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })), },
        IlluminanceMeasurement: {
            ID: 0x0400,
            Name: `Illuminance measurement`,
            Desc: `provides attributes and commands for configuring the measurement of
Illuminance, and reporting Illumince measurements`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        IlluminanceLevelSensing: {
            ID: 0x0401,
            Name: `Illuminance level sensing`,
            Desc: `provides attributes and commands for configuring the sensing of
Illuminance levels, and reporting whether Illuminance is above, below,
or on target`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        TemperatureMeasurement: {
            ID: 0x0402,
            Name: `Temperature measurement`,
            Desc: `provides attributes and commands for configuring the measurement of
temperature, and reporting temperature measurements`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        PressureMeasurement: {
            ID: 0x0403,
            Name: `Pressure measurement`,
            Desc: `provides attributes and commands for configuring the measurement of pressure,
and reporting pressure measurements`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        FlowMeasurement: {
            ID: 0x0404,
            Name: `Flow measurement`,
            Desc: `provides attributes and commands for configuring the measurement of flow,
and reporting flow rates`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        RelativeHumidityMeasurement: {
            ID: 0x0405,
            Name: `Relative Humidity measurement`,
            Desc: `provides attributes and commands for configuring the measurement of
relative humidity, and reporting relative humidity measurement`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        OccupancySensing: {
            ID: 0x0406,
            Name: `Occupancy sensing`,
            Desc: `provides attributes and commands for configuring occupancy sensing, and
reporting occupancy status`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        Pm25Measurement: {
            ID: 0x042A,
            Name: `PM2.5 Measurement`,
            Desc: `(Particulate Matter 2.5 microns or less) in µg/m³`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        ElectricalMeasurement: {
            ID: 0x0B04,
            Name: `Electrical Measurement`,
            Desc: `provides a mechanism for querying data about the electrical properties as measured by the device.
This cluster may be implemented on any device type and be implemented on a per-endpoint basis.
For example, a power strip device could represent each outlet on a different endpoint and report
electrical information for each individual outlet. The only caveat is that if you implement an
attribute that has an associated multiplier and divisor, then you must implement the associated
multiplier and divisor attributes. For example if you implement DCVoltage, you must also implement
DCVoltageMultiplier and DCVoltageDivisor`,
            
            GetProfileInfo: makeType<ZigBee.IMeasurementAndSensing.ElectricalMeasurement.ICmdGetProfileInfo, ZigBee.IMeasurementAndSensing.ElectricalMeasurement.ICmdGetProfileInfoPayload>(command, () => ({
                name: `Get Profile Info`,
                description: ``,
                id: 0x0000,
                payload: {}
            })),

            GetMeasurementProfile: makeType<ZigBee.IMeasurementAndSensing.ElectricalMeasurement.ICmdGetMeasurementProfile, ZigBee.IMeasurementAndSensing.ElectricalMeasurement.ICmdGetMeasurementProfilePayload>(command, () => ({
                name: `Get Measurement Profile`,
                description: ``,
                id: 0x0001,
                payload: { 
                    AttributeId: ZigBee.MeasurementAndSensing.Types.AttributeId,
                    StartTime: ZigBee.MeasurementAndSensing.Types.StartTime,
                    NumberOfIntervals: ZigBee.MeasurementAndSensing.Types.NumberOfIntervals,
                }
            })),

            
            GetProfileInfoResponse: makeType<ZigBee.IMeasurementAndSensing.ElectricalMeasurement.ICmdGetProfileInfoResponse, ZigBee.IMeasurementAndSensing.ElectricalMeasurement.ICmdGetProfileInfoResponsePayload>(command, () => ({
                name: `Get Profile Info Response`,
                description: ``,
                id: 0x0000,
                payload: { 
                    ProfileCount: ZigBee.MeasurementAndSensing.Types.ProfileCount,
                    ProfileIntervalPeriod: ZigBee.MeasurementAndSensing.Types.ProfileIntervalPeriod,
                    MaxNumberOfIntervals: ZigBee.MeasurementAndSensing.Types.MaxNumberOfIntervals,
                    ListOfAttributes: ZigBee.MeasurementAndSensing.Types.ListOfAttributes,
                }
            })),

            GetMeasurementProfileResponse: makeType<ZigBee.IMeasurementAndSensing.ElectricalMeasurement.ICmdGetMeasurementProfileResponse, ZigBee.IMeasurementAndSensing.ElectricalMeasurement.ICmdGetMeasurementProfileResponsePayload>(command, () => ({
                name: `Get Measurement Profile Response`,
                description: ``,
                id: 0x0001,
                payload: { 
                    StartTime: ZigBee.MeasurementAndSensing.Types.StartTime,
                    MeasurementResponseStatus: ZigBee.MeasurementAndSensing.Types.MeasurementResponseStatus,
                    ProfileIntervalPeriod: ZigBee.MeasurementAndSensing.Types.ProfileIntervalPeriod,
                    NumberOfIntervals: ZigBee.MeasurementAndSensing.Types.NumberOfIntervals,
                    AttributeId: ZigBee.MeasurementAndSensing.Types.AttributeId,
                    Intervals: ZigBee.MeasurementAndSensing.Types.Intervals,
                }
            })),

            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        CarbonMonoxideMeasurement: {
            ID: 0x040C,
            Name: `Carbon Monoxide Measurement`,
            Desc: `(CO) as concentration in Air`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        CarbonDioxideMeasurement: {
            ID: 0x040D,
            Name: `Carbon Dioxide Measurement`,
            Desc: `(CO2) as concentration in Air`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        EthyleneMeasurement: {
            ID: 0x040E,
            Name: `Ethylene Measurement`,
            Desc: `(CH2) as concentration in Air`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        EthyleneOxideMeasurement: {
            ID: 0x040F,
            Name: `Ethylene Oxide Measurement`,
            Desc: `(C2H40) as concentration in Air`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        HydrogenMeasurement: {
            ID: 0x0410,
            Name: `Hydrogen Measurement`,
            Desc: `(H) as concentration in Air`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        HydrogenSulfideMeasurement: {
            ID: 0x0411,
            Name: `Hydrogen Sulfide Measurement`,
            Desc: `(H2S) as concentration in Air`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        NitricOxideMeasurement: {
            ID: 0x0412,
            Name: `Nitric Oxide Measurement`,
            Desc: `(NO) as concentration in Air`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        NitrogenDioxideMeasurement: {
            ID: 0x0413,
            Name: `Nitrogen Dioxide Measurement`,
            Desc: `(NO2) as concentration in Air`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        OxygenMeasurement: {
            ID: 0x0414,
            Name: `Oxygen Measurement`,
            Desc: `(O2) as concentration in Air`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        OzoneMeasurement: {
            ID: 0x0415,
            Name: `Ozone Measurement`,
            Desc: `(O3) as concentration in Air`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        SulfurDioxideMeasurement: {
            ID: 0x0416,
            Name: `Sulfur Dioxide Measurement`,
            Desc: `(SO2) as concentration in Air`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        DissolvedOxygenMeasurement: {
            ID: 0x0417,
            Name: `Dissolved Oxygen Measurement`,
            Desc: `(DO) as concentration in Water`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        BromateMeasurement: {
            ID: 0x0418,
            Name: `Bromate Measurement`,
            Desc: `as concentration in Drinking Water`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        ChloraminesMeasurement: {
            ID: 0x0419,
            Name: `Chloramines Measurement`,
            Desc: `as concentration in Drinking Water`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        ChlorineMeasurement: {
            ID: 0x041A,
            Name: `Chlorine Measurement`,
            Desc: `as concentration in Drinking Water`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        FecalColiformEColiMeasurement: {
            ID: 0x041B,
            Name: `Fecal coliform & E. Coli Measurement`,
            Desc: `as concentration in Drinking Water`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        FluorideMeasurement: {
            ID: 0x041C,
            Name: `Fluoride Measurement`,
            Desc: `as concentration in Drinking Water`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        HaloaceticAcidsMeasurement: {
            ID: 0x041D,
            Name: `Haloacetic Acids Measurement`,
            Desc: `as concentration in Drinking Water`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        TotalTrihalomethanesMeasurement: {
            ID: 0x041E,
            Name: `Total Trihalomethanes Measurement`,
            Desc: `as concentration in Drinking Water`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        TotalColiformBacteriaMeasurement: {
            ID: 0x041F,
            Name: `Total Coliform Bacteria Measurement`,
            Desc: `as concentration in Drinking Water`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        TurbidityMeasurement: {
            ID: 0x0420,
            Name: `Turbidity Measurement`,
            Desc: `as concentration in Drinking Water`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        CopperMeasurement: {
            ID: 0x0421,
            Name: `Copper Measurement`,
            Desc: `as concentration in Drinking Water`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        LeadMeasurement: {
            ID: 0x0422,
            Name: `Lead Measurement`,
            Desc: `as concentration in Drinking Water`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        ManganeseMeasurement: {
            ID: 0x0423,
            Name: `Manganese Measurement`,
            Desc: `as concentration in Drinking Water`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        SulfateMeasurement: {
            ID: 0x0424,
            Name: `Sulfate Measurement`,
            Desc: `as concentration in Drinking Water`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        BromodichloromethaneMeasurement: {
            ID: 0x0425,
            Name: `Bromodichloromethane Measurement`,
            Desc: `as concentration in Drinking Water`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        BromoformMeasurement: {
            ID: 0x0426,
            Name: `Bromoform Measurement`,
            Desc: `as concentration in Drinking Water`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        ChlorodibromomethaneMeasurement: {
            ID: 0x0427,
            Name: `Chlorodibromomethane Measurement`,
            Desc: `as concentration in Drinking Water`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        ChloroformMeasurement: {
            ID: 0x0428,
            Name: `Chloroform Measurement`,
            Desc: `as concentration in Drinking Water`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        SodiumMeasurement: {
            ID: 0x0429,
            Name: `Sodium Measurement`,
            Desc: `as concentration in Drinking Water`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        },
        FormaldehydeMeasurement: {
            ID: 0x042B,
            Name: `Formaldehyde Measurement`,
            Desc: `as concentration in Air`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        }
    };
    
    ZigBee.MeasurementAndSensing.IlluminanceMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredIlluminance,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredIlluminance,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredIlluminance,
        0x0003: ZigBee.MeasurementAndSensing.Types.Tolerance,
        0x0004: ZigBee.MeasurementAndSensing.Types.LightSensorType,
    };
    ZigBee.MeasurementAndSensing.IlluminanceMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.IlluminanceMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.IlluminanceMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.IlluminanceLevelSensing.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.LevelStatus,
        0x0001: ZigBee.MeasurementAndSensing.Types.IlluminanceLightSensorType,
        0x0010: ZigBee.MeasurementAndSensing.Types.IlluminanceTargetLevel,
    };
    ZigBee.MeasurementAndSensing.IlluminanceLevelSensing.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.IlluminanceLevelSensing.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.IlluminanceLevelSensing.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.TemperatureMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredTemperature,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredTemperature,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredTemperature,
        0x0003: ZigBee.MeasurementAndSensing.Types.Tolerance,
    };
    ZigBee.MeasurementAndSensing.TemperatureMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.TemperatureMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.TemperatureMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.PressureMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredPressure,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredPressure,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredPressure,
        0x0003: ZigBee.MeasurementAndSensing.Types.Tolerance,
        0x0010: ZigBee.MeasurementAndSensing.Types.ScaledPressure,
        0x0011: ZigBee.MeasurementAndSensing.Types.ScaledMinPressure,
        0x0012: ZigBee.MeasurementAndSensing.Types.ScaledMaxPressure,
        0x0013: ZigBee.MeasurementAndSensing.Types.ScaledTolerance,
        0x0014: ZigBee.MeasurementAndSensing.Types.Scale,
    };
    ZigBee.MeasurementAndSensing.PressureMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.PressureMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.PressureMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.FlowMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredFlow,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredFlow,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredFlow,
        0x0003: ZigBee.MeasurementAndSensing.Types.Tolerance,
    };
    ZigBee.MeasurementAndSensing.FlowMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.FlowMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.FlowMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.RelativeHumidityMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredRelativeHumidity,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredRelativeHumidity,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredRelativeHumidity,
        0x0003: ZigBee.MeasurementAndSensing.Types.Tolerance,
    };
    ZigBee.MeasurementAndSensing.RelativeHumidityMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.RelativeHumidityMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.RelativeHumidityMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.OccupancySensing.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.Occupancy,
        0x0001: ZigBee.MeasurementAndSensing.Types.OccupancyType,
        0x0010: ZigBee.MeasurementAndSensing.Types.PirOccupiedToUnoccupiedDelay,
        0x0011: ZigBee.MeasurementAndSensing.Types.PirUnoccupiedToOccupiedDelay,
        0x0012: ZigBee.MeasurementAndSensing.Types.PirUnoccupiedToOccupiedThreshold,
        0x0020: ZigBee.MeasurementAndSensing.Types.UltrasonicOccupiedToUnoccupiedDelay,
        0x0021: ZigBee.MeasurementAndSensing.Types.UltrasonicUnoccupiedToOccupiedDelay,
        0x0022: ZigBee.MeasurementAndSensing.Types.UltrasonicUnoccupiedToOccupiedThreshold,
    };
    ZigBee.MeasurementAndSensing.OccupancySensing.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.OccupancySensing.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.OccupancySensing.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.Pm25Measurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredPm25,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredPm25,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredPm25,
        0x0003: ZigBee.MeasurementAndSensing.Types.Pm25Tolerance,
    };
    ZigBee.MeasurementAndSensing.Pm25Measurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.Pm25Measurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.Pm25Measurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.ElectricalMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.ElectricalMeasurementType,
        0x0100: ZigBee.MeasurementAndSensing.Types.DcVoltage,
        0x0101: ZigBee.MeasurementAndSensing.Types.DcVoltageMin,
        0x0102: ZigBee.MeasurementAndSensing.Types.DcVoltageMax,
        0x0103: ZigBee.MeasurementAndSensing.Types.DcCurrent,
        0x0104: ZigBee.MeasurementAndSensing.Types.DcCurrentMin,
        0x0105: ZigBee.MeasurementAndSensing.Types.DcCurrentMax,
        0x0106: ZigBee.MeasurementAndSensing.Types.DcPower,
        0x0107: ZigBee.MeasurementAndSensing.Types.DcPowerMin,
        0x0108: ZigBee.MeasurementAndSensing.Types.DcPowerMax,
        0x0200: ZigBee.MeasurementAndSensing.Types.DcVoltageMultiplier,
        0x0201: ZigBee.MeasurementAndSensing.Types.DcVoltageDivisor,
        0x0202: ZigBee.MeasurementAndSensing.Types.DcCurrentMultiplier,
        0x0203: ZigBee.MeasurementAndSensing.Types.DcCurrentDivisor,
        0x0204: ZigBee.MeasurementAndSensing.Types.DcPowerMultiplier,
        0x0205: ZigBee.MeasurementAndSensing.Types.DcPowerDivisor,
        0x0300: ZigBee.MeasurementAndSensing.Types.AcFrequency,
        0x0301: ZigBee.MeasurementAndSensing.Types.AcFrequencyMin,
        0x0302: ZigBee.MeasurementAndSensing.Types.AcFrequencyMax,
        0x0303: ZigBee.MeasurementAndSensing.Types.NeutralCurrent,
        0x0304: ZigBee.MeasurementAndSensing.Types.TotalActivePower,
        0x0305: ZigBee.MeasurementAndSensing.Types.TotalReactivePower,
        0x0306: ZigBee.MeasurementAndSensing.Types.TotalApparentPower,
        0x0307: ZigBee.MeasurementAndSensing.Types.Measured1StHarmonicCurrent,
        0x0308: ZigBee.MeasurementAndSensing.Types.Measured3RdHarmonicCurrent,
        0x0309: ZigBee.MeasurementAndSensing.Types.Measured5ThHarmonicCurrent,
        0x030A: ZigBee.MeasurementAndSensing.Types.Measured7ThHarmonicCurrent,
        0x030B: ZigBee.MeasurementAndSensing.Types.Measured9ThHarmonicCurrent,
        0x030C: ZigBee.MeasurementAndSensing.Types.Measured11ThHarmonicCurrent,
        0x030D: ZigBee.MeasurementAndSensing.Types.MeasuredPhase1StHarmonicCurrent,
        0x030E: ZigBee.MeasurementAndSensing.Types.MeasuredPhase3RdHarmonicCurrent,
        0x030F: ZigBee.MeasurementAndSensing.Types.MeasuredPhase5ThHarmonicCurrent,
        0x0310: ZigBee.MeasurementAndSensing.Types.MeasuredPhase7ThHarmonicCurrent,
        0x0311: ZigBee.MeasurementAndSensing.Types.MeasuredPhase9ThHarmonicCurrent,
        0x0312: ZigBee.MeasurementAndSensing.Types.MeasuredPhase11ThHarmonicCurrent,
        0x0400: ZigBee.MeasurementAndSensing.Types.AcFrequencyMultiplier,
        0x0401: ZigBee.MeasurementAndSensing.Types.AcFrequencyDivisor,
        0x0402: ZigBee.MeasurementAndSensing.Types.PowerMultiplier,
        0x0403: ZigBee.MeasurementAndSensing.Types.PowerDivisor,
        0x0404: ZigBee.MeasurementAndSensing.Types.HarmonicCurrentMultiplier,
        0x0405: ZigBee.MeasurementAndSensing.Types.PhaseHarmonicCurrentMultiplier,
        0x0501: ZigBee.MeasurementAndSensing.Types.LineCurrent,
        0x0502: ZigBee.MeasurementAndSensing.Types.ActiveCurrent,
        0x0503: ZigBee.MeasurementAndSensing.Types.ReactiveCurrent,
        0x0505: ZigBee.MeasurementAndSensing.Types.RmsVoltage,
        0x0506: ZigBee.MeasurementAndSensing.Types.RmsVoltageMin,
        0x0507: ZigBee.MeasurementAndSensing.Types.RmsVoltageMax,
        0x0508: ZigBee.MeasurementAndSensing.Types.RmsCurrent,
        0x0509: ZigBee.MeasurementAndSensing.Types.RmsCurrentMin,
        0x050A: ZigBee.MeasurementAndSensing.Types.RmsCurrentMax,
        0x050B: ZigBee.MeasurementAndSensing.Types.ActivePower,
        0x050C: ZigBee.MeasurementAndSensing.Types.ActivePowerMin,
        0x050D: ZigBee.MeasurementAndSensing.Types.ActivePowerMax,
        0x050E: ZigBee.MeasurementAndSensing.Types.ReactivePower,
        0x050F: ZigBee.MeasurementAndSensing.Types.ApparentPower,
        0x0510: ZigBee.MeasurementAndSensing.Types.PowerFactor,
        0x0511: ZigBee.MeasurementAndSensing.Types.AverageRmsVoltageMeasurementPeriod,
        0x0512: ZigBee.MeasurementAndSensing.Types.AverageRmsOverVoltageCounter,
        0x0513: ZigBee.MeasurementAndSensing.Types.AverageRmsUnderVoltageCounter,
        0x0514: ZigBee.MeasurementAndSensing.Types.RmsExtremeOverVoltagePeriod,
        0x0515: ZigBee.MeasurementAndSensing.Types.RmsExtremeUnderVoltagePeriod,
        0x0516: ZigBee.MeasurementAndSensing.Types.RmsVoltageSagPeriod,
        0x0517: ZigBee.MeasurementAndSensing.Types.RmsVoltageSwellPeriod,
        0x0600: ZigBee.MeasurementAndSensing.Types.AcVoltageMultiplier,
        0x0601: ZigBee.MeasurementAndSensing.Types.AcVoltageDivisor,
        0x0602: ZigBee.MeasurementAndSensing.Types.AcCurrentMultiplier,
        0x0603: ZigBee.MeasurementAndSensing.Types.AcCurrentDivisor,
        0x0604: ZigBee.MeasurementAndSensing.Types.AcPowerMultiplier,
        0x0605: ZigBee.MeasurementAndSensing.Types.AcPowerDivisor,
        0x0700: ZigBee.MeasurementAndSensing.Types.DcOverloadAlarmsMask,
        0x0701: ZigBee.MeasurementAndSensing.Types.DcVoltageOverload,
        0x0702: ZigBee.MeasurementAndSensing.Types.DcCurrentOverload,
        0x0800: ZigBee.MeasurementAndSensing.Types.AcAlarmsMask,
        0x0801: ZigBee.MeasurementAndSensing.Types.AcVoltageOverload,
        0x0802: ZigBee.MeasurementAndSensing.Types.AcCurrentOverload,
        0x0803: ZigBee.MeasurementAndSensing.Types.AcActivePowerOverload,
        0x0804: ZigBee.MeasurementAndSensing.Types.AcReactivePowerOverload,
        0x0805: ZigBee.MeasurementAndSensing.Types.AverageRmsOverVoltage,
        0x0806: ZigBee.MeasurementAndSensing.Types.AverageRmsUnderVoltage,
        0x0807: ZigBee.MeasurementAndSensing.Types.RmsExtremeOverVoltage,
        0x0808: ZigBee.MeasurementAndSensing.Types.RmsExtremeUnderVoltage,
        0x0809: ZigBee.MeasurementAndSensing.Types.RmsVoltageSag,
        0x080A: ZigBee.MeasurementAndSensing.Types.RmsVoltageSwell,
        0x0901: ZigBee.MeasurementAndSensing.Types.LineCurrentPhB,
        0x0902: ZigBee.MeasurementAndSensing.Types.ActiveCurrentPhB,
        0x0903: ZigBee.MeasurementAndSensing.Types.ReactiveCurrentPhB,
        0x0905: ZigBee.MeasurementAndSensing.Types.RmsVoltagePhB,
        0x0906: ZigBee.MeasurementAndSensing.Types.RmsVoltageMinPhB,
        0x0907: ZigBee.MeasurementAndSensing.Types.RmsVoltageMaxPhB,
        0x0908: ZigBee.MeasurementAndSensing.Types.RmsCurrentPhB,
        0x0909: ZigBee.MeasurementAndSensing.Types.RmsCurrentMinPhB,
        0x090A: ZigBee.MeasurementAndSensing.Types.RmsCurrentMaxPhB,
        0x090B: ZigBee.MeasurementAndSensing.Types.ActivePowerPhB,
        0x090C: ZigBee.MeasurementAndSensing.Types.ActivePowerMinPhB,
        0x090D: ZigBee.MeasurementAndSensing.Types.ActivePowerMaxPhB,
        0x090E: ZigBee.MeasurementAndSensing.Types.ReactivePowerPhB,
        0x090F: ZigBee.MeasurementAndSensing.Types.ApparentPowerPhB,
        0x0910: ZigBee.MeasurementAndSensing.Types.PowerFactorPhB,
        0x0911: ZigBee.MeasurementAndSensing.Types.AverageRmsVoltageMeasurementPeriodPhB,
        0x0912: ZigBee.MeasurementAndSensing.Types.AverageRmsOverVoltageCounterPhB,
        0x0913: ZigBee.MeasurementAndSensing.Types.AverageRmsUnderVoltageCounterPhB,
        0x0914: ZigBee.MeasurementAndSensing.Types.RmsExtremeOverVoltagePeriodPhB,
        0x0915: ZigBee.MeasurementAndSensing.Types.RmsExtremeUnderVoltagePeriodPhB,
        0x0916: ZigBee.MeasurementAndSensing.Types.RmsVoltageSagPeriodPhB,
        0x0917: ZigBee.MeasurementAndSensing.Types.RmsVoltageSwellPeriodPhB,
        0x0A01: ZigBee.MeasurementAndSensing.Types.LineCurrentPhC,
        0x0A02: ZigBee.MeasurementAndSensing.Types.ActiveCurrentPhC,
        0x0A03: ZigBee.MeasurementAndSensing.Types.ReactiveCurrentPhC,
        0x0A05: ZigBee.MeasurementAndSensing.Types.RmsVoltagePhC,
        0x0A06: ZigBee.MeasurementAndSensing.Types.RmsVoltageMinPhC,
        0x0A07: ZigBee.MeasurementAndSensing.Types.RmsVoltageMaxPhC,
        0x0A08: ZigBee.MeasurementAndSensing.Types.RmsCurrentPhC,
        0x0A09: ZigBee.MeasurementAndSensing.Types.RmsCurrentMinPhC,
        0x0A0A: ZigBee.MeasurementAndSensing.Types.RmsCurrentMaxPhC,
        0x0A0B: ZigBee.MeasurementAndSensing.Types.ActivePowerPhC,
        0x0A0C: ZigBee.MeasurementAndSensing.Types.ActivePowerMinPhC,
        0x0A0D: ZigBee.MeasurementAndSensing.Types.ActivePowerMaxPhC,
        0x0A0E: ZigBee.MeasurementAndSensing.Types.ReactivePowerPhC,
        0x0A0F: ZigBee.MeasurementAndSensing.Types.ApparentPowerPhC,
        0x0A10: ZigBee.MeasurementAndSensing.Types.PowerFactorPhC,
        0x0A11: ZigBee.MeasurementAndSensing.Types.AverageRmsVoltageMeasurementPeriodPhC,
        0x0A12: ZigBee.MeasurementAndSensing.Types.AverageRmsOverVoltageCounterPhC,
        0x0A13: ZigBee.MeasurementAndSensing.Types.AverageRmsUnderVoltageCounterPhC,
        0x0A14: ZigBee.MeasurementAndSensing.Types.RmsExtremeOverVoltagePeriodPhC,
        0x0A15: ZigBee.MeasurementAndSensing.Types.RmsExtremeUnderVoltagePeriodPhC,
        0x0A16: ZigBee.MeasurementAndSensing.Types.RmsVoltageSagPeriodPhC,
        0x0A17: ZigBee.MeasurementAndSensing.Types.RmsVoltageSwellPeriodPhC,
    };
    ZigBee.MeasurementAndSensing.ElectricalMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.ElectricalMeasurement.Server.Command = { 
        0x00: ZigBee.MeasurementAndSensing.ElectricalMeasurement.GetProfileInfo,
        0x01: ZigBee.MeasurementAndSensing.ElectricalMeasurement.GetMeasurementProfile,
    };
    ZigBee.MeasurementAndSensing.ElectricalMeasurement.Client.Command = { 
        0x00: ZigBee.MeasurementAndSensing.ElectricalMeasurement.GetProfileInfoResponse,
        0x01: ZigBee.MeasurementAndSensing.ElectricalMeasurement.GetMeasurementProfileResponse,
    };
    ZigBee.MeasurementAndSensing.CarbonMonoxideMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.CarbonMonoxideMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.CarbonMonoxideMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.CarbonMonoxideMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.CarbonDioxideMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.CarbonDioxideMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.CarbonDioxideMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.CarbonDioxideMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.EthyleneMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.EthyleneMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.EthyleneMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.EthyleneMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.EthyleneOxideMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.EthyleneOxideMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.EthyleneOxideMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.EthyleneOxideMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.HydrogenMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.HydrogenMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.HydrogenMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.HydrogenMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.HydrogenSulfideMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.HydrogenSulfideMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.HydrogenSulfideMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.HydrogenSulfideMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.NitricOxideMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.NitricOxideMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.NitricOxideMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.NitricOxideMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.NitrogenDioxideMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.NitrogenDioxideMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.NitrogenDioxideMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.NitrogenDioxideMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.OxygenMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.OxygenMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.OxygenMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.OxygenMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.OzoneMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.OzoneMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.OzoneMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.OzoneMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.SulfurDioxideMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.SulfurDioxideMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.SulfurDioxideMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.SulfurDioxideMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.DissolvedOxygenMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.DissolvedOxygenMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.DissolvedOxygenMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.DissolvedOxygenMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.BromateMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.BromateMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.BromateMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.BromateMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.ChloraminesMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.ChloraminesMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.ChloraminesMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.ChloraminesMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.ChlorineMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.ChlorineMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.ChlorineMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.ChlorineMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.FecalColiformEColiMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.FecalColiformEColiMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.FecalColiformEColiMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.FecalColiformEColiMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.FluorideMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.FluorideMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.FluorideMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.FluorideMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.HaloaceticAcidsMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.HaloaceticAcidsMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.HaloaceticAcidsMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.HaloaceticAcidsMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.TotalTrihalomethanesMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.TotalTrihalomethanesMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.TotalTrihalomethanesMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.TotalTrihalomethanesMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.TotalColiformBacteriaMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.TotalColiformBacteriaMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.TotalColiformBacteriaMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.TotalColiformBacteriaMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.TurbidityMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.TurbidityMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.TurbidityMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.TurbidityMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.CopperMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.CopperMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.CopperMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.CopperMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.LeadMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.LeadMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.LeadMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.LeadMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.ManganeseMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.ManganeseMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.ManganeseMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.ManganeseMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.SulfateMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.SulfateMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.SulfateMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.SulfateMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.BromodichloromethaneMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.BromodichloromethaneMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.BromodichloromethaneMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.BromodichloromethaneMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.BromoformMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.BromoformMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.BromoformMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.BromoformMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.ChlorodibromomethaneMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.ChlorodibromomethaneMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.ChlorodibromomethaneMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.ChlorodibromomethaneMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.ChloroformMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.ChloroformMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.ChloroformMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.ChloroformMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.SodiumMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.SodiumMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.SodiumMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.SodiumMeasurement.Client.Command = { 
    };
    ZigBee.MeasurementAndSensing.FormaldehydeMeasurement.Server.Attribute = { 
        0x0000: ZigBee.MeasurementAndSensing.Types.MeasuredConcentration,
        0x0001: ZigBee.MeasurementAndSensing.Types.MinMeasuredConcentration,
        0x0002: ZigBee.MeasurementAndSensing.Types.MaxMeasuredConcentration,
        0x0003: ZigBee.MeasurementAndSensing.Types.ConcentrationTolerance,
    };
    ZigBee.MeasurementAndSensing.FormaldehydeMeasurement.Client.Attribute = { 
    };
    ZigBee.MeasurementAndSensing.FormaldehydeMeasurement.Server.Command = { 
    };
    ZigBee.MeasurementAndSensing.FormaldehydeMeasurement.Client.Command = { 
    };
    export const Otau = {
        Types: { 
            BlockRequestOptions: makeType<ZigBee.IOtau.IArgBlockRequestOptions, ZigBee.IOtau.IArgBlockRequestOptionsPayload>(base.bmp8, ()=>({
                name: `Block Request Options`,
                description: ``,
                bits: { 
                1: `Request Node Address`, 
                2: `Minimum Block Period`,  },
                
            })),
            CurrentFileVersion: makeType<ZigBee.IOtau.IArgCurrentFileVersion, ZigBee.IOtau.IArgCurrentFileVersionPayload>(base.u32, ()=>({
                name: `Current File Version`,
                description: ``,
                id: 0x0002,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            CurrentStackVersion: makeType<ZigBee.IOtau.IArgCurrentStackVersion, ZigBee.IOtau.IArgCurrentStackVersionPayload>(base.u16, ()=>({
                name: `Current Stack Version`,
                description: ``,
                id: 0x0003,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            CurrentTime: makeType<ZigBee.IOtau.IArgCurrentTime, ZigBee.IOtau.IArgCurrentTimePayload>(base.time, ()=>({
                name: `Current Time`,
                description: ``,
                
            })),
            DataSize: makeType<ZigBee.IOtau.IArgDataSize, ZigBee.IOtau.IArgDataSizePayload>(base.u8, ()=>({
                name: `Data Size`,
                description: ``,
                
            })),
            DownloadedFileVersion: makeType<ZigBee.IOtau.IArgDownloadedFileVersion, ZigBee.IOtau.IArgDownloadedFileVersionPayload>(base.u32, ()=>({
                name: `Downloaded File Version`,
                description: ``,
                id: 0x0004,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            DownloadedStackVersion: makeType<ZigBee.IOtau.IArgDownloadedStackVersion, ZigBee.IOtau.IArgDownloadedStackVersionPayload>(base.u16, ()=>({
                name: `Downloaded Stack Version`,
                description: ``,
                id: 0x0005,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            FileOffset: makeType<ZigBee.IOtau.IArgFileOffset, ZigBee.IOtau.IArgFileOffsetPayload>(base.u32, ()=>({
                name: `File Offset`,
                description: ``,
                id: 0x0001,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            FileVersion: makeType<ZigBee.IOtau.IArgFileVersion, ZigBee.IOtau.IArgFileVersionPayload>(base.u32, ()=>({
                name: `File Version`,
                description: ``,
                
            })),
            HardwareVersion: makeType<ZigBee.IOtau.IArgHardwareVersion, ZigBee.IOtau.IArgHardwareVersionPayload>(base.u16, ()=>({
                name: `Hardware Version`,
                description: ``,
                
            })),
            HardwareVersionPresent: makeType<ZigBee.IOtau.IArgHardwareVersionPresent, ZigBee.IOtau.IArgHardwareVersionPresentPayload>(base.bool, ()=>({
                name: `Hardware Version Present`,
                description: ``,
                
            })),
            ImageNotifyPayloadType: makeType<ZigBee.IOtau.IArgImageNotifyPayloadType, ZigBee.IOtau.IArgImageNotifyPayloadTypePayload>(base.enum8, ()=>({
                name: `Image Notify Payload Type`,
                description: ``,
                values: { 
                0x00: `Query jitter`, 
                0x01: `Query jitter and manufacturer code`, 
                0x02: `Query jitter, manufacturer code, and image type`, 
                0x03: `Query jitter, manufacturer code, image type, and new file version`,  },
                
            })),
            ImageSize: makeType<ZigBee.IOtau.IArgImageSize, ZigBee.IOtau.IArgImageSizePayload>(base.u32, ()=>({
                name: `Image Size`,
                description: ``,
                unit: units.Bytes,
                
            })),
            ImageType: makeType<ZigBee.IOtau.IArgImageType, ZigBee.IOtau.IArgImageTypePayload>(base.u16, ()=>({
                name: `Image type`,
                description: ``,
                values: { 
                0x0000: `Specific image`, 
                0xFFC0: `Security credential`, 
                0xFFC1: `Configuration`, 
                0xFFC2: `Log`, 
                0xFFFF: `Wild card`,  },
                
            })),
            ImageUpgradeStatus: makeType<ZigBee.IOtau.IArgImageUpgradeStatus, ZigBee.IOtau.IArgImageUpgradeStatusPayload>(base.enum8, ()=>({
                name: `Image upgrade status`,
                description: ``,
                id: 0x0006,
                report: false,
                read: true,
                write: true,
                require: false,
                values: { 
                0x00: `Normal`, 
                0x01: `Download in progress`, 
                0x02: `Download complete`, 
                0x03: `Waiting to upgrade`, 
                0x04: `Count down`, 
                0x05: `Wait for more`,  },
                
            })),
            ManufacturerCode: makeType<ZigBee.IOtau.IArgManufacturerCode, ZigBee.IOtau.IArgManufacturerCodePayload>(base.enum16, ()=>({
                name: `Manufacturer Code`,
                description: ``,
                values: { 
                0x117C: `Ikea`,  },
                
            })),
            MinBlockRequestDelay: makeType<ZigBee.IOtau.IArgMinBlockRequestDelay, ZigBee.IOtau.IArgMinBlockRequestDelayPayload>(base.u16, ()=>({
                name: `Min block request delay`,
                description: ``,
                id: 0x0009,
                report: false,
                read: true,
                write: false,
                require: false,
                
            })),
            NextImageStatus: makeType<ZigBee.IOtau.IArgNextImageStatus, ZigBee.IOtau.IArgNextImageStatusPayload>(base.enum8, ()=>({
                name: `Next Image Status`,
                description: ``,
                values: { 
                0x00: `Success`, 
                0x98: `No updates`,  },
                
            })),
            Payload: makeType<ZigBee.IOtau.IArgPayload, ZigBee.IOtau.IArgPayloadPayload>(base.ostring, ()=>({
                name: `Payload`,
                description: ``,
                
            })),
            QueryJitter: makeType<ZigBee.IOtau.IArgQueryJitter, ZigBee.IOtau.IArgQueryJitterPayload>(base.u8, ()=>({
                name: `Query jitter`,
                description: ``,
                
            })),
            RequestTime: makeType<ZigBee.IOtau.IArgRequestTime, ZigBee.IOtau.IArgRequestTimePayload>(base.time, ()=>({
                name: `Request Time`,
                description: ``,
                
            })),
            Status: makeType<ZigBee.IOtau.IArgStatus, ZigBee.IOtau.IArgStatusPayload>(base.enum8, ()=>({
                name: `Status`,
                description: ``,
                values: { 
                0x00: `Success`, 
                0x7E: `Not Authorized`, 
                0x80: `Malformed Command`, 
                0x81: `Unsupported Command`, 
                0x95: `Abort`, 
                0x96: `Invalid Image`, 
                0x97: `Wait for data`, 
                0x98: `No Image Available`, 
                0x99: `Require More Image`,  },
                
            })),
            UpgradeServer: makeType<ZigBee.IOtau.IArgUpgradeServer, ZigBee.IOtau.IArgUpgradeServerPayload>(base.uid, ()=>({
                name: `Upgrade server`,
                description: ``,
                id: 0x0000,
                report: false,
                read: true,
                write: true,
                require: false,
                
            })), },
        Otau: {
            ID: 0x0019,
            Name: `OTAU`,
            Desc: `Over the air upgrade.`,
            
            QueryNextImage: makeType<ZigBee.IOtau.Otau.ICmdQueryNextImage, ZigBee.IOtau.Otau.ICmdQueryNextImagePayload>(command, () => ({
                name: `Query next image`,
                description: ``,
                id: 0x0001,
                payload: { 
                    HardwareVersionPresent: ZigBee.Otau.Types.HardwareVersionPresent,
                    ManufacturerCode: ZigBee.Otau.Types.ManufacturerCode,
                    ImageType: ZigBee.Otau.Types.ImageType,
                    FileVersion: ZigBee.Otau.Types.FileVersion,
                    HardwareVersion: ZigBee.Otau.Types.HardwareVersion,
                }
            })),

            ImageBlockRequest: makeType<ZigBee.IOtau.Otau.ICmdImageBlockRequest, ZigBee.IOtau.Otau.ICmdImageBlockRequestPayload>(command, () => ({
                name: `Image Block Request`,
                description: ``,
                id: 0x0003,
                payload: { 
                    BlockRequestOptions: ZigBee.Otau.Types.BlockRequestOptions,
                    ManufacturerCode: ZigBee.Otau.Types.ManufacturerCode,
                    ImageType: ZigBee.Otau.Types.ImageType,
                    FileVersion: ZigBee.Otau.Types.FileVersion,
                    FileOffset: ZigBee.Otau.Types.FileOffset,
                    MaxDataSize: ZigBee.Otau.Types.DataSize,
                }
            })),

            UpgradeEndRequest: makeType<ZigBee.IOtau.Otau.ICmdUpgradeEndRequest, ZigBee.IOtau.Otau.ICmdUpgradeEndRequestPayload>(command, () => ({
                name: `Upgrade End Request`,
                description: ``,
                id: 0x0006,
                payload: { 
                    Status: ZigBee.Otau.Types.Status,
                    ManufacturerCode: ZigBee.Otau.Types.ManufacturerCode,
                    ImageType: ZigBee.Otau.Types.ImageType,
                    FileVersion: ZigBee.Otau.Types.FileVersion,
                }
            })),

            
            ImageNotify: makeType<ZigBee.IOtau.Otau.ICmdImageNotify, ZigBee.IOtau.Otau.ICmdImageNotifyPayload>(command, () => ({
                name: `Image notify`,
                description: ``,
                id: 0x0000,
                payload: { 
                    ImageNotifyPayloadType: ZigBee.Otau.Types.ImageNotifyPayloadType,
                    QueryJitter: ZigBee.Otau.Types.QueryJitter,
                    ManufacturerCode: ZigBee.Otau.Types.ManufacturerCode,
                    ImageType: ZigBee.Otau.Types.ImageType,
                    FileVersion: ZigBee.Otau.Types.FileVersion,
                }
            })),

            QueryNextImageResponse: makeType<ZigBee.IOtau.Otau.ICmdQueryNextImageResponse, ZigBee.IOtau.Otau.ICmdQueryNextImageResponsePayload>(command, () => ({
                name: `Query next image response`,
                description: ``,
                id: 0x0002,
                payload: { 
                    Status: ZigBee.Otau.Types.Status,
                    ManufacturerCode: ZigBee.Otau.Types.ManufacturerCode,
                    ImageType: ZigBee.Otau.Types.ImageType,
                    FileVersion: ZigBee.Otau.Types.FileVersion,
                    ImageSize: ZigBee.Otau.Types.ImageSize,
                }
            })),

            ImageBlockResponse: makeType<ZigBee.IOtau.Otau.ICmdImageBlockResponse, ZigBee.IOtau.Otau.ICmdImageBlockResponsePayload>(command, () => ({
                name: `Image Block Response`,
                description: ``,
                id: 0x0005,
                payload: { 
                    Status: ZigBee.Otau.Types.Status,
                    ManufacturerCode: ZigBee.Otau.Types.ManufacturerCode,
                    ImageType: ZigBee.Otau.Types.ImageType,
                    FileVersion: ZigBee.Otau.Types.FileVersion,
                    FileOffset: ZigBee.Otau.Types.FileOffset,
                    Payload: ZigBee.Otau.Types.Payload,
                    CurrentTime: ZigBee.Otau.Types.CurrentTime,
                    RequestTime: ZigBee.Otau.Types.RequestTime,
                }
            })),

            UpgradeEndResponse: makeType<ZigBee.IOtau.Otau.ICmdUpgradeEndResponse, ZigBee.IOtau.Otau.ICmdUpgradeEndResponsePayload>(command, () => ({
                name: `Upgrade End Response`,
                description: ``,
                id: 0x0007,
                payload: { 
                    ManufacturerCode: ZigBee.Otau.Types.ManufacturerCode,
                    ImageType: ZigBee.Otau.Types.ImageType,
                    FileVersion: ZigBee.Otau.Types.FileVersion,
                    CurrentTime: ZigBee.Otau.Types.CurrentTime,
                    RequestTime: ZigBee.Otau.Types.RequestTime,
                }
            })),

            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        }
    };
    
    ZigBee.Otau.Otau.Server.Attribute = { 
    };
    ZigBee.Otau.Otau.Client.Attribute = { 
        0x0000: ZigBee.Otau.Types.UpgradeServer,
        0x0001: ZigBee.Otau.Types.FileOffset,
        0x0002: ZigBee.Otau.Types.CurrentFileVersion,
        0x0003: ZigBee.Otau.Types.CurrentStackVersion,
        0x0004: ZigBee.Otau.Types.DownloadedFileVersion,
        0x0005: ZigBee.Otau.Types.DownloadedStackVersion,
        0x0006: ZigBee.Otau.Types.ImageUpgradeStatus,
        0x0009: ZigBee.Otau.Types.MinBlockRequestDelay,
    };
    ZigBee.Otau.Otau.Server.Command = { 
        0x01: ZigBee.Otau.Otau.QueryNextImage,
        0x03: ZigBee.Otau.Otau.ImageBlockRequest,
        0x06: ZigBee.Otau.Otau.UpgradeEndRequest,
    };
    ZigBee.Otau.Otau.Client.Command = { 
        0x00: ZigBee.Otau.Otau.ImageNotify,
        0x02: ZigBee.Otau.Otau.QueryNextImageResponse,
        0x05: ZigBee.Otau.Otau.ImageBlockResponse,
        0x07: ZigBee.Otau.Otau.UpgradeEndResponse,
    };
    export const SmartEnergy = {
        Types: { 
            ControlTemperature: makeType<ZigBee.ISmartEnergy.IArgControlTemperature, ZigBee.ISmartEnergy.IArgControlTemperaturePayload>(base.s24, ()=>({
                name: `Control Temperature`,
                description: ``,
                id: 0x0019,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            CurrentBlock: makeType<ZigBee.ISmartEnergy.IArgCurrentBlock, ZigBee.ISmartEnergy.IArgCurrentBlockPayload>(base.enum8, ()=>({
                name: `Current Block`,
                description: `is an 8-bit Enumeration which indicates the currently active block. If blocks are active then the current
active block is based on the CurrentBlockPeriodConsumptionDelivered and the block thresholds.
Block 1 is active when the value of CurrentBlockPeriodConsumptionDelivered is less than Block1Threshold value;
Block 2 is active when CurrentBlockPeriodConsumptionDelivered is greater than Block1Threshold value and less than
Block2Threshold value, and so on. Block 16 is active when the value of CurrentBlockPeriodConsumptionDelivered is
greater than Block15Threshold value.`,
                id: 0x000E,
                report: false,
                read: false,
                write: false,
                require: false,
                values: { 
                0x00: `No blocks in use`, 
                0x01: `Block1`, 
                0x02: `Block2`, 
                0x03: `Block3`, 
                0x04: `Block4`, 
                0x05: `Block5`, 
                0x06: `Block6`, 
                0x07: `Block7`, 
                0x08: `Block8`, 
                0x09: `Block9`, 
                0x0A: `Block10`, 
                0x0B: `Block11`, 
                0x0C: `Block12`, 
                0x0D: `Block13`, 
                0x0E: `Block14`, 
                0x0F: `Block15`, 
                0x10: `Block16`,  },
                
            })),
            CurrentBlockPeriodConsumptionDelivered: makeType<ZigBee.ISmartEnergy.IArgCurrentBlockPeriodConsumptionDelivered, ZigBee.ISmartEnergy.IArgCurrentBlockPeriodConsumptionDeliveredPayload>(base.u48, ()=>({
                name: `Current Block Period Consumption Delivered`,
                description: `represents the most recent summed value of Energy, Gas or Water delivered and consumed in the premises during the
Block Tariff Period. The CurrentBlockPeriodConsumptionDelivered is reset at the start of each Block Tariff Period.`,
                id: 0x000C,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            CurrentInletEnergyCarrierDemand: makeType<ZigBee.ISmartEnergy.IArgCurrentInletEnergyCarrierDemand, ZigBee.ISmartEnergy.IArgCurrentInletEnergyCarrierDemandPayload>(base.s24, ()=>({
                name: `Current Inlet Energy Carrier Demand`,
                description: ``,
                id: 0x001A,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            CurrentInletEnergyCarrierSummation: makeType<ZigBee.ISmartEnergy.IArgCurrentInletEnergyCarrierSummation, ZigBee.ISmartEnergy.IArgCurrentInletEnergyCarrierSummationPayload>(base.u48, ()=>({
                name: `Current Inlet Energy Carrier Summation`,
                description: ``,
                id: 0x0015,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            CurrentMaxDemandDelivered: makeType<ZigBee.ISmartEnergy.IArgCurrentMaxDemandDelivered, ZigBee.ISmartEnergy.IArgCurrentMaxDemandDeliveredPayload>(base.u48, ()=>({
                name: `Current Max Demand Delivered`,
                description: `represents the maximum demand or rate of delivered value of Energy, Gas, or Water being utilized at the premises`,
                id: 0x0002,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            CurrentMaxDemandDeliveredTime: makeType<ZigBee.ISmartEnergy.IArgCurrentMaxDemandDeliveredTime, ZigBee.ISmartEnergy.IArgCurrentMaxDemandDeliveredTimePayload>(base.utc, ()=>({
                name: `Current Max Demand Delivered Time`,
                description: `represents the time when CurrentMaxDemandDelivered reading was captured`,
                id: 0x0008,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            CurrentMaxDemandReceived: makeType<ZigBee.ISmartEnergy.IArgCurrentMaxDemandReceived, ZigBee.ISmartEnergy.IArgCurrentMaxDemandReceivedPayload>(base.u48, ()=>({
                name: `Current Max Demand Received`,
                description: `represents the maximum demand or rate of received value of Energy, Gas, or Water being utilized by the utility`,
                id: 0x0003,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            CurrentMaxDemandReceivedTime: makeType<ZigBee.ISmartEnergy.IArgCurrentMaxDemandReceivedTime, ZigBee.ISmartEnergy.IArgCurrentMaxDemandReceivedTimePayload>(base.utc, ()=>({
                name: `Current Max Demand Received Time`,
                description: `represents the time when CurrentMaxDemandReceived reading was captured`,
                id: 0x0009,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            CurrentOutletEnergyCarrierDemand: makeType<ZigBee.ISmartEnergy.IArgCurrentOutletEnergyCarrierDemand, ZigBee.ISmartEnergy.IArgCurrentOutletEnergyCarrierDemandPayload>(base.s24, ()=>({
                name: `Current Outlet Energy Carrier Demand`,
                description: ``,
                id: 0x001B,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            CurrentOutletEnergyCarrierSummation: makeType<ZigBee.ISmartEnergy.IArgCurrentOutletEnergyCarrierSummation, ZigBee.ISmartEnergy.IArgCurrentOutletEnergyCarrierSummationPayload>(base.u48, ()=>({
                name: `Current Outlet Energy Carrier Summation`,
                description: ``,
                id: 0x0016,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            CurrentSummationDelivered: makeType<ZigBee.ISmartEnergy.IArgCurrentSummationDelivered, ZigBee.ISmartEnergy.IArgCurrentSummationDeliveredPayload>(base.u48, ()=>({
                name: `Current Summation Delivered`,
                description: `represents the most recent summed value of Energy, Gas, or Water delivered and consumed in the premises.`,
                id: 0x0000,
                report: true,
                read: false,
                write: false,
                require: false,
                
            })),
            CurrentSummationReceived: makeType<ZigBee.ISmartEnergy.IArgCurrentSummationReceived, ZigBee.ISmartEnergy.IArgCurrentSummationReceivedPayload>(base.u48, ()=>({
                name: `Current Summation Received`,
                description: `represents the most recent summed value of Energy, Gas, or Water generated and delivered from the premises.`,
                id: 0x0001,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            DailyConsumptionTarget: makeType<ZigBee.ISmartEnergy.IArgDailyConsumptionTarget, ZigBee.ISmartEnergy.IArgDailyConsumptionTargetPayload>(base.u24, ()=>({
                name: `Daily Consumption Target`,
                description: `is a daily target consumption amount that can be displayed to the consumer on a HAN device, with the intent that
it can be used to compare to actual daily consumption (e.g. compare to the CurrentDayConsumptionDelivered).
This may be sent from the utility to the ESI, or it may be derived. Although intended to be based on Block
Thresholds, it can be used for other targets not related to blocks. The formatting will be based on the
HistoricalConsumptionFormatting attribute.
Example: If based on a Block Threshold, the DailyConsumptionTarget could be calculated based on the
number of days specified in the Block Tariff Period and a given Block Threshold as follows:
DailyConsumptionTarget = BlockNThreshold / ((BlockPeriodDuration /60) / 24). Example: If the target is
based on a Block1Threshold of 675kWh and where 43200 BlockThresholdPeriod is the number of minutes in
the billing period (30 days), the ConsumptionDailyTarget would be 675 / ((43200 / 60) / 24) = 22.5 kWh per day`,
                id: 0x000D,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            DailyFreezeTime: makeType<ZigBee.ISmartEnergy.IArgDailyFreezeTime, ZigBee.ISmartEnergy.IArgDailyFreezeTimePayload>(base.u16, ()=>({
                name: `Daily Freeze Time`,
                description: `represents the time of day when DFTSummation.
Bits 0 to 7: Range of 0 to 0x3C representing the number of minutes past the top of the hour.
Bits 8 to 15: Range of 0 to 0x17 representing the hour of the day (in 24-hour format).`,
                id: 0x0005,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            DefaultUpdatePeriod: makeType<ZigBee.ISmartEnergy.IArgDefaultUpdatePeriod, ZigBee.ISmartEnergy.IArgDefaultUpdatePeriodPayload>(base.u8, ()=>({
                name: `Default Update Period`,
                description: `represents the interval (seconds) at which the InstantaneousDemand
attribute is updated when not in fast poll mode. InstantaneousDemand may be continuously updated as new
measurements are acquired, but at a minimum InstantaneousDemand must be updated at the
DefaultUpdatePeriod. The DefaultUpdatePeriod may apply to other attributes as defined by the device
manufacturer.`,
                id: 0x000A,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            DftSummation: makeType<ZigBee.ISmartEnergy.IArgDftSummation, ZigBee.ISmartEnergy.IArgDftSummationPayload>(base.u48, ()=>({
                name: `DFT Summation`,
                description: `represents a snapshot of attribute CurrentSummationDelivered captured at the time indicated by attribute
DailyFreezeTime`,
                id: 0x0004,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            FastPollUpdatePeriod: makeType<ZigBee.ISmartEnergy.IArgFastPollUpdatePeriod, ZigBee.ISmartEnergy.IArgFastPollUpdatePeriodPayload>(base.u8, ()=>({
                name: `Fast Poll Update Period`,
                description: `represents the interval (seconds) at which the InstantaneousDemand
attribute is updated when in fast poll mode. InstantaneousDemand may be continuously updated as new
measurements are acquired, but at a minimum, InstantaneousDemand must be updated at the
FastPollUpdatePeriod. The FastPollUpdatePeriod may apply to other attributes as defined by the device
manufacturer.`,
                id: 0x000B,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            FlowRestriction: makeType<ZigBee.ISmartEnergy.IArgFlowRestriction, ZigBee.ISmartEnergy.IArgFlowRestrictionPayload>(base.u8, ()=>({
                name: `Flow Restriction`,
                description: ``,
                id: 0x0013,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            InletTemperature: makeType<ZigBee.ISmartEnergy.IArgInletTemperature, ZigBee.ISmartEnergy.IArgInletTemperaturePayload>(base.s24, ()=>({
                name: `Inlet Temperature`,
                description: ``,
                id: 0x0017,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            IntervalReadReportingPeriod: makeType<ZigBee.ISmartEnergy.IArgIntervalReadReportingPeriod, ZigBee.ISmartEnergy.IArgIntervalReadReportingPeriodPayload>(base.u16, ()=>({
                name: `Interval Read Reporting Period`,
                description: ``,
                id: 0x0010,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            OutletTemperature: makeType<ZigBee.ISmartEnergy.IArgOutletTemperature, ZigBee.ISmartEnergy.IArgOutletTemperaturePayload>(base.s24, ()=>({
                name: `Outlet Temperature`,
                description: ``,
                id: 0x0018,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            PowerFactor: makeType<ZigBee.ISmartEnergy.IArgPowerFactor, ZigBee.ISmartEnergy.IArgPowerFactorPayload>(base.s8, ()=>({
                name: `Power Factor`,
                description: `contains the Average Power Factor ratio in 1/100ths. Valid values are 0 to 99`,
                id: 0x0006,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            PresetReadingTime: makeType<ZigBee.ISmartEnergy.IArgPresetReadingTime, ZigBee.ISmartEnergy.IArgPresetReadingTimePayload>(base.u16, ()=>({
                name: `Preset Reading Time`,
                description: ``,
                id: 0x0011,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            PreviousBlockPeriodConsumptionDelivered: makeType<ZigBee.ISmartEnergy.IArgPreviousBlockPeriodConsumptionDelivered, ZigBee.ISmartEnergy.IArgPreviousBlockPeriodConsumptionDeliveredPayload>(base.u48, ()=>({
                name: `Previous Block Period Consumption Delivered`,
                description: ``,
                id: 0x001C,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            ProfileIntervalPeriod: makeType<ZigBee.ISmartEnergy.IArgProfileIntervalPeriod, ZigBee.ISmartEnergy.IArgProfileIntervalPeriodPayload>(base.enum8, ()=>({
                name: `Profile Interval Period`,
                description: `is currently included in the Get Profile Response command payload, but does not appear in an attribute set.
This represents the duration of each interval. ProfileIntervalPeriod represents the interval or time frame used
to capture metered Energy, Gas, and Water consumption for profiling purposes.`,
                id: 0x000F,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            ReadingSnapShotTime: makeType<ZigBee.ISmartEnergy.IArgReadingSnapShotTime, ZigBee.ISmartEnergy.IArgReadingSnapShotTimePayload>(base.utc, ()=>({
                name: `Reading Snap Shot Time`,
                description: `represents the last time all of the CurrentSummationDelivered, CurrentSummationReceived,
CurrentMaxDemandDelivered, and CurrentMaxDemandReceived attributes that are supported by the device were updated`,
                id: 0x0007,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            SupplyStatus: makeType<ZigBee.ISmartEnergy.IArgSupplyStatus, ZigBee.ISmartEnergy.IArgSupplyStatusPayload>(base.enum8, ()=>({
                name: `Supply Status`,
                description: ``,
                id: 0x0014,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })),
            VolumePerReport: makeType<ZigBee.ISmartEnergy.IArgVolumePerReport, ZigBee.ISmartEnergy.IArgVolumePerReportPayload>(base.u16, ()=>({
                name: `Volume Per Report`,
                description: ``,
                id: 0x0012,
                report: false,
                read: false,
                write: false,
                require: false,
                
            })), },
        Metering: {
            ID: 0x0702,
            Name: `Metering`,
            Desc: `provides a mechanism to retrieve usage information from Electric, Gas, Water, and
potentially Thermal metering devices.`,
            
            
            Server: {
                Attribute: {},
                Command: {},
            },
            Client: {
                Attribute: {},
                Command: {}
            },
        }
    };
    
    ZigBee.SmartEnergy.Metering.Server.Attribute = { 
        0x0000: ZigBee.SmartEnergy.Types.CurrentSummationDelivered,
        0x0001: ZigBee.SmartEnergy.Types.CurrentSummationReceived,
        0x0002: ZigBee.SmartEnergy.Types.CurrentMaxDemandDelivered,
        0x0003: ZigBee.SmartEnergy.Types.CurrentMaxDemandReceived,
        0x0004: ZigBee.SmartEnergy.Types.DftSummation,
        0x0005: ZigBee.SmartEnergy.Types.DailyFreezeTime,
        0x0006: ZigBee.SmartEnergy.Types.PowerFactor,
        0x0007: ZigBee.SmartEnergy.Types.ReadingSnapShotTime,
        0x0008: ZigBee.SmartEnergy.Types.CurrentMaxDemandDeliveredTime,
        0x0009: ZigBee.SmartEnergy.Types.CurrentMaxDemandReceivedTime,
        0x000A: ZigBee.SmartEnergy.Types.DefaultUpdatePeriod,
        0x000B: ZigBee.SmartEnergy.Types.FastPollUpdatePeriod,
        0x000C: ZigBee.SmartEnergy.Types.CurrentBlockPeriodConsumptionDelivered,
        0x000D: ZigBee.SmartEnergy.Types.DailyConsumptionTarget,
        0x000E: ZigBee.SmartEnergy.Types.CurrentBlock,
        0x000F: ZigBee.SmartEnergy.Types.ProfileIntervalPeriod,
        0x0010: ZigBee.SmartEnergy.Types.IntervalReadReportingPeriod,
        0x0011: ZigBee.SmartEnergy.Types.PresetReadingTime,
        0x0012: ZigBee.SmartEnergy.Types.VolumePerReport,
        0x0013: ZigBee.SmartEnergy.Types.FlowRestriction,
        0x0014: ZigBee.SmartEnergy.Types.SupplyStatus,
        0x0015: ZigBee.SmartEnergy.Types.CurrentInletEnergyCarrierSummation,
        0x0016: ZigBee.SmartEnergy.Types.CurrentOutletEnergyCarrierSummation,
        0x0017: ZigBee.SmartEnergy.Types.InletTemperature,
        0x0018: ZigBee.SmartEnergy.Types.OutletTemperature,
        0x0019: ZigBee.SmartEnergy.Types.ControlTemperature,
        0x001A: ZigBee.SmartEnergy.Types.CurrentInletEnergyCarrierDemand,
        0x001B: ZigBee.SmartEnergy.Types.CurrentOutletEnergyCarrierDemand,
        0x001C: ZigBee.SmartEnergy.Types.PreviousBlockPeriodConsumptionDelivered,
    };
    ZigBee.SmartEnergy.Metering.Client.Attribute = { 
    };
    ZigBee.SmartEnergy.Metering.Server.Command = { 
    };
    ZigBee.SmartEnergy.Metering.Client.Command = { 
    };

    export const cluster: {[id: number]: ICluster} = { 
        0x0101: ZigBee.Closures.DoorLock,
        0x0102: ZigBee.Closures.WindowCovering,
        0x0000: ZigBee.General.Basic,
        0x0001: ZigBee.General.PowerConfiguration,
        0x0002: ZigBee.General.DeviceTemperatureConfiguration,
        0x0003: ZigBee.General.Identify,
        0x0004: ZigBee.General.Groups,
        0x0005: ZigBee.General.Scenes,
        0x0006: ZigBee.General.OnOff,
        0x0007: ZigBee.General.OnOffSwitchConfiguration,
        0x0008: ZigBee.General.LevelControl,
        0x0009: ZigBee.General.Alarms,
        0x000A: ZigBee.General.Time,
        0x000B: ZigBee.General.Location,
        0x000C: ZigBee.General.AnalogInput,
        0x000D: ZigBee.General.AnalogOutput,
        0x000E: ZigBee.General.AnalogValue,
        0x000F: ZigBee.General.BinaryInput,
        0x0010: ZigBee.General.BinaryOutput,
        0x0011: ZigBee.General.BinaryValue,
        0x0012: ZigBee.General.MultistateInput,
        0x0013: ZigBee.General.MultistateOutput,
        0x0014: ZigBee.General.MultistateValue,
        0x0020: ZigBee.General.PollControl,
        0x0B01: ZigBee.General.MeterIdentification,
        0x0B05: ZigBee.General.Diagnostics,
        0x0200: ZigBee.Hvac.PumpConfigurationAndControl,
        0x0201: ZigBee.Hvac.Thermostat,
        0x0202: ZigBee.Hvac.FanControl,
        0x0203: ZigBee.Hvac.DehumidificationControl,
        0x0204: ZigBee.Hvac.ThermostatUiConfiguration,
        0x0500: ZigBee.SecurityAndSafety.IasZone,
        0xFC7E: ZigBee.Ikea.IkeaAirQuality,
        0x0300: ZigBee.Lighting.ColorControl,
        0x0301: ZigBee.Lighting.BallastConfiguration,
        0x0400: ZigBee.MeasurementAndSensing.IlluminanceMeasurement,
        0x0401: ZigBee.MeasurementAndSensing.IlluminanceLevelSensing,
        0x0402: ZigBee.MeasurementAndSensing.TemperatureMeasurement,
        0x0403: ZigBee.MeasurementAndSensing.PressureMeasurement,
        0x0404: ZigBee.MeasurementAndSensing.FlowMeasurement,
        0x0405: ZigBee.MeasurementAndSensing.RelativeHumidityMeasurement,
        0x0406: ZigBee.MeasurementAndSensing.OccupancySensing,
        0x042a: ZigBee.MeasurementAndSensing.Pm25Measurement,
        0x0B04: ZigBee.MeasurementAndSensing.ElectricalMeasurement,
        0x040c: ZigBee.MeasurementAndSensing.CarbonMonoxideMeasurement,
        0x040d: ZigBee.MeasurementAndSensing.CarbonDioxideMeasurement,
        0x040e: ZigBee.MeasurementAndSensing.EthyleneMeasurement,
        0x040f: ZigBee.MeasurementAndSensing.EthyleneOxideMeasurement,
        0x0410: ZigBee.MeasurementAndSensing.HydrogenMeasurement,
        0x0411: ZigBee.MeasurementAndSensing.HydrogenSulfideMeasurement,
        0x0412: ZigBee.MeasurementAndSensing.NitricOxideMeasurement,
        0x0413: ZigBee.MeasurementAndSensing.NitrogenDioxideMeasurement,
        0x0414: ZigBee.MeasurementAndSensing.OxygenMeasurement,
        0x0415: ZigBee.MeasurementAndSensing.OzoneMeasurement,
        0x0416: ZigBee.MeasurementAndSensing.SulfurDioxideMeasurement,
        0x0417: ZigBee.MeasurementAndSensing.DissolvedOxygenMeasurement,
        0x0418: ZigBee.MeasurementAndSensing.BromateMeasurement,
        0x0419: ZigBee.MeasurementAndSensing.ChloraminesMeasurement,
        0x041a: ZigBee.MeasurementAndSensing.ChlorineMeasurement,
        0x041b: ZigBee.MeasurementAndSensing.FecalColiformEColiMeasurement,
        0x041c: ZigBee.MeasurementAndSensing.FluorideMeasurement,
        0x041d: ZigBee.MeasurementAndSensing.HaloaceticAcidsMeasurement,
        0x041e: ZigBee.MeasurementAndSensing.TotalTrihalomethanesMeasurement,
        0x041f: ZigBee.MeasurementAndSensing.TotalColiformBacteriaMeasurement,
        0x0420: ZigBee.MeasurementAndSensing.TurbidityMeasurement,
        0x0421: ZigBee.MeasurementAndSensing.CopperMeasurement,
        0x0422: ZigBee.MeasurementAndSensing.LeadMeasurement,
        0x0423: ZigBee.MeasurementAndSensing.ManganeseMeasurement,
        0x0424: ZigBee.MeasurementAndSensing.SulfateMeasurement,
        0x0425: ZigBee.MeasurementAndSensing.BromodichloromethaneMeasurement,
        0x0426: ZigBee.MeasurementAndSensing.BromoformMeasurement,
        0x0427: ZigBee.MeasurementAndSensing.ChlorodibromomethaneMeasurement,
        0x0428: ZigBee.MeasurementAndSensing.ChloroformMeasurement,
        0x0429: ZigBee.MeasurementAndSensing.SodiumMeasurement,
        0x042b: ZigBee.MeasurementAndSensing.FormaldehydeMeasurement,
        0x0019: ZigBee.Otau.Otau,
        0x0702: ZigBee.SmartEnergy.Metering,
    };

    export namespace IZDP {
        // noinspection ES6UnusedImports
        import ICommand = ZigBee.ICommand;
        // noinspection ES6UnusedImports
        import IArgument = ZigBee.IArgument;
        // noinspection ES6UnusedImports
        import ValueType = ZigBee.ValueType;

        
            export type ICmdNwkAddressRequestPayload = { IeeeAddress?: IArgIeeeAddressPayload, RequestType?: IArgRequestTypePayload, StartIndex?: IArgStartIndexPayload, }
            export interface ICmdNwkAddressRequest extends ICommand { value: ICmdNwkAddressRequestPayload }
            export type ICmdNwkAddressResponsePayload = { Status?: IArgStatusPayload, IeeeAddress?: IArgIeeeAddressPayload, NwkAddress?: IArgNwkAddressPayload, StartIndex?: IArgStartIndexPayload, AssociatedDevices?: IArgAssociatedDevicesPayload, }
            export interface ICmdNwkAddressResponse extends ICommand { value: ICmdNwkAddressResponsePayload }
            export type ICmdIeeeAddressRequestPayload = { NwkAddress?: IArgNwkAddressPayload, RequestType?: IArgRequestTypePayload, StartIndex?: IArgStartIndexPayload, }
            export interface ICmdIeeeAddressRequest extends ICommand { value: ICmdIeeeAddressRequestPayload }
            export type ICmdIeeeAddressResponsePayload = { Status?: IArgStatusPayload, IeeeAddress?: IArgIeeeAddressPayload, NwkAddress?: IArgNwkAddressPayload, StartIndex?: IArgStartIndexPayload, AssociatedDevices?: IArgAssociatedDevicesPayload, }
            export interface ICmdIeeeAddressResponse extends ICommand { value: ICmdIeeeAddressResponsePayload }
            export type ICmdNodeDescRequestPayload = { NwkAddress?: IArgNwkAddressPayload, }
            export interface ICmdNodeDescRequest extends ICommand { value: ICmdNodeDescRequestPayload }
            export type ICmdNodeDescResponsePayload = { Status?: IArgStatusPayload, NwkAddress?: IArgNwkAddressPayload, NodeDescriptor?: IArgNodeDescriptorPayload, }
            export interface ICmdNodeDescResponse extends ICommand { value: ICmdNodeDescResponsePayload }
            export type ICmdPowerDescRequestPayload = { NwkAddress?: IArgNwkAddressPayload, }
            export interface ICmdPowerDescRequest extends ICommand { value: ICmdPowerDescRequestPayload }
            export type ICmdPowerDescResponsePayload = { Status?: IArgStatusPayload, NwkAddress?: IArgNwkAddressPayload, PowerDescriptor?: IArgPowerDescriptorPayload, }
            export interface ICmdPowerDescResponse extends ICommand { value: ICmdPowerDescResponsePayload }
            export type ICmdSimpleDescRequestPayload = { NwkAddress?: IArgNwkAddressPayload, Endpoint?: IArgEndpointPayload, }
            export interface ICmdSimpleDescRequest extends ICommand { value: ICmdSimpleDescRequestPayload }
            export type ICmdSimpleDescResponsePayload = { Status?: IArgStatusPayload, NwkAddress?: IArgNwkAddressPayload, SimpleDescriptor?: IArgSimpleDescriptorPayload, }
            export interface ICmdSimpleDescResponse extends ICommand { value: ICmdSimpleDescResponsePayload }
            export type ICmdActiveEndpointRequestPayload = { NwkAddress?: IArgNwkAddressPayload, }
            export interface ICmdActiveEndpointRequest extends ICommand { value: ICmdActiveEndpointRequestPayload }
            export type ICmdActiveEndpointResponsePayload = { Status?: IArgStatusPayload, NwkAddress?: IArgNwkAddressPayload, EndpointList?: IArgEndpointListPayload, }
            export interface ICmdActiveEndpointResponse extends ICommand { value: ICmdActiveEndpointResponsePayload }
            export type ICmdMatchDescRequestPayload = { NwkAddress?: IArgNwkAddressPayload, ProfileId?: IArgProfileIdPayload, InClusterList?: IArgInClusterListPayload, OutClusterList?: IArgOutClusterListPayload, }
            export interface ICmdMatchDescRequest extends ICommand { value: ICmdMatchDescRequestPayload }
            export type ICmdMatchDescResponsePayload = { Status?: IArgStatusPayload, NwkAddress?: IArgNwkAddressPayload, EndpointList?: IArgEndpointListPayload, }
            export interface ICmdMatchDescResponse extends ICommand { value: ICmdMatchDescResponsePayload }
            export type ICmdComplexDescRequestPayload = { NwkAddress?: IArgNwkAddressPayload, }
            export interface ICmdComplexDescRequest extends ICommand { value: ICmdComplexDescRequestPayload }
            export type ICmdComplexDescResponsePayload = { Status?: IArgStatusPayload, NwkAddress?: IArgNwkAddressPayload, ComplexDescriptor?: IArgComplexDescriptorPayload, }
            export interface ICmdComplexDescResponse extends ICommand { value: ICmdComplexDescResponsePayload }
            export type ICmdUserDescRequestPayload = { NwkAddress?: IArgNwkAddressPayload, }
            export interface ICmdUserDescRequest extends ICommand { value: ICmdUserDescRequestPayload }
            export type ICmdUserDescResponsePayload = { Status?: IArgStatusPayload, NwkAddress?: IArgNwkAddressPayload, UserDescriptor?: IArgUserDescriptorPayload, }
            export interface ICmdUserDescResponse extends ICommand { value: ICmdUserDescResponsePayload }
            export type ICmdDiscoveryCacheRequestPayload = { NwkAddress?: IArgNwkAddressPayload, IeeeAddress?: IArgIeeeAddressPayload, }
            export interface ICmdDiscoveryCacheRequest extends ICommand { value: ICmdDiscoveryCacheRequestPayload }
            export type ICmdDeviceAnnouncePayload = { NwkAddress?: IArgNwkAddressPayload, IeeeAddress?: IArgIeeeAddressPayload, Capability?: IArgCapabilityPayload, }
            export interface ICmdDeviceAnnounce extends ICommand { value: ICmdDeviceAnnouncePayload }
            export type ICmdUserDescSetRequestPayload = { NwkAddress?: IArgNwkAddressPayload, UserDescriptor?: IArgUserDescriptorPayload, }
            export interface ICmdUserDescSetRequest extends ICommand { value: ICmdUserDescSetRequestPayload }
            export type ICmdSystemServerDiscoverRequestPayload = { ServerMask?: IArgServerMaskPayload, }
            export interface ICmdSystemServerDiscoverRequest extends ICommand { value: ICmdSystemServerDiscoverRequestPayload }
            export type ICmdDiscoveryStoreRequestPayload = { NwkAddress?: IArgNwkAddressPayload, IeeeAddress?: IArgIeeeAddressPayload, NodeDescSize?: IArgNodeDescSizePayload, PowerDescSize?: IArgPowerDescSizePayload, ActiveEndpointSize?: IArgActiveEndpointSizePayload, SimpleDescSizeList?: IArgSimpleDescSizeListPayload, }
            export interface ICmdDiscoveryStoreRequest extends ICommand { value: ICmdDiscoveryStoreRequestPayload }
            export type ICmdNodeDescStoreRequestPayload = { NwkAddress?: IArgNwkAddressPayload, IeeeAddress?: IArgIeeeAddressPayload, Capability?: IArgCapabilityPayload, ManufacturerCode?: IArgManufacturerCodePayload, MaxBufferSize?: IArgMaxBufferSizePayload, MaxRxSize?: IArgMaxRxSizePayload, ServerMask?: IArgServerMaskPayload, MaxTxSize?: IArgMaxTxSizePayload, }
            export interface ICmdNodeDescStoreRequest extends ICommand { value: ICmdNodeDescStoreRequestPayload }
            export type ICmdPowerDescStoreRequestPayload = { NwkAddress?: IArgNwkAddressPayload, IeeeAddress?: IArgIeeeAddressPayload, PowerMode?: IArgPowerModePayload, PowerSource?: IArgPowerSourcePayload, }
            export interface ICmdPowerDescStoreRequest extends ICommand { value: ICmdPowerDescStoreRequestPayload }
            export type ICmdActiveEndpointStoreRequestPayload = { NwkAddress?: IArgNwkAddressPayload, IeeeAddress?: IArgIeeeAddressPayload, EndpointList?: IArgEndpointListPayload, }
            export interface ICmdActiveEndpointStoreRequest extends ICommand { value: ICmdActiveEndpointStoreRequestPayload }
            export type ICmdSimpleDescStoreRequestPayload = { NwkAddress?: IArgNwkAddressPayload, IeeeAddress?: IArgIeeeAddressPayload, SimpleDescriptor?: IArgSimpleDescriptorPayload, }
            export interface ICmdSimpleDescStoreRequest extends ICommand { value: ICmdSimpleDescStoreRequestPayload }
            export type ICmdRemoveNodeCacheRequestPayload = { NwkAddress?: IArgNwkAddressPayload, IeeeAddress?: IArgIeeeAddressPayload, }
            export interface ICmdRemoveNodeCacheRequest extends ICommand { value: ICmdRemoveNodeCacheRequestPayload }
            export type ICmdFindNodeCacheRequestPayload = { NwkAddress?: IArgNwkAddressPayload, IeeeAddress?: IArgIeeeAddressPayload, }
            export interface ICmdFindNodeCacheRequest extends ICommand { value: ICmdFindNodeCacheRequestPayload }
            export type ICmdExtendedSimpleDescRequestPayload = { NwkAddress?: IArgNwkAddressPayload, Endpoint?: IArgEndpointPayload, StartIndex?: IArgStartIndexPayload, }
            export interface ICmdExtendedSimpleDescRequest extends ICommand { value: ICmdExtendedSimpleDescRequestPayload }
            export type ICmdExtendedActiveEndpointRequestPayload = { NwkAddress?: IArgNwkAddressPayload, StartIndex?: IArgStartIndexPayload, }
            export interface ICmdExtendedActiveEndpointRequest extends ICommand { value: ICmdExtendedActiveEndpointRequestPayload }
            export type ICmdEndDeviceBindRequestPayload = { BindingTarget?: IArgBindingTargetPayload, SourceAddress?: IArgSourceAddressPayload, SourceEndpoint?: IArgSourceEndpointPayload, ProfileId?: IArgProfileIdPayload, InClusterList?: IArgInClusterListPayload, OutClusterList?: IArgOutClusterListPayload, }
            export interface ICmdEndDeviceBindRequest extends ICommand { value: ICmdEndDeviceBindRequestPayload }
            export type ICmdBindRequestPayload = { BindingTableEntry?: IArgBindingTableEntryPayload, }
            export interface ICmdBindRequest extends ICommand { value: ICmdBindRequestPayload }
            export type ICmdBindResponsePayload = { Status?: IArgStatusPayload, }
            export interface ICmdBindResponse extends ICommand { value: ICmdBindResponsePayload }
            export type ICmdMgmtBindRequestPayload = { StartIndex?: IArgStartIndexPayload, }
            export interface ICmdMgmtBindRequest extends ICommand { value: ICmdMgmtBindRequestPayload }
            export type ICmdMgmtBindResponsePayload = { Status?: IArgStatusPayload, TotalEntries?: IArgTotalEntriesPayload, StartIndex?: IArgStartIndexPayload, BindingTable?: IArgBindingTablePayload, }
            export interface ICmdMgmtBindResponse extends ICommand { value: ICmdMgmtBindResponsePayload }
            export type ICmdUnbindRequestPayload = { BindingTableEntry?: IArgBindingTableEntryPayload, }
            export interface ICmdUnbindRequest extends ICommand { value: ICmdUnbindRequestPayload }
            export type ICmdUnbindResponsePayload = { Status?: IArgStatusPayload, }
            export interface ICmdUnbindResponse extends ICommand { value: ICmdUnbindResponsePayload }
            export type ICmdMgmtLqiRequestPayload = { StartIndex?: IArgStartIndexPayload, }
            export interface ICmdMgmtLqiRequest extends ICommand { value: ICmdMgmtLqiRequestPayload }
            export type ICmdMgmtLqiResponsePayload = { Status?: IArgStatusPayload, TotalEntries?: IArgTotalEntriesPayload, StartIndex?: IArgStartIndexPayload, NeighborTable?: IArgNeighborTablePayload, }
            export interface ICmdMgmtLqiResponse extends ICommand { value: ICmdMgmtLqiResponsePayload }
            export type ICmdMgmtNwkUpdatePayload = { Status?: IArgStatusPayload, ScannedChannels?: IArgScannedChannelsPayload, TotalTransmissions?: IArgTotalTransmissionsPayload, TransmissionFailures?: IArgTransmissionFailuresPayload, EnergyValues?: IArgEnergyValuesPayload, }
            export interface ICmdMgmtNwkUpdate extends ICommand { value: ICmdMgmtNwkUpdatePayload }
            export type IArgApsFlagsPayload = ValueType;
            export interface IArgApsFlags extends IArgument { value: IArgApsFlagsPayload }
            export type IArgActiveEndpointListPayload = ValueType;
            export interface IArgActiveEndpointList extends IArgument { value: IArgActiveEndpointListPayload }
            export type IArgActiveEndpointSizePayload = ValueType;
            export interface IArgActiveEndpointSize extends IArgument { value: IArgActiveEndpointSizePayload }
            export type IArgAddressModePayload = ValueType;
            export interface IArgAddressMode extends IArgument { value: IArgAddressModePayload }
            export type IArgAssociatedDevicesPayload = ValueType;
            export interface IArgAssociatedDevices extends IArgument { value: IArgAssociatedDevicesPayload }
            export type IArgBindingTablePayload = ValueType;
            export interface IArgBindingTable extends IArgument { value: IArgBindingTablePayload }
            export type IArgBindingTableEntryPayload = { SourceAddress?: IArgSourceAddressPayload, SourceEndpoint?: IArgSourceEndpointPayload, ClusterId?: IArgClusterIdPayload, AddressMode?: IArgAddressModePayload, NwkAddress?: IArgNwkAddressPayload, IeeeAddress?: IArgIeeeAddressPayload, Endpoint?: IArgEndpointPayload, };
            export interface IArgBindingTableEntry extends IArgument { value: IArgBindingTableEntryPayload }
            export type IArgBindingTargetPayload = ValueType;
            export interface IArgBindingTarget extends IArgument { value: IArgBindingTargetPayload }
            export type IArgCapabilityPayload = ValueType;
            export interface IArgCapability extends IArgument { value: IArgCapabilityPayload }
            export type IArgClusterIdPayload = ValueType;
            export interface IArgClusterId extends IArgument { value: IArgClusterIdPayload }
            export type IArgComplexDescriptorPayload = ValueType;
            export interface IArgComplexDescriptor extends IArgument { value: IArgComplexDescriptorPayload }
            export type IArgComplexDescriptorAvailablePayload = ValueType;
            export interface IArgComplexDescriptorAvailable extends IArgument { value: IArgComplexDescriptorAvailablePayload }
            export type IArgDepthPayload = ValueType;
            export interface IArgDepth extends IArgument { value: IArgDepthPayload }
            export type IArgDescriptorCapabilityPayload = ValueType;
            export interface IArgDescriptorCapability extends IArgument { value: IArgDescriptorCapabilityPayload }
            export type IArgDeviceTypePayload = ValueType;
            export interface IArgDeviceType extends IArgument { value: IArgDeviceTypePayload }
            export type IArgDeviceVersionPayload = ValueType;
            export interface IArgDeviceVersion extends IArgument { value: IArgDeviceVersionPayload }
            export type IArgEndpointPayload = ValueType;
            export interface IArgEndpoint extends IArgument { value: IArgEndpointPayload }
            export type IArgEndpointListPayload = ValueType;
            export interface IArgEndpointList extends IArgument { value: IArgEndpointListPayload }
            export type IArgEnergyValuesPayload = ValueType;
            export interface IArgEnergyValues extends IArgument { value: IArgEnergyValuesPayload }
            export type IArgFrequencyBandsPayload = ValueType;
            export interface IArgFrequencyBands extends IArgument { value: IArgFrequencyBandsPayload }
            export type IArgIeeeAddressPayload = ValueType;
            export interface IArgIeeeAddress extends IArgument { value: IArgIeeeAddressPayload }
            export type IArgInClusterListPayload = ValueType;
            export interface IArgInClusterList extends IArgument { value: IArgInClusterListPayload }
            export type IArgLogicalTypePayload = ValueType;
            export interface IArgLogicalType extends IArgument { value: IArgLogicalTypePayload }
            export type IArgLqiPayload = ValueType;
            export interface IArgLqi extends IArgument { value: IArgLqiPayload }
            export type IArgMacCapabilityFlagsPayload = ValueType;
            export interface IArgMacCapabilityFlags extends IArgument { value: IArgMacCapabilityFlagsPayload }
            export type IArgManufacturerCodePayload = ValueType;
            export interface IArgManufacturerCode extends IArgument { value: IArgManufacturerCodePayload }
            export type IArgMaxBufferSizePayload = ValueType;
            export interface IArgMaxBufferSize extends IArgument { value: IArgMaxBufferSizePayload }
            export type IArgMaxRxSizePayload = ValueType;
            export interface IArgMaxRxSize extends IArgument { value: IArgMaxRxSizePayload }
            export type IArgMaxTxSizePayload = ValueType;
            export interface IArgMaxTxSize extends IArgument { value: IArgMaxTxSizePayload }
            export type IArgNwkAddressPayload = ValueType;
            export interface IArgNwkAddress extends IArgument { value: IArgNwkAddressPayload }
            export type IArgNeighborTablePayload = ValueType;
            export interface IArgNeighborTable extends IArgument { value: IArgNeighborTablePayload }
            export type IArgNeighborTableEntryPayload = { PanId?: IArgIeeeAddressPayload, IeeeAddress?: IArgIeeeAddressPayload, NwkAddress?: IArgNwkAddressPayload, NeighborType?: IArgNeighborTypePayload, RxOnWhenIdle?: IArgRxOnWhenIdlePayload, Relationship?: IArgRelationshipPayload, PermitJoining?: IArgPermitJoiningPayload, Depth?: IArgDepthPayload, Lqi?: IArgLqiPayload, };
            export interface IArgNeighborTableEntry extends IArgument { value: IArgNeighborTableEntryPayload }
            export type IArgNeighborTypePayload = ValueType;
            export interface IArgNeighborType extends IArgument { value: IArgNeighborTypePayload }
            export type IArgNodeDescSizePayload = ValueType;
            export interface IArgNodeDescSize extends IArgument { value: IArgNodeDescSizePayload }
            export type IArgNodeDescriptorPayload = { LogicalType?: IArgLogicalTypePayload, ComplexDescriptorAvailable?: IArgComplexDescriptorAvailablePayload, UserDescriptorAvailable?: IArgUserDescriptorAvailablePayload, ApsFlags?: IArgApsFlagsPayload, FrequencyBands?: IArgFrequencyBandsPayload, MacCapabilityFlags?: IArgMacCapabilityFlagsPayload, ManufacturerCode?: IArgManufacturerCodePayload, MaxBufferSize?: IArgMaxBufferSizePayload, MaxRxSize?: IArgMaxRxSizePayload, ServerMask?: IArgServerMaskPayload, MaxTxSize?: IArgMaxTxSizePayload, DescriptorCapability?: IArgDescriptorCapabilityPayload, };
            export interface IArgNodeDescriptor extends IArgument { value: IArgNodeDescriptorPayload }
            export type IArgOutClusterListPayload = ValueType;
            export interface IArgOutClusterList extends IArgument { value: IArgOutClusterListPayload }
            export type IArgPermitJoiningPayload = ValueType;
            export interface IArgPermitJoining extends IArgument { value: IArgPermitJoiningPayload }
            export type IArgPowerDescSizePayload = ValueType;
            export interface IArgPowerDescSize extends IArgument { value: IArgPowerDescSizePayload }
            export type IArgPowerDescriptorPayload = { PowerMode?: IArgPowerModePayload, ActivePowerSource?: IArgPowerSourcePayload, CurrentPowerSource?: IArgPowerSourcePayload, PowerLevel?: IArgPowerLevelPayload, };
            export interface IArgPowerDescriptor extends IArgument { value: IArgPowerDescriptorPayload }
            export type IArgPowerLevelPayload = ValueType;
            export interface IArgPowerLevel extends IArgument { value: IArgPowerLevelPayload }
            export type IArgPowerModePayload = ValueType;
            export interface IArgPowerMode extends IArgument { value: IArgPowerModePayload }
            export type IArgPowerSourcePayload = ValueType;
            export interface IArgPowerSource extends IArgument { value: IArgPowerSourcePayload }
            export type IArgProfileIdPayload = ValueType;
            export interface IArgProfileId extends IArgument { value: IArgProfileIdPayload }
            export type IArgRelationshipPayload = ValueType;
            export interface IArgRelationship extends IArgument { value: IArgRelationshipPayload }
            export type IArgRequestTypePayload = ValueType;
            export interface IArgRequestType extends IArgument { value: IArgRequestTypePayload }
            export type IArgRxOnWhenIdlePayload = ValueType;
            export interface IArgRxOnWhenIdle extends IArgument { value: IArgRxOnWhenIdlePayload }
            export type IArgScannedChannelsPayload = ValueType;
            export interface IArgScannedChannels extends IArgument { value: IArgScannedChannelsPayload }
            export type IArgServerMaskPayload = ValueType;
            export interface IArgServerMask extends IArgument { value: IArgServerMaskPayload }
            export type IArgSimpleDescSizeListPayload = ValueType;
            export interface IArgSimpleDescSizeList extends IArgument { value: IArgSimpleDescSizeListPayload }
            export type IArgSimpleDescriptorPayload = { Endpoint?: IArgEndpointPayload, ProfileId?: IArgProfileIdPayload, DeviceType?: IArgDeviceTypePayload, DeviceVersion?: IArgDeviceVersionPayload, InClusterList?: IArgInClusterListPayload, OutClusterList?: IArgOutClusterListPayload, };
            export interface IArgSimpleDescriptor extends IArgument { value: IArgSimpleDescriptorPayload }
            export type IArgSourceAddressPayload = ValueType;
            export interface IArgSourceAddress extends IArgument { value: IArgSourceAddressPayload }
            export type IArgSourceEndpointPayload = ValueType;
            export interface IArgSourceEndpoint extends IArgument { value: IArgSourceEndpointPayload }
            export type IArgStartIndexPayload = ValueType;
            export interface IArgStartIndex extends IArgument { value: IArgStartIndexPayload }
            export type IArgStatusPayload = ValueType;
            export interface IArgStatus extends IArgument { value: IArgStatusPayload }
            export type IArgTotalEntriesPayload = ValueType;
            export interface IArgTotalEntries extends IArgument { value: IArgTotalEntriesPayload }
            export type IArgTotalTransmissionsPayload = ValueType;
            export interface IArgTotalTransmissions extends IArgument { value: IArgTotalTransmissionsPayload }
            export type IArgTransmissionFailuresPayload = ValueType;
            export interface IArgTransmissionFailures extends IArgument { value: IArgTransmissionFailuresPayload }
            export type IArgUnknownU8Payload = ValueType;
            export interface IArgUnknownU8 extends IArgument { value: IArgUnknownU8Payload }
            export type IArgUserDescriptorPayload = ValueType;
            export interface IArgUserDescriptor extends IArgument { value: IArgUserDescriptorPayload }
            export type IArgUserDescriptorAvailablePayload = ValueType;
            export interface IArgUserDescriptorAvailable extends IArgument { value: IArgUserDescriptorAvailablePayload }
    }
    
    export namespace IClosures {
        // noinspection ES6UnusedImports
        import IArgument = ZigBee.IArgument;
        // noinspection ES6UnusedImports
        import IAttribute = ZigBee.IAttribute;
        // noinspection ES6UnusedImports
        import ValueType = ZigBee.ValueType;
        
        export namespace DoorLock {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace WindowCovering {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
            export type ICmdUpOpenPayload = { }
            export interface ICmdUpOpen extends ICommand { value: ICmdUpOpenPayload }
            export type ICmdDownClosePayload = { }
            export interface ICmdDownClose extends ICommand { value: ICmdDownClosePayload }
            export type ICmdStopPayload = { }
            export interface ICmdStop extends ICommand { value: ICmdStopPayload }
            export type ICmdGoToLiftValuePayload = { Position?: IArgPositionPayload, }
            export interface ICmdGoToLiftValue extends ICommand { value: ICmdGoToLiftValuePayload }
            export type ICmdGoToLiftPercentagePayload = { Percentage?: IArgPercentagePayload, }
            export interface ICmdGoToLiftPercentage extends ICommand { value: ICmdGoToLiftPercentagePayload }
            export type ICmdGoToTiltValuePayload = { Position?: IArgPositionPayload, }
            export interface ICmdGoToTiltValue extends ICommand { value: ICmdGoToTiltValuePayload }
            export type ICmdGoToTiltPercentagePayload = { Percentage?: IArgPercentagePayload, }
            export interface ICmdGoToTiltPercentage extends ICommand { value: ICmdGoToTiltPercentagePayload }
        }

            export type IArgConfigStatusPayload = ValueType;
            export interface IArgConfigStatus extends IAttribute { value: IArgConfigStatusPayload }
            export type IArgLiftAccelerationTimePayload = ValueType;
            export interface IArgLiftAccelerationTime extends IAttribute { value: IArgLiftAccelerationTimePayload }
            export type IArgLiftCurrentPositionPayload = ValueType;
            export interface IArgLiftCurrentPosition extends IAttribute { value: IArgLiftCurrentPositionPayload }
            export type IArgLiftDecelerationTimePayload = ValueType;
            export interface IArgLiftDecelerationTime extends IAttribute { value: IArgLiftDecelerationTimePayload }
            export type IArgLiftInstalledClosedLimitPayload = ValueType;
            export interface IArgLiftInstalledClosedLimit extends IAttribute { value: IArgLiftInstalledClosedLimitPayload }
            export type IArgLiftInstalledOpenLimitPayload = ValueType;
            export interface IArgLiftInstalledOpenLimit extends IAttribute { value: IArgLiftInstalledOpenLimitPayload }
            export type IArgLiftIntermediateSetpointsPayload = ValueType;
            export interface IArgLiftIntermediateSetpoints extends IAttribute { value: IArgLiftIntermediateSetpointsPayload }
            export type IArgLiftNumberOfActuationsPayload = ValueType;
            export interface IArgLiftNumberOfActuations extends IAttribute { value: IArgLiftNumberOfActuationsPayload }
            export type IArgLiftPhysicalClosedLimitPayload = ValueType;
            export interface IArgLiftPhysicalClosedLimit extends IAttribute { value: IArgLiftPhysicalClosedLimitPayload }
            export type IArgLiftVelocityPayload = ValueType;
            export interface IArgLiftVelocity extends IAttribute { value: IArgLiftVelocityPayload }
            export type IArgLiftCurrentPositionPercentagePayload = ValueType;
            export interface IArgLiftCurrentPositionPercentage extends IAttribute { value: IArgLiftCurrentPositionPercentagePayload }
            export type IArgPercentagePayload = ValueType;
            export interface IArgPercentage extends IArgument { value: IArgPercentagePayload }
            export type IArgPositionPayload = ValueType;
            export interface IArgPosition extends IArgument { value: IArgPositionPayload }
            export type IArgTiltCurrentPositionPayload = ValueType;
            export interface IArgTiltCurrentPosition extends IAttribute { value: IArgTiltCurrentPositionPayload }
            export type IArgTiltIntermediateSetpointsPayload = ValueType;
            export interface IArgTiltIntermediateSetpoints extends IAttribute { value: IArgTiltIntermediateSetpointsPayload }
            export type IArgTiltNumberOfActuationsPayload = ValueType;
            export interface IArgTiltNumberOfActuations extends IAttribute { value: IArgTiltNumberOfActuationsPayload }
            export type IArgTiltPhysicalClosedLimitPayload = ValueType;
            export interface IArgTiltPhysicalClosedLimit extends IAttribute { value: IArgTiltPhysicalClosedLimitPayload }
            export type IArgTiltAInstalledOpenLimitPayload = ValueType;
            export interface IArgTiltAInstalledOpenLimit extends IAttribute { value: IArgTiltAInstalledOpenLimitPayload }
            export type IArgTiltBInstalledOpenLimitPayload = ValueType;
            export interface IArgTiltBInstalledOpenLimit extends IAttribute { value: IArgTiltBInstalledOpenLimitPayload }
            export type IArgTiltCurrentPositionPercentagePayload = ValueType;
            export interface IArgTiltCurrentPositionPercentage extends IAttribute { value: IArgTiltCurrentPositionPercentagePayload }
            export type IArgWindowCoveringModePayload = ValueType;
            export interface IArgWindowCoveringMode extends IAttribute { value: IArgWindowCoveringModePayload }
            export type IArgWindowCoveringTypePayload = ValueType;
            export interface IArgWindowCoveringType extends IAttribute { value: IArgWindowCoveringTypePayload }    }

    export namespace IGeneral {
        // noinspection ES6UnusedImports
        import IArgument = ZigBee.IArgument;
        // noinspection ES6UnusedImports
        import IAttribute = ZigBee.IAttribute;
        // noinspection ES6UnusedImports
        import ValueType = ZigBee.ValueType;
        
        export namespace Basic {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
            export type ICmdResetToFactoryDefaultsPayload = { }
            export interface ICmdResetToFactoryDefaults extends ICommand { value: ICmdResetToFactoryDefaultsPayload }
        }

        export namespace PowerConfiguration {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace DeviceTemperatureConfiguration {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace Identify {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
            export type ICmdIdentifyPayload = { IdentifyTime?: IArgIdentifyTimePayload, }
            export interface ICmdIdentify extends ICommand { value: ICmdIdentifyPayload }
            export type ICmdIdentifyQueryPayload = { }
            export interface ICmdIdentifyQuery extends ICommand { value: ICmdIdentifyQueryPayload }
            export type ICmdTriggerEffectPayload = { IdentifyEffect?: IArgIdentifyEffectPayload, IdentifyEffectVariant?: IArgIdentifyEffectVariantPayload, }
            export interface ICmdTriggerEffect extends ICommand { value: ICmdTriggerEffectPayload }
            export type ICmdIdentifyQueryResponsePayload = { IdentifyTimeout?: IArgIdentifyTimeoutPayload, }
            export interface ICmdIdentifyQueryResponse extends ICommand { value: ICmdIdentifyQueryResponsePayload }
        }

        export namespace Groups {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
            export type ICmdAddGroupPayload = { GroupId?: IArgGroupIdPayload, GroupName?: IArgGroupNamePayload, }
            export interface ICmdAddGroup extends ICommand { value: ICmdAddGroupPayload }
            export type ICmdViewGroupPayload = { GroupId?: IArgGroupIdPayload, }
            export interface ICmdViewGroup extends ICommand { value: ICmdViewGroupPayload }
            export type ICmdGetGroupMembershipPayload = { GroupList?: IArgGroupListPayload, }
            export interface ICmdGetGroupMembership extends ICommand { value: ICmdGetGroupMembershipPayload }
            export type ICmdRemoveGroupPayload = { GroupList?: IArgGroupListPayload, }
            export interface ICmdRemoveGroup extends ICommand { value: ICmdRemoveGroupPayload }
            export type ICmdRemoveAllGroupsPayload = { }
            export interface ICmdRemoveAllGroups extends ICommand { value: ICmdRemoveAllGroupsPayload }
            export type ICmdAddGroupIfIdentifyingPayload = { GroupId?: IArgGroupIdPayload, GroupName?: IArgGroupNamePayload, }
            export interface ICmdAddGroupIfIdentifying extends ICommand { value: ICmdAddGroupIfIdentifyingPayload }
            export type ICmdAddGroupResponsePayload = { Status?: IArgStatusPayload, GroupId?: IArgGroupIdPayload, }
            export interface ICmdAddGroupResponse extends ICommand { value: ICmdAddGroupResponsePayload }
            export type ICmdViewGroupResponsePayload = { Status?: IArgStatusPayload, GroupId?: IArgGroupIdPayload, GroupName?: IArgGroupNamePayload, }
            export interface ICmdViewGroupResponse extends ICommand { value: ICmdViewGroupResponsePayload }
            export type ICmdGetGroupMembershipResponsePayload = { GroupCapacity?: IArgGroupCapacityPayload, GroupList?: IArgGroupListPayload, }
            export interface ICmdGetGroupMembershipResponse extends ICommand { value: ICmdGetGroupMembershipResponsePayload }
            export type ICmdRemoveGroupResponsePayload = { Status?: IArgStatusPayload, GroupId?: IArgGroupIdPayload, }
            export interface ICmdRemoveGroupResponse extends ICommand { value: ICmdRemoveGroupResponsePayload }
        }

        export namespace Scenes {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
            export type ICmdAddScenePayload = { GroupId?: IArgGroupIdPayload, SceneId?: IArgSceneIdPayload, TransitionTimeSec?: IArgTransitionTimeSecPayload, SceneName?: IArgSceneNamePayload, SceneExtensions?: IArgSceneExtensionsPayload, }
            export interface ICmdAddScene extends ICommand { value: ICmdAddScenePayload }
            export type ICmdViewScenePayload = { GroupId?: IArgGroupIdPayload, SceneId?: IArgSceneIdPayload, }
            export interface ICmdViewScene extends ICommand { value: ICmdViewScenePayload }
            export type ICmdRemoveScenePayload = { GroupId?: IArgGroupIdPayload, SceneId?: IArgSceneIdPayload, }
            export interface ICmdRemoveScene extends ICommand { value: ICmdRemoveScenePayload }
            export type ICmdRemoveAllScenesPayload = { GroupId?: IArgGroupIdPayload, }
            export interface ICmdRemoveAllScenes extends ICommand { value: ICmdRemoveAllScenesPayload }
            export type ICmdStoreScenePayload = { GroupId?: IArgGroupIdPayload, SceneId?: IArgSceneIdPayload, }
            export interface ICmdStoreScene extends ICommand { value: ICmdStoreScenePayload }
            export type ICmdRecallScenePayload = { GroupId?: IArgGroupIdPayload, SceneId?: IArgSceneIdPayload, }
            export interface ICmdRecallScene extends ICommand { value: ICmdRecallScenePayload }
            export type ICmdGetSceneMembershipPayload = { GroupId?: IArgGroupIdPayload, }
            export interface ICmdGetSceneMembership extends ICommand { value: ICmdGetSceneMembershipPayload }
            export type ICmdEnhancedAddScenePayload = { GroupId?: IArgGroupIdPayload, SceneId?: IArgSceneIdPayload, TransitionTime?: IArgTransitionTimePayload, SceneName?: IArgSceneNamePayload, SceneExtensions?: IArgSceneExtensionsPayload, }
            export interface ICmdEnhancedAddScene extends ICommand { value: ICmdEnhancedAddScenePayload }
            export type ICmdEnhancedViewScenePayload = { GroupId?: IArgGroupIdPayload, SceneId?: IArgSceneIdPayload, }
            export interface ICmdEnhancedViewScene extends ICommand { value: ICmdEnhancedViewScenePayload }
            export type ICmdCopyScenePayload = { SceneCopyMode?: IArgSceneCopyModePayload, FromGroupId?: IArgGroupIdPayload, FromSceneId?: IArgSceneIdPayload, ToGroupId?: IArgGroupIdPayload, ToSceneId?: IArgSceneIdPayload, }
            export interface ICmdCopyScene extends ICommand { value: ICmdCopyScenePayload }
            export type ICmdIkeaRemotePressPayload = { IkeaRemoteDirection?: IArgIkeaRemoteDirectionPayload, }
            export interface ICmdIkeaRemotePress extends ICommand { value: ICmdIkeaRemotePressPayload }
            export type ICmdIkeaRemoteLongpressStartPayload = { IkeaRemoteDirection?: IArgIkeaRemoteDirectionPayload, }
            export interface ICmdIkeaRemoteLongpressStart extends ICommand { value: ICmdIkeaRemoteLongpressStartPayload }
            export type ICmdIkeaRemoteLongpressStopPayload = { }
            export interface ICmdIkeaRemoteLongpressStop extends ICommand { value: ICmdIkeaRemoteLongpressStopPayload }
            export type ICmdAddSceneResponsePayload = { Status?: IArgStatusPayload, GroupId?: IArgGroupIdPayload, SceneId?: IArgSceneIdPayload, }
            export interface ICmdAddSceneResponse extends ICommand { value: ICmdAddSceneResponsePayload }
            export type ICmdViewSceneResponsePayload = { Status?: IArgStatusPayload, GroupId?: IArgGroupIdPayload, SceneId?: IArgSceneIdPayload, TransitionTimeSec?: IArgTransitionTimeSecPayload, SceneName?: IArgSceneNamePayload, SceneExtensions?: IArgSceneExtensionsPayload, }
            export interface ICmdViewSceneResponse extends ICommand { value: ICmdViewSceneResponsePayload }
            export type ICmdRemoveSceneResponsePayload = { Status?: IArgStatusPayload, GroupId?: IArgGroupIdPayload, SceneId?: IArgSceneIdPayload, }
            export interface ICmdRemoveSceneResponse extends ICommand { value: ICmdRemoveSceneResponsePayload }
            export type ICmdRemoveAllScenesResponsePayload = { Status?: IArgStatusPayload, GroupId?: IArgGroupIdPayload, }
            export interface ICmdRemoveAllScenesResponse extends ICommand { value: ICmdRemoveAllScenesResponsePayload }
            export type ICmdStoreSceneResponsePayload = { Status?: IArgStatusPayload, GroupId?: IArgGroupIdPayload, SceneId?: IArgSceneIdPayload, }
            export interface ICmdStoreSceneResponse extends ICommand { value: ICmdStoreSceneResponsePayload }
            export type ICmdGetSceneMembershipResponsePayload = { Status?: IArgStatusPayload, SceneCapacity?: IArgSceneCapacityPayload, GroupId?: IArgGroupIdPayload, SceneList?: IArgSceneListPayload, }
            export interface ICmdGetSceneMembershipResponse extends ICommand { value: ICmdGetSceneMembershipResponsePayload }
            export type ICmdEnhancedAddSceneResponsePayload = { Status?: IArgStatusPayload, GroupId?: IArgGroupIdPayload, SceneId?: IArgSceneIdPayload, }
            export interface ICmdEnhancedAddSceneResponse extends ICommand { value: ICmdEnhancedAddSceneResponsePayload }
            export type ICmdEnhancedViewSceneResponsePayload = { Status?: IArgStatusPayload, GroupId?: IArgGroupIdPayload, SceneId?: IArgSceneIdPayload, TransitionTime?: IArgTransitionTimePayload, SceneName?: IArgSceneNamePayload, SceneExtensions?: IArgSceneExtensionsPayload, }
            export interface ICmdEnhancedViewSceneResponse extends ICommand { value: ICmdEnhancedViewSceneResponsePayload }
            export type ICmdCopySceneResponsePayload = { Status?: IArgStatusPayload, FromGroupId?: IArgGroupIdPayload, FromSceneId?: IArgSceneIdPayload, }
            export interface ICmdCopySceneResponse extends ICommand { value: ICmdCopySceneResponsePayload }
        }

        export namespace OnOff {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
            export type ICmdOffPayload = { }
            export interface ICmdOff extends ICommand { value: ICmdOffPayload }
            export type ICmdOnPayload = { }
            export interface ICmdOn extends ICommand { value: ICmdOnPayload }
            export type ICmdTogglePayload = { }
            export interface ICmdToggle extends ICommand { value: ICmdTogglePayload }
            export type ICmdOffWithEffectPayload = { EffectIdentifier?: IArgEffectIdentifierPayload, EffectVariant?: IArgEffectVariantPayload, }
            export interface ICmdOffWithEffect extends ICommand { value: ICmdOffWithEffectPayload }
            export type ICmdOnWithRecallGlobalScenePayload = { }
            export interface ICmdOnWithRecallGlobalScene extends ICommand { value: ICmdOnWithRecallGlobalScenePayload }
            export type ICmdOnWithTimedOffPayload = { OnOffControl?: IArgOnOffControlPayload, OnTime?: IArgOnTimePayload, OffWaitTime?: IArgOffWaitTimePayload, }
            export interface ICmdOnWithTimedOff extends ICommand { value: ICmdOnWithTimedOffPayload }
        }

        export namespace OnOffSwitchConfiguration {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace LevelControl {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
            export type ICmdMoveToLevelPayload = { Level?: IArgLevelPayload, TransitionTime?: IArgTransitionTimePayload, }
            export interface ICmdMoveToLevel extends ICommand { value: ICmdMoveToLevelPayload }
            export type ICmdMovePayload = { LevelDirection?: IArgLevelDirectionPayload, Rate?: IArgRatePayload, }
            export interface ICmdMove extends ICommand { value: ICmdMovePayload }
            export type ICmdStepPayload = { LevelDirection?: IArgLevelDirectionPayload, StepSize?: IArgStepSizePayload, TransitionTime?: IArgTransitionTimePayload, }
            export interface ICmdStep extends ICommand { value: ICmdStepPayload }
            export type ICmdStopPayload = { }
            export interface ICmdStop extends ICommand { value: ICmdStopPayload }
            export type ICmdMoveToLevelWithOnOffPayload = { Level?: IArgLevelPayload, TransitionTime?: IArgTransitionTimePayload, }
            export interface ICmdMoveToLevelWithOnOff extends ICommand { value: ICmdMoveToLevelWithOnOffPayload }
            export type ICmdMoveWithOnOffPayload = { LevelDirection?: IArgLevelDirectionPayload, Rate?: IArgRatePayload, }
            export interface ICmdMoveWithOnOff extends ICommand { value: ICmdMoveWithOnOffPayload }
            export type ICmdStepWithOnOffPayload = { LevelDirection?: IArgLevelDirectionPayload, StepSize?: IArgStepSizePayload, TransitionTime?: IArgTransitionTimePayload, }
            export interface ICmdStepWithOnOff extends ICommand { value: ICmdStepWithOnOffPayload }
            export type ICmdStopWithOnOffPayload = { }
            export interface ICmdStopWithOnOff extends ICommand { value: ICmdStopWithOnOffPayload }
        }

        export namespace Alarms {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
            export type ICmdResetAlarmPayload = { AlarmCode?: IArgAlarmCodePayload, ClusterId?: IArgClusterIdPayload, }
            export interface ICmdResetAlarm extends ICommand { value: ICmdResetAlarmPayload }
            export type ICmdResetAllAlarmsPayload = { }
            export interface ICmdResetAllAlarms extends ICommand { value: ICmdResetAllAlarmsPayload }
            export type ICmdGetAlarmPayload = { }
            export interface ICmdGetAlarm extends ICommand { value: ICmdGetAlarmPayload }
            export type ICmdResetAlarmLogPayload = { }
            export interface ICmdResetAlarmLog extends ICommand { value: ICmdResetAlarmLogPayload }
            export type ICmdAlarmPayload = { AlarmCode?: IArgAlarmCodePayload, ClusterId?: IArgClusterIdPayload, }
            export interface ICmdAlarm extends ICommand { value: ICmdAlarmPayload }
            export type ICmdGetAlarmResponsePayload = { Status?: IArgStatusPayload, AlarmCode?: IArgAlarmCodePayload, ClusterId?: IArgClusterIdPayload, Time?: IArgTimePayload, }
            export interface ICmdGetAlarmResponse extends ICommand { value: ICmdGetAlarmResponsePayload }
        }

        export namespace Time {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace Location {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
            export type ICmdSetAbsoluteLocationPayload = { XCoordinate?: IArgXCoordinatePayload, YCoordinate?: IArgYCoordinatePayload, ZCoordinate?: IArgZCoordinatePayload, Power?: IArgPowerPayload, PathLossExponent?: IArgPathLossExponentPayload, }
            export interface ICmdSetAbsoluteLocation extends ICommand { value: ICmdSetAbsoluteLocationPayload }
            export type ICmdSetDeviceConfigurationPayload = { Power?: IArgPowerPayload, PathLossExponent?: IArgPathLossExponentPayload, CalculationPeriod?: IArgCalculationPeriodPayload, NumberRssiMeasurements?: IArgNumberRssiMeasurementsPayload, ReportingPeriod?: IArgReportingPeriodPayload, }
            export interface ICmdSetDeviceConfiguration extends ICommand { value: ICmdSetDeviceConfigurationPayload }
            export type ICmdGetDeviceConfigurationPayload = { TargetAddress?: IArgDevicePayload, }
            export interface ICmdGetDeviceConfiguration extends ICommand { value: ICmdGetDeviceConfigurationPayload }
            export type ICmdGetLocationDataPayload = { LocationFlags?: IArgLocationFlagsPayload, NumberResponses?: IArgNumberResponsesPayload, TargetAddress?: IArgDevicePayload, }
            export interface ICmdGetLocationData extends ICommand { value: ICmdGetLocationDataPayload }
            export type ICmdRssiResponsePayload = { Device?: IArgDevicePayload, XCoordinate?: IArgXCoordinatePayload, YCoordinate?: IArgYCoordinatePayload, ZCoordinate?: IArgZCoordinatePayload, Rssi?: IArgRssiPayload, NumberRssiMeasurements?: IArgNumberRssiMeasurementsPayload, }
            export interface ICmdRssiResponse extends ICommand { value: ICmdRssiResponsePayload }
            export type ICmdSendPingsPayload = { TargetAddress?: IArgDevicePayload, NumberRssiMeasurements?: IArgNumberRssiMeasurementsPayload, CalculationPeriod?: IArgCalculationPeriodPayload, }
            export interface ICmdSendPings extends ICommand { value: ICmdSendPingsPayload }
            export type ICmdAnchorNodeAnnouncePayload = { Device?: IArgDevicePayload, XCoordinate?: IArgXCoordinatePayload, YCoordinate?: IArgYCoordinatePayload, ZCoordinate?: IArgZCoordinatePayload, }
            export interface ICmdAnchorNodeAnnounce extends ICommand { value: ICmdAnchorNodeAnnouncePayload }
            export type ICmdDistanceMeasurePayload = { TargetAddress?: IArgDevicePayload, Resolution?: IArgResolutionPayload, }
            export interface ICmdDistanceMeasure extends ICommand { value: ICmdDistanceMeasurePayload }
            export type ICmdDeviceConfigurationResponsePayload = { Status?: IArgStatusPayload, Power?: IArgPowerPayload, PathLossExponent?: IArgPathLossExponentPayload, CalculationPeriod?: IArgCalculationPeriodPayload, NumberRssiMeasurements?: IArgNumberRssiMeasurementsPayload, ReportingPeriod?: IArgReportingPeriodPayload, }
            export interface ICmdDeviceConfigurationResponse extends ICommand { value: ICmdDeviceConfigurationResponsePayload }
            export type ICmdLocationDataResponsePayload = { Status?: IArgStatusPayload, LocationType?: IArgLocationTypePayload, XCoordinate?: IArgXCoordinatePayload, YCoordinate?: IArgYCoordinatePayload, ZCoordinate?: IArgZCoordinatePayload, Power?: IArgPowerPayload, PathLossExponent?: IArgPathLossExponentPayload, LocationMethod?: IArgLocationMethodPayload, QualityMeasure?: IArgQualityMeasurePayload, LocationAge?: IArgLocationAgePayload, }
            export interface ICmdLocationDataResponse extends ICommand { value: ICmdLocationDataResponsePayload }
            export type ICmdLocationDataNotificationPayload = { LocationType?: IArgLocationTypePayload, XCoordinate?: IArgXCoordinatePayload, YCoordinate?: IArgYCoordinatePayload, ZCoordinate?: IArgZCoordinatePayload, Power?: IArgPowerPayload, PathLossExponent?: IArgPathLossExponentPayload, LocationMethod?: IArgLocationMethodPayload, QualityMeasure?: IArgQualityMeasurePayload, LocationAge?: IArgLocationAgePayload, }
            export interface ICmdLocationDataNotification extends ICommand { value: ICmdLocationDataNotificationPayload }
            export type ICmdCompactLocationDataNotificationPayload = { LocationType?: IArgLocationTypePayload, XCoordinate?: IArgXCoordinatePayload, YCoordinate?: IArgYCoordinatePayload, ZCoordinate?: IArgZCoordinatePayload, QualityMeasure?: IArgQualityMeasurePayload, LocationAge?: IArgLocationAgePayload, }
            export interface ICmdCompactLocationDataNotification extends ICommand { value: ICmdCompactLocationDataNotificationPayload }
            export type ICmdRssiPingPayload = { LocationType?: IArgLocationTypePayload, }
            export interface ICmdRssiPing extends ICommand { value: ICmdRssiPingPayload }
            export type ICmdRssiRequestPayload = { }
            export interface ICmdRssiRequest extends ICommand { value: ICmdRssiRequestPayload }
            export type ICmdReportRssiMeasurementsPayload = { Device?: IArgDevicePayload, NeighborsInfoList?: IArgNeighborsInfoListPayload, }
            export interface ICmdReportRssiMeasurements extends ICommand { value: ICmdReportRssiMeasurementsPayload }
            export type ICmdRequestOwnLocationPayload = { BlindNodeAddress?: IArgDevicePayload, }
            export interface ICmdRequestOwnLocation extends ICommand { value: ICmdRequestOwnLocationPayload }
            export type ICmdDistanceMeasureResponsePayload = { TargetAddress?: IArgDevicePayload, Distance?: IArgDistancePayload, QualityIndex?: IArgQualityIndexPayload, }
            export interface ICmdDistanceMeasureResponse extends ICommand { value: ICmdDistanceMeasureResponsePayload }
        }

        export namespace AnalogInput {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace AnalogOutput {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace AnalogValue {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace BinaryInput {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace BinaryOutput {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace BinaryValue {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace MultistateInput {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace MultistateOutput {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace MultistateValue {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace PollControl {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
            export type ICmdCheckInPayload = { }
            export interface ICmdCheckIn extends ICommand { value: ICmdCheckInPayload }
        }

        export namespace MeterIdentification {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace Diagnostics {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

            export type IArgApsDecryptFailuresPayload = ValueType;
            export interface IArgApsDecryptFailures extends IAttribute { value: IArgApsDecryptFailuresPayload }
            export type IArgApsFcFailurePayload = ValueType;
            export interface IArgApsFcFailure extends IAttribute { value: IArgApsFcFailurePayload }
            export type IArgApsRxBcastPayload = ValueType;
            export interface IArgApsRxBcast extends IAttribute { value: IArgApsRxBcastPayload }
            export type IArgApsRxUcastPayload = ValueType;
            export interface IArgApsRxUcast extends IAttribute { value: IArgApsRxUcastPayload }
            export type IArgApsTxBcastPayload = ValueType;
            export interface IArgApsTxBcast extends IAttribute { value: IArgApsTxBcastPayload }
            export type IArgApsTxUcastFailPayload = ValueType;
            export interface IArgApsTxUcastFail extends IAttribute { value: IArgApsTxUcastFailPayload }
            export type IArgApsTxUcastRetryPayload = ValueType;
            export interface IArgApsTxUcastRetry extends IAttribute { value: IArgApsTxUcastRetryPayload }
            export type IArgApsTxUcastSuccessPayload = ValueType;
            export interface IArgApsTxUcastSuccess extends IAttribute { value: IArgApsTxUcastSuccessPayload }
            export type IArgApsUnauthorizedKeyPayload = ValueType;
            export interface IArgApsUnauthorizedKey extends IAttribute { value: IArgApsUnauthorizedKeyPayload }
            export type IArgAlarmCountPayload = ValueType;
            export interface IArgAlarmCount extends IAttribute { value: IArgAlarmCountPayload }
            export type IArgAlarmMaskPayload = ValueType;
            export interface IArgAlarmMask extends IAttribute { value: IArgAlarmMaskPayload }
            export type IArgAlarmCodePayload = ValueType;
            export interface IArgAlarmCode extends IArgument { value: IArgAlarmCodePayload }
            export type IArgAnalogMaxPresentValuePayload = ValueType;
            export interface IArgAnalogMaxPresentValue extends IAttribute { value: IArgAnalogMaxPresentValuePayload }
            export type IArgAnalogMinPresentValuePayload = ValueType;
            export interface IArgAnalogMinPresentValue extends IAttribute { value: IArgAnalogMinPresentValuePayload }
            export type IArgAnalogPresentValuePayload = ValueType;
            export interface IArgAnalogPresentValue extends IAttribute { value: IArgAnalogPresentValuePayload }
            export type IArgAnalogPriorityPayload = { };
            export interface IArgAnalogPriority extends IArgument { value: IArgAnalogPriorityPayload }
            export type IArgAnalogPriorityArrayPayload = ValueType;
            export interface IArgAnalogPriorityArray extends IAttribute { value: IArgAnalogPriorityArrayPayload }
            export type IArgAnalogRelinquishDefaultPayload = ValueType;
            export interface IArgAnalogRelinquishDefault extends IAttribute { value: IArgAnalogRelinquishDefaultPayload }
            export type IArgAnalogResolutionPayload = ValueType;
            export interface IArgAnalogResolution extends IAttribute { value: IArgAnalogResolutionPayload }
            export type IArgApplicationVersionPayload = ValueType;
            export interface IArgApplicationVersion extends IAttribute { value: IArgApplicationVersionPayload }
            export type IArgAvgMacRetryPerApsMsgSentPayload = ValueType;
            export interface IArgAvgMacRetryPerApsMsgSent extends IAttribute { value: IArgAvgMacRetryPerApsMsgSentPayload }
            export type IArgBatteryAlarmMaskPayload = ValueType;
            export interface IArgBatteryAlarmMask extends IAttribute { value: IArgBatteryAlarmMaskPayload }
            export type IArgBatteryAlarmStatePayload = ValueType;
            export interface IArgBatteryAlarmState extends IAttribute { value: IArgBatteryAlarmStatePayload }
            export type IArgBatteryManufacturerPayload = ValueType;
            export interface IArgBatteryManufacturer extends IAttribute { value: IArgBatteryManufacturerPayload }
            export type IArgBatteryPercentageMinThresholdPayload = ValueType;
            export interface IArgBatteryPercentageMinThreshold extends IAttribute { value: IArgBatteryPercentageMinThresholdPayload }
            export type IArgBatteryPercentageThreshold1Payload = ValueType;
            export interface IArgBatteryPercentageThreshold1 extends IAttribute { value: IArgBatteryPercentageThreshold1Payload }
            export type IArgBatteryPercentageThreshold2Payload = ValueType;
            export interface IArgBatteryPercentageThreshold2 extends IAttribute { value: IArgBatteryPercentageThreshold2Payload }
            export type IArgBatteryPercentageThreshold3Payload = ValueType;
            export interface IArgBatteryPercentageThreshold3 extends IAttribute { value: IArgBatteryPercentageThreshold3Payload }
            export type IArgBatteryQuantityPayload = ValueType;
            export interface IArgBatteryQuantity extends IAttribute { value: IArgBatteryQuantityPayload }
            export type IArgBatteryRatedVoltagePayload = ValueType;
            export interface IArgBatteryRatedVoltage extends IAttribute { value: IArgBatteryRatedVoltagePayload }
            export type IArgBatteryRemainingPayload = ValueType;
            export interface IArgBatteryRemaining extends IAttribute { value: IArgBatteryRemainingPayload }
            export type IArgBatterySizePayload = ValueType;
            export interface IArgBatterySize extends IAttribute { value: IArgBatterySizePayload }
            export type IArgBatteryVoltagePayload = ValueType;
            export interface IArgBatteryVoltage extends IAttribute { value: IArgBatteryVoltagePayload }
            export type IArgBatteryVoltageMinThresholdPayload = ValueType;
            export interface IArgBatteryVoltageMinThreshold extends IAttribute { value: IArgBatteryVoltageMinThresholdPayload }
            export type IArgBatteryVoltageThreshold1Payload = ValueType;
            export interface IArgBatteryVoltageThreshold1 extends IAttribute { value: IArgBatteryVoltageThreshold1Payload }
            export type IArgBatteryVoltageThreshold2Payload = ValueType;
            export interface IArgBatteryVoltageThreshold2 extends IAttribute { value: IArgBatteryVoltageThreshold2Payload }
            export type IArgBatteryVoltageThreshold3Payload = ValueType;
            export interface IArgBatteryVoltageThreshold3 extends IAttribute { value: IArgBatteryVoltageThreshold3Payload }
            export type IArgBatteryCapacityPayload = ValueType;
            export interface IArgBatteryCapacity extends IAttribute { value: IArgBatteryCapacityPayload }
            export type IArgBinaryActiveTextPayload = ValueType;
            export interface IArgBinaryActiveText extends IAttribute { value: IArgBinaryActiveTextPayload }
            export type IArgBinaryInactiveTextPayload = ValueType;
            export interface IArgBinaryInactiveText extends IAttribute { value: IArgBinaryInactiveTextPayload }
            export type IArgBinaryMaxOffTimePayload = ValueType;
            export interface IArgBinaryMaxOffTime extends IAttribute { value: IArgBinaryMaxOffTimePayload }
            export type IArgBinaryMinOffTimePayload = ValueType;
            export interface IArgBinaryMinOffTime extends IAttribute { value: IArgBinaryMinOffTimePayload }
            export type IArgBinaryPolarityPayload = ValueType;
            export interface IArgBinaryPolarity extends IAttribute { value: IArgBinaryPolarityPayload }
            export type IArgBinaryPresentValuePayload = ValueType;
            export interface IArgBinaryPresentValue extends IAttribute { value: IArgBinaryPresentValuePayload }
            export type IArgBinaryPriorityPayload = { };
            export interface IArgBinaryPriority extends IArgument { value: IArgBinaryPriorityPayload }
            export type IArgBinaryPriorityArrayPayload = ValueType;
            export interface IArgBinaryPriorityArray extends IAttribute { value: IArgBinaryPriorityArrayPayload }
            export type IArgBinaryRelinquishDefaultPayload = ValueType;
            export interface IArgBinaryRelinquishDefault extends IAttribute { value: IArgBinaryRelinquishDefaultPayload }
            export type IArgCalculationPeriodPayload = ValueType;
            export interface IArgCalculationPeriod extends IAttribute { value: IArgCalculationPeriodPayload }
            export type IArgCheckInIntervalPayload = ValueType;
            export interface IArgCheckInInterval extends IAttribute { value: IArgCheckInIntervalPayload }
            export type IArgCheckInIntervalMinPayload = ValueType;
            export interface IArgCheckInIntervalMin extends IAttribute { value: IArgCheckInIntervalMinPayload }
            export type IArgChildMovedPayload = ValueType;
            export interface IArgChildMoved extends IAttribute { value: IArgChildMovedPayload }
            export type IArgClusterIdPayload = ValueType;
            export interface IArgClusterId extends IArgument { value: IArgClusterIdPayload }
            export type IArgClusterRevisionPayload = ValueType;
            export interface IArgClusterRevision extends IAttribute { value: IArgClusterRevisionPayload }
            export type IArgConfigurationPayload = ValueType;
            export interface IArgConfiguration extends IAttribute { value: IArgConfigurationPayload }
            export type IArgCurrentGroupPayload = ValueType;
            export interface IArgCurrentGroup extends IAttribute { value: IArgCurrentGroupPayload }
            export type IArgCurrentLevelPayload = ValueType;
            export interface IArgCurrentLevel extends IAttribute { value: IArgCurrentLevelPayload }
            export type IArgCurrentScenePayload = ValueType;
            export interface IArgCurrentScene extends IAttribute { value: IArgCurrentScenePayload }
            export type IArgCurrentTemperaturePayload = ValueType;
            export interface IArgCurrentTemperature extends IAttribute { value: IArgCurrentTemperaturePayload }
            export type IArgDateCodePayload = ValueType;
            export interface IArgDateCode extends IAttribute { value: IArgDateCodePayload }
            export type IArgDefaultMoveRatePayload = ValueType;
            export interface IArgDefaultMoveRate extends IAttribute { value: IArgDefaultMoveRatePayload }
            export type IArgDevicePayload = ValueType;
            export interface IArgDevice extends IArgument { value: IArgDevicePayload }
            export type IArgDeviceEnabledPayload = ValueType;
            export interface IArgDeviceEnabled extends IAttribute { value: IArgDeviceEnabledPayload }
            export type IArgDeviceTempAlarmMaskPayload = ValueType;
            export interface IArgDeviceTempAlarmMask extends IAttribute { value: IArgDeviceTempAlarmMaskPayload }
            export type IArgDisableLocalConfigPayload = ValueType;
            export interface IArgDisableLocalConfig extends IAttribute { value: IArgDisableLocalConfigPayload }
            export type IArgDistancePayload = ValueType;
            export interface IArgDistance extends IArgument { value: IArgDistancePayload }
            export type IArgDstEndPayload = ValueType;
            export interface IArgDstEnd extends IAttribute { value: IArgDstEndPayload }
            export type IArgDstShiftPayload = ValueType;
            export interface IArgDstShift extends IAttribute { value: IArgDstShiftPayload }
            export type IArgDstStartPayload = ValueType;
            export interface IArgDstStart extends IAttribute { value: IArgDstStartPayload }
            export type IArgEffectIdentifierPayload = ValueType;
            export interface IArgEffectIdentifier extends IArgument { value: IArgEffectIdentifierPayload }
            export type IArgEffectVariantPayload = ValueType;
            export interface IArgEffectVariant extends IArgument { value: IArgEffectVariantPayload }
            export type IArgFastPollTimeoutPayload = ValueType;
            export interface IArgFastPollTimeout extends IAttribute { value: IArgFastPollTimeoutPayload }
            export type IArgFastPollTimeoutMaxPayload = ValueType;
            export interface IArgFastPollTimeoutMax extends IAttribute { value: IArgFastPollTimeoutMaxPayload }
            export type IArgGenericDeviceClassPayload = ValueType;
            export interface IArgGenericDeviceClass extends IAttribute { value: IArgGenericDeviceClassPayload }
            export type IArgGenericDeviceTypePayload = ValueType;
            export interface IArgGenericDeviceType extends IAttribute { value: IArgGenericDeviceTypePayload }
            export type IArgGlobalSceneControlPayload = ValueType;
            export interface IArgGlobalSceneControl extends IAttribute { value: IArgGlobalSceneControlPayload }
            export type IArgGroupIdPayload = ValueType;
            export interface IArgGroupId extends IArgument { value: IArgGroupIdPayload }
            export type IArgGroupNameSupportPayload = ValueType;
            export interface IArgGroupNameSupport extends IAttribute { value: IArgGroupNameSupportPayload }
            export type IArgGroupCapacityPayload = ValueType;
            export interface IArgGroupCapacity extends IArgument { value: IArgGroupCapacityPayload }
            export type IArgGroupListPayload = ValueType;
            export interface IArgGroupList extends IArgument { value: IArgGroupListPayload }
            export type IArgGroupNamePayload = ValueType;
            export interface IArgGroupName extends IArgument { value: IArgGroupNamePayload }
            export type IArgHwVersionPayload = ValueType;
            export interface IArgHwVersion extends IAttribute { value: IArgHwVersionPayload }
            export type IArgHighTempDwellTripPointPayload = ValueType;
            export interface IArgHighTempDwellTripPoint extends IAttribute { value: IArgHighTempDwellTripPointPayload }
            export type IArgHighTempThresholdPayload = ValueType;
            export interface IArgHighTempThreshold extends IAttribute { value: IArgHighTempThresholdPayload }
            export type IArgIOApplicationTypePayload = ValueType;
            export interface IArgIOApplicationType extends IAttribute { value: IArgIOApplicationTypePayload }
            export type IArgIODescriptionPayload = ValueType;
            export interface IArgIODescription extends IAttribute { value: IArgIODescriptionPayload }
            export type IArgIOOutOfServicePayload = ValueType;
            export interface IArgIOOutOfService extends IAttribute { value: IArgIOOutOfServicePayload }
            export type IArgIOReliabilityPayload = ValueType;
            export interface IArgIOReliability extends IAttribute { value: IArgIOReliabilityPayload }
            export type IArgIOStatusFlagsPayload = ValueType;
            export interface IArgIOStatusFlags extends IAttribute { value: IArgIOStatusFlagsPayload }
            export type IArgIOUnitTypePayload = ValueType;
            export interface IArgIOUnitType extends IAttribute { value: IArgIOUnitTypePayload }
            export type IArgIdentifyEffectPayload = ValueType;
            export interface IArgIdentifyEffect extends IArgument { value: IArgIdentifyEffectPayload }
            export type IArgIdentifyEffectVariantPayload = ValueType;
            export interface IArgIdentifyEffectVariant extends IArgument { value: IArgIdentifyEffectVariantPayload }
            export type IArgIdentifyTimePayload = ValueType;
            export interface IArgIdentifyTime extends IAttribute { value: IArgIdentifyTimePayload }
            export type IArgIdentifyTimeoutPayload = ValueType;
            export interface IArgIdentifyTimeout extends IArgument { value: IArgIdentifyTimeoutPayload }
            export type IArgIkeaRemoteDirectionPayload = ValueType;
            export interface IArgIkeaRemoteDirection extends IArgument { value: IArgIkeaRemoteDirectionPayload }
            export type IArgJoinIndicationPayload = ValueType;
            export interface IArgJoinIndication extends IAttribute { value: IArgJoinIndicationPayload }
            export type IArgLedIndicationPayload = ValueType;
            export interface IArgLedIndication extends IAttribute { value: IArgLedIndicationPayload }
            export type IArgLastMessageLqiPayload = ValueType;
            export interface IArgLastMessageLqi extends IAttribute { value: IArgLastMessageLqiPayload }
            export type IArgLastMessageRssiPayload = ValueType;
            export interface IArgLastMessageRssi extends IAttribute { value: IArgLastMessageRssiPayload }
            export type IArgLastSetTimePayload = ValueType;
            export interface IArgLastSetTime extends IAttribute { value: IArgLastSetTimePayload }
            export type IArgLevelPayload = ValueType;
            export interface IArgLevel extends IArgument { value: IArgLevelPayload }
            export type IArgLevelControlOptionsPayload = ValueType;
            export interface IArgLevelControlOptions extends IAttribute { value: IArgLevelControlOptionsPayload }
            export type IArgLevelDirectionPayload = ValueType;
            export interface IArgLevelDirection extends IArgument { value: IArgLevelDirectionPayload }
            export type IArgLocalTimePayload = ValueType;
            export interface IArgLocalTime extends IAttribute { value: IArgLocalTimePayload }
            export type IArgLocationAgePayload = ValueType;
            export interface IArgLocationAge extends IAttribute { value: IArgLocationAgePayload }
            export type IArgLocationDescriptionPayload = ValueType;
            export interface IArgLocationDescription extends IAttribute { value: IArgLocationDescriptionPayload }
            export type IArgLocationMethodPayload = ValueType;
            export interface IArgLocationMethod extends IAttribute { value: IArgLocationMethodPayload }
            export type IArgLocationTypePayload = ValueType;
            export interface IArgLocationType extends IAttribute { value: IArgLocationTypePayload }
            export type IArgLocationFlagsPayload = ValueType;
            export interface IArgLocationFlags extends IArgument { value: IArgLocationFlagsPayload }
            export type IArgLongPollIntervalPayload = ValueType;
            export interface IArgLongPollInterval extends IAttribute { value: IArgLongPollIntervalPayload }
            export type IArgLongPollIntervalMinPayload = ValueType;
            export interface IArgLongPollIntervalMin extends IAttribute { value: IArgLongPollIntervalMinPayload }
            export type IArgLowTempDwellTripPointPayload = ValueType;
            export interface IArgLowTempDwellTripPoint extends IAttribute { value: IArgLowTempDwellTripPointPayload }
            export type IArgLowTempThresholdPayload = ValueType;
            export interface IArgLowTempThreshold extends IAttribute { value: IArgLowTempThresholdPayload }
            export type IArgMacRxBcastPayload = ValueType;
            export interface IArgMacRxBcast extends IAttribute { value: IArgMacRxBcastPayload }
            export type IArgMacRxUcastPayload = ValueType;
            export interface IArgMacRxUcast extends IAttribute { value: IArgMacRxUcastPayload }
            export type IArgMacTxBcastPayload = ValueType;
            export interface IArgMacTxBcast extends IAttribute { value: IArgMacTxBcastPayload }
            export type IArgMacTxUcastPayload = ValueType;
            export interface IArgMacTxUcast extends IAttribute { value: IArgMacTxUcastPayload }
            export type IArgMacTxUcastFailPayload = ValueType;
            export interface IArgMacTxUcastFail extends IAttribute { value: IArgMacTxUcastFailPayload }
            export type IArgMacTxUcastRetryPayload = ValueType;
            export interface IArgMacTxUcastRetry extends IAttribute { value: IArgMacTxUcastRetryPayload }
            export type IArgMainsAlarmMaskPayload = ValueType;
            export interface IArgMainsAlarmMask extends IAttribute { value: IArgMainsAlarmMaskPayload }
            export type IArgMainsFrequencyPayload = ValueType;
            export interface IArgMainsFrequency extends IAttribute { value: IArgMainsFrequencyPayload }
            export type IArgMainsVoltagePayload = ValueType;
            export interface IArgMainsVoltage extends IAttribute { value: IArgMainsVoltagePayload }
            export type IArgMainsVoltageDwellTripPointPayload = ValueType;
            export interface IArgMainsVoltageDwellTripPoint extends IAttribute { value: IArgMainsVoltageDwellTripPointPayload }
            export type IArgMainsVoltageMaxThresholdPayload = ValueType;
            export interface IArgMainsVoltageMaxThreshold extends IAttribute { value: IArgMainsVoltageMaxThresholdPayload }
            export type IArgMainsVoltageMinThresholdPayload = ValueType;
            export interface IArgMainsVoltageMinThreshold extends IAttribute { value: IArgMainsVoltageMinThresholdPayload }
            export type IArgManufacturerNamePayload = ValueType;
            export interface IArgManufacturerName extends IAttribute { value: IArgManufacturerNamePayload }
            export type IArgMaxTempExperiencedPayload = ValueType;
            export interface IArgMaxTempExperienced extends IAttribute { value: IArgMaxTempExperiencedPayload }
            export type IArgMinTempExperiencedPayload = ValueType;
            export interface IArgMinTempExperienced extends IAttribute { value: IArgMinTempExperiencedPayload }
            export type IArgModelIdentifierPayload = ValueType;
            export interface IArgModelIdentifier extends IAttribute { value: IArgModelIdentifierPayload }
            export type IArgMultistateNumberOfStatesPayload = ValueType;
            export interface IArgMultistateNumberOfStates extends IAttribute { value: IArgMultistateNumberOfStatesPayload }
            export type IArgMultistatePresentValuePayload = ValueType;
            export interface IArgMultistatePresentValue extends IAttribute { value: IArgMultistatePresentValuePayload }
            export type IArgMultistatePriorityPayload = { };
            export interface IArgMultistatePriority extends IArgument { value: IArgMultistatePriorityPayload }
            export type IArgMultistatePriorityArrayPayload = ValueType;
            export interface IArgMultistatePriorityArray extends IAttribute { value: IArgMultistatePriorityArrayPayload }
            export type IArgMultistateRelinquishDefaultPayload = ValueType;
            export interface IArgMultistateRelinquishDefault extends IAttribute { value: IArgMultistateRelinquishDefaultPayload }
            export type IArgMultistateTextPayload = ValueType;
            export interface IArgMultistateText extends IAttribute { value: IArgMultistateTextPayload }
            export type IArgNwkDecryptFailuresPayload = ValueType;
            export interface IArgNwkDecryptFailures extends IAttribute { value: IArgNwkDecryptFailuresPayload }
            export type IArgNwkFcFailurePayload = ValueType;
            export interface IArgNwkFcFailure extends IAttribute { value: IArgNwkFcFailurePayload }
            export type IArgNeighborAddedPayload = ValueType;
            export interface IArgNeighborAdded extends IAttribute { value: IArgNeighborAddedPayload }
            export type IArgNeighborInfoPayload = { Device?: IArgDevicePayload, XCoordinate?: IArgXCoordinatePayload, YCoordinate?: IArgYCoordinatePayload, ZCoordinate?: IArgZCoordinatePayload, Rssi?: IArgRssiPayload, NumberRssiMeasurements?: IArgNumberRssiMeasurementsPayload, };
            export interface IArgNeighborInfo extends IArgument { value: IArgNeighborInfoPayload }
            export type IArgNeighborRemovedPayload = ValueType;
            export interface IArgNeighborRemoved extends IAttribute { value: IArgNeighborRemovedPayload }
            export type IArgNeighborStalePayload = ValueType;
            export interface IArgNeighborStale extends IAttribute { value: IArgNeighborStalePayload }
            export type IArgNeighborsInfoListPayload = ValueType;
            export interface IArgNeighborsInfoList extends IArgument { value: IArgNeighborsInfoListPayload }
            export type IArgNumberRssiMeasurementsPayload = ValueType;
            export interface IArgNumberRssiMeasurements extends IAttribute { value: IArgNumberRssiMeasurementsPayload }
            export type IArgNumberResponsesPayload = ValueType;
            export interface IArgNumberResponses extends IArgument { value: IArgNumberResponsesPayload }
            export type IArgNumberOfDevicesPayload = ValueType;
            export interface IArgNumberOfDevices extends IAttribute { value: IArgNumberOfDevicesPayload }
            export type IArgNumberOfResetsPayload = ValueType;
            export interface IArgNumberOfResets extends IAttribute { value: IArgNumberOfResetsPayload }
            export type IArgOffTransitionTimePayload = ValueType;
            export interface IArgOffTransitionTime extends IAttribute { value: IArgOffTransitionTimePayload }
            export type IArgOffWaitTimePayload = ValueType;
            export interface IArgOffWaitTime extends IAttribute { value: IArgOffWaitTimePayload }
            export type IArgOnLevelPayload = ValueType;
            export interface IArgOnLevel extends IAttribute { value: IArgOnLevelPayload }
            export type IArgOnOffPayload = ValueType;
            export interface IArgOnOff extends IAttribute { value: IArgOnOffPayload }
            export type IArgOnTimePayload = ValueType;
            export interface IArgOnTime extends IAttribute { value: IArgOnTimePayload }
            export type IArgOnTransitionTimePayload = ValueType;
            export interface IArgOnTransitionTime extends IAttribute { value: IArgOnTransitionTimePayload }
            export type IArgOnOffTransistionTimePayload = ValueType;
            export interface IArgOnOffTransistionTime extends IAttribute { value: IArgOnOffTransistionTimePayload }
            export type IArgOnOffControlPayload = ValueType;
            export interface IArgOnOffControl extends IArgument { value: IArgOnOffControlPayload }
            export type IArgOverTempTotalDwellPayload = ValueType;
            export interface IArgOverTempTotalDwell extends IAttribute { value: IArgOverTempTotalDwellPayload }
            export type IArgPacketBufferAllocFailuresPayload = ValueType;
            export interface IArgPacketBufferAllocFailures extends IAttribute { value: IArgPacketBufferAllocFailuresPayload }
            export type IArgPacketValidateDropcountPayload = ValueType;
            export interface IArgPacketValidateDropcount extends IAttribute { value: IArgPacketValidateDropcountPayload }
            export type IArgPathLossExponentPayload = ValueType;
            export interface IArgPathLossExponent extends IAttribute { value: IArgPathLossExponentPayload }
            export type IArgPersistensMemoryWritesPayload = ValueType;
            export interface IArgPersistensMemoryWrites extends IAttribute { value: IArgPersistensMemoryWritesPayload }
            export type IArgPhyToMacQueueLimitReachedPayload = ValueType;
            export interface IArgPhyToMacQueueLimitReached extends IAttribute { value: IArgPhyToMacQueueLimitReachedPayload }
            export type IArgPhysicalEnvironmentPayload = ValueType;
            export interface IArgPhysicalEnvironment extends IAttribute { value: IArgPhysicalEnvironmentPayload }
            export type IArgPowerPayload = ValueType;
            export interface IArgPower extends IAttribute { value: IArgPowerPayload }
            export type IArgPowerOnLevelPayload = ValueType;
            export interface IArgPowerOnLevel extends IAttribute { value: IArgPowerOnLevelPayload }
            export type IArgPowerSourcePayload = ValueType;
            export interface IArgPowerSource extends IAttribute { value: IArgPowerSourcePayload }
            export type IArgPoweronOnOffPayload = ValueType;
            export interface IArgPoweronOnOff extends IAttribute { value: IArgPoweronOnOffPayload }
            export type IArgProductUrlPayload = ValueType;
            export interface IArgProductUrl extends IAttribute { value: IArgProductUrlPayload }
            export type IArgProductCodePayload = ValueType;
            export interface IArgProductCode extends IAttribute { value: IArgProductCodePayload }
            export type IArgQualityMeasurePayload = ValueType;
            export interface IArgQualityMeasure extends IAttribute { value: IArgQualityMeasurePayload }
            export type IArgQualityIndexPayload = ValueType;
            export interface IArgQualityIndex extends IArgument { value: IArgQualityIndexPayload }
            export type IArgRssiPayload = ValueType;
            export interface IArgRssi extends IArgument { value: IArgRssiPayload }
            export type IArgRatePayload = ValueType;
            export interface IArgRate extends IArgument { value: IArgRatePayload }
            export type IArgRelayedUcastPayload = ValueType;
            export interface IArgRelayedUcast extends IAttribute { value: IArgRelayedUcastPayload }
            export type IArgRemainingTimePayload = ValueType;
            export interface IArgRemainingTime extends IAttribute { value: IArgRemainingTimePayload }
            export type IArgReportingPeriodPayload = ValueType;
            export interface IArgReportingPeriod extends IAttribute { value: IArgReportingPeriodPayload }
            export type IArgResolutionPayload = ValueType;
            export interface IArgResolution extends IArgument { value: IArgResolutionPayload }
            export type IArgRouteDiscInitiatedPayload = ValueType;
            export interface IArgRouteDiscInitiated extends IAttribute { value: IArgRouteDiscInitiatedPayload }
            export type IArgSwBuildIdPayload = ValueType;
            export interface IArgSwBuildId extends IAttribute { value: IArgSwBuildIdPayload }
            export type IArgSceneCapacityPayload = ValueType;
            export interface IArgSceneCapacity extends IArgument { value: IArgSceneCapacityPayload }
            export type IArgSceneCopyModePayload = ValueType;
            export interface IArgSceneCopyMode extends IArgument { value: IArgSceneCopyModePayload }
            export type IArgSceneCountPayload = ValueType;
            export interface IArgSceneCount extends IAttribute { value: IArgSceneCountPayload }
            export type IArgSceneExtensionsPayload = ValueType;
            export interface IArgSceneExtensions extends IArgument { value: IArgSceneExtensionsPayload }
            export type IArgSceneIdPayload = ValueType;
            export interface IArgSceneId extends IArgument { value: IArgSceneIdPayload }
            export type IArgSceneLastConfiguredByPayload = ValueType;
            export interface IArgSceneLastConfiguredBy extends IAttribute { value: IArgSceneLastConfiguredByPayload }
            export type IArgSceneNamePayload = ValueType;
            export interface IArgSceneName extends IArgument { value: IArgSceneNamePayload }
            export type IArgSceneNameSupportPayload = ValueType;
            export interface IArgSceneNameSupport extends IAttribute { value: IArgSceneNameSupportPayload }
            export type IArgSceneValidPayload = ValueType;
            export interface IArgSceneValid extends IAttribute { value: IArgSceneValidPayload }
            export type IArgSceneListPayload = ValueType;
            export interface IArgSceneList extends IArgument { value: IArgSceneListPayload }
            export type IArgSensitivityPayload = ValueType;
            export interface IArgSensitivity extends IAttribute { value: IArgSensitivityPayload }
            export type IArgShortPollIntervalPayload = ValueType;
            export interface IArgShortPollInterval extends IAttribute { value: IArgShortPollIntervalPayload }
            export type IArgStackVersionPayload = ValueType;
            export interface IArgStackVersion extends IAttribute { value: IArgStackVersionPayload }
            export type IArgStandardTimePayload = ValueType;
            export interface IArgStandardTime extends IAttribute { value: IArgStandardTimePayload }
            export type IArgStatusPayload = ValueType;
            export interface IArgStatus extends IArgument { value: IArgStatusPayload }
            export type IArgStepSizePayload = ValueType;
            export interface IArgStepSize extends IArgument { value: IArgStepSizePayload }
            export type IArgSwitchActionsPayload = ValueType;
            export interface IArgSwitchActions extends IAttribute { value: IArgSwitchActionsPayload }
            export type IArgSwitchTypePayload = ValueType;
            export interface IArgSwitchType extends IAttribute { value: IArgSwitchTypePayload }
            export type IArgTimePayload = ValueType;
            export interface IArgTime extends IAttribute { value: IArgTimePayload }
            export type IArgTimeStatusPayload = ValueType;
            export interface IArgTimeStatus extends IAttribute { value: IArgTimeStatusPayload }
            export type IArgTimeZonePayload = ValueType;
            export interface IArgTimeZone extends IAttribute { value: IArgTimeZonePayload }
            export type IArgTransitionTimePayload = ValueType;
            export interface IArgTransitionTime extends IArgument { value: IArgTransitionTimePayload }
            export type IArgTransitionTimeSecPayload = ValueType;
            export interface IArgTransitionTimeSec extends IArgument { value: IArgTransitionTimeSecPayload }
            export type IArgUserTestPayload = ValueType;
            export interface IArgUserTest extends IAttribute { value: IArgUserTestPayload }
            export type IArgValidUntilTimePayload = ValueType;
            export interface IArgValidUntilTime extends IAttribute { value: IArgValidUntilTimePayload }
            export type IArgXCoordinatePayload = ValueType;
            export interface IArgXCoordinate extends IAttribute { value: IArgXCoordinatePayload }
            export type IArgYCoordinatePayload = ValueType;
            export interface IArgYCoordinate extends IAttribute { value: IArgYCoordinatePayload }
            export type IArgZCoordinatePayload = ValueType;
            export interface IArgZCoordinate extends IAttribute { value: IArgZCoordinatePayload }
            export type IArgZclVersionPayload = ValueType;
            export interface IArgZclVersion extends IAttribute { value: IArgZclVersionPayload }    }

    export namespace IHvac {
        // noinspection ES6UnusedImports
        import IArgument = ZigBee.IArgument;
        // noinspection ES6UnusedImports
        import IAttribute = ZigBee.IAttribute;
        // noinspection ES6UnusedImports
        import ValueType = ZigBee.ValueType;
        
        export namespace PumpConfigurationAndControl {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace Thermostat {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
            export type ICmdSetpointRaiseLowerPayload = { SetpointMode?: IArgSetpointModePayload, SetpointAmount?: IArgSetpointAmountPayload, }
            export interface ICmdSetpointRaiseLower extends ICommand { value: ICmdSetpointRaiseLowerPayload }
            export type ICmdSetWeeklySchedulePayload = { SetWeeklyNumberOfTransitions?: IArgSetWeeklyNumberOfTransitionsPayload, SetWeeklyDayOfWeek?: IArgSetWeeklyDayOfWeekPayload, SetWeeklyMode?: IArgSetWeeklyModePayload, SetWeeklyTransitionTime1?: IArgSetWeeklyTransitionTime1Payload, SetWeeklyHeatSetpoint1?: IArgSetWeeklyHeatSetpoint1Payload, SetWeeklyCoolSetpoint1?: IArgSetWeeklyCoolSetpoint1Payload, SetWeeklyTransitionTime10?: IArgSetWeeklyTransitionTime10Payload, SetWeeklyHeatSetpoint10?: IArgSetWeeklyHeatSetpoint10Payload, SetWeeklyCoolSetpoint10?: IArgSetWeeklyCoolSetpoint10Payload, }
            export interface ICmdSetWeeklySchedule extends ICommand { value: ICmdSetWeeklySchedulePayload }
            export type ICmdGetWeeklySchedulePayload = { GetWeeklyDaysToReturn?: IArgGetWeeklyDaysToReturnPayload, GetWeeklyModeToReturn?: IArgGetWeeklyModeToReturnPayload, }
            export interface ICmdGetWeeklySchedule extends ICommand { value: ICmdGetWeeklySchedulePayload }
            export type ICmdClearWeeklySchedulePayload = { }
            export interface ICmdClearWeeklySchedule extends ICommand { value: ICmdClearWeeklySchedulePayload }
            export type ICmdGetRelayStatusLogPayload = { }
            export interface ICmdGetRelayStatusLog extends ICommand { value: ICmdGetRelayStatusLogPayload }
            export type ICmdGetWeeklyScheduleResponsePayload = { SetWeeklyNumberOfTransitions?: IArgSetWeeklyNumberOfTransitionsPayload, SetWeeklyDayOfWeek?: IArgSetWeeklyDayOfWeekPayload, SetWeeklyMode?: IArgSetWeeklyModePayload, SetWeeklyTransitionTime1?: IArgSetWeeklyTransitionTime1Payload, SetWeeklyHeatSetpoint1?: IArgSetWeeklyHeatSetpoint1Payload, SetWeeklyCoolSetpoint1?: IArgSetWeeklyCoolSetpoint1Payload, SetWeeklyTransitionTime10?: IArgSetWeeklyTransitionTime10Payload, SetWeeklyHeatSetpoint10?: IArgSetWeeklyHeatSetpoint10Payload, SetWeeklyCoolSetpoint10?: IArgSetWeeklyCoolSetpoint10Payload, }
            export interface ICmdGetWeeklyScheduleResponse extends ICommand { value: ICmdGetWeeklyScheduleResponsePayload }
            export type ICmdGetRelayStatusLogResponsePayload = { RelayStatusLogTimeOfDay?: IArgRelayStatusLogTimeOfDayPayload, RelayStatus?: IArgRelayStatusPayload, RelayStatusLocalTemperature?: IArgRelayStatusLocalTemperaturePayload, RelayStatusHumidity?: IArgRelayStatusHumidityPayload, RelayStatusSetpoint?: IArgRelayStatusSetpointPayload, RelayStatusUnreadEntries?: IArgRelayStatusUnreadEntriesPayload, }
            export interface ICmdGetRelayStatusLogResponse extends ICommand { value: ICmdGetRelayStatusLogResponsePayload }
        }

        export namespace FanControl {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace DehumidificationControl {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace ThermostatUiConfiguration {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

            export type IArgAcCapacityPayload = ValueType;
            export interface IArgAcCapacity extends IAttribute { value: IArgAcCapacityPayload }
            export type IArgAcCapacityFormatPayload = ValueType;
            export interface IArgAcCapacityFormat extends IAttribute { value: IArgAcCapacityFormatPayload }
            export type IArgAcCoilTemperaturePayload = ValueType;
            export interface IArgAcCoilTemperature extends IAttribute { value: IArgAcCoilTemperaturePayload }
            export type IArgAcCompressorTypePayload = ValueType;
            export interface IArgAcCompressorType extends IAttribute { value: IArgAcCompressorTypePayload }
            export type IArgAcErrorCodePayload = ValueType;
            export interface IArgAcErrorCode extends IAttribute { value: IArgAcErrorCodePayload }
            export type IArgAcLouverPositionPayload = ValueType;
            export interface IArgAcLouverPosition extends IAttribute { value: IArgAcLouverPositionPayload }
            export type IArgAcRefrigerantTypePayload = ValueType;
            export interface IArgAcRefrigerantType extends IAttribute { value: IArgAcRefrigerantTypePayload }
            export type IArgAcTypePayload = ValueType;
            export interface IArgAcType extends IAttribute { value: IArgAcTypePayload }
            export type IArgAbsMaxCoolSetpointLimitPayload = ValueType;
            export interface IArgAbsMaxCoolSetpointLimit extends IAttribute { value: IArgAbsMaxCoolSetpointLimitPayload }
            export type IArgAbsMaxHeatSetpointLimitPayload = ValueType;
            export interface IArgAbsMaxHeatSetpointLimit extends IAttribute { value: IArgAbsMaxHeatSetpointLimitPayload }
            export type IArgAbsMinCoolSetpointLimitPayload = ValueType;
            export interface IArgAbsMinCoolSetpointLimit extends IAttribute { value: IArgAbsMinCoolSetpointLimitPayload }
            export type IArgAbsMinHeatSetpointLimitPayload = ValueType;
            export interface IArgAbsMinHeatSetpointLimit extends IAttribute { value: IArgAbsMinHeatSetpointLimitPayload }
            export type IArgCapacityPayload = ValueType;
            export interface IArgCapacity extends IAttribute { value: IArgCapacityPayload }
            export type IArgControlModePayload = ValueType;
            export interface IArgControlMode extends IAttribute { value: IArgControlModePayload }
            export type IArgControlSequenceOfOperationPayload = ValueType;
            export interface IArgControlSequenceOfOperation extends IAttribute { value: IArgControlSequenceOfOperationPayload }
            export type IArgDehumidificationCoolingPayload = ValueType;
            export interface IArgDehumidificationCooling extends IAttribute { value: IArgDehumidificationCoolingPayload }
            export type IArgDehumidificationHysteresisPayload = ValueType;
            export interface IArgDehumidificationHysteresis extends IAttribute { value: IArgDehumidificationHysteresisPayload }
            export type IArgDehumidificationLockoutPayload = ValueType;
            export interface IArgDehumidificationLockout extends IAttribute { value: IArgDehumidificationLockoutPayload }
            export type IArgDehumidificationMaxCoolPayload = ValueType;
            export interface IArgDehumidificationMaxCool extends IAttribute { value: IArgDehumidificationMaxCoolPayload }
            export type IArgEffectiveControlModePayload = ValueType;
            export interface IArgEffectiveControlMode extends IAttribute { value: IArgEffectiveControlModePayload }
            export type IArgEffectiveOperationModePayload = ValueType;
            export interface IArgEffectiveOperationMode extends IAttribute { value: IArgEffectiveOperationModePayload }
            export type IArgEmergencyHeatDeltaPayload = ValueType;
            export interface IArgEmergencyHeatDelta extends IAttribute { value: IArgEmergencyHeatDeltaPayload }
            export type IArgFanModePayload = ValueType;
            export interface IArgFanMode extends IAttribute { value: IArgFanModePayload }
            export type IArgFanModeSequencePayload = ValueType;
            export interface IArgFanModeSequence extends IAttribute { value: IArgFanModeSequencePayload }
            export type IArgGetWeeklyDaysToReturnPayload = ValueType;
            export interface IArgGetWeeklyDaysToReturn extends IArgument { value: IArgGetWeeklyDaysToReturnPayload }
            export type IArgGetWeeklyModeToReturnPayload = ValueType;
            export interface IArgGetWeeklyModeToReturn extends IArgument { value: IArgGetWeeklyModeToReturnPayload }
            export type IArgHvacSystemTypeConfigurationPayload = ValueType;
            export interface IArgHvacSystemTypeConfiguration extends IAttribute { value: IArgHvacSystemTypeConfigurationPayload }
            export type IArgKeypadLockoutPayload = ValueType;
            export interface IArgKeypadLockout extends IAttribute { value: IArgKeypadLockoutPayload }
            export type IArgLifetimeEnergyConsumedPayload = ValueType;
            export interface IArgLifetimeEnergyConsumed extends IAttribute { value: IArgLifetimeEnergyConsumedPayload }
            export type IArgLifetimeRunningHoursPayload = ValueType;
            export interface IArgLifetimeRunningHours extends IAttribute { value: IArgLifetimeRunningHoursPayload }
            export type IArgLocalTemperaturePayload = ValueType;
            export interface IArgLocalTemperature extends IAttribute { value: IArgLocalTemperaturePayload }
            export type IArgLocalTemperatureCalibrationPayload = ValueType;
            export interface IArgLocalTemperatureCalibration extends IAttribute { value: IArgLocalTemperatureCalibrationPayload }
            export type IArgMaxCompPressurePayload = ValueType;
            export interface IArgMaxCompPressure extends IAttribute { value: IArgMaxCompPressurePayload }
            export type IArgMaxConstFlowPayload = ValueType;
            export interface IArgMaxConstFlow extends IAttribute { value: IArgMaxConstFlowPayload }
            export type IArgMaxConstPressurePayload = ValueType;
            export interface IArgMaxConstPressure extends IAttribute { value: IArgMaxConstPressurePayload }
            export type IArgMaxConstSpeedPayload = ValueType;
            export interface IArgMaxConstSpeed extends IAttribute { value: IArgMaxConstSpeedPayload }
            export type IArgMaxConstTempPayload = ValueType;
            export interface IArgMaxConstTemp extends IAttribute { value: IArgMaxConstTempPayload }
            export type IArgMaxCoolSetpointLimitPayload = ValueType;
            export interface IArgMaxCoolSetpointLimit extends IAttribute { value: IArgMaxCoolSetpointLimitPayload }
            export type IArgMaxFlowPayload = ValueType;
            export interface IArgMaxFlow extends IAttribute { value: IArgMaxFlowPayload }
            export type IArgMaxHeatSetpointLimitPayload = ValueType;
            export interface IArgMaxHeatSetpointLimit extends IAttribute { value: IArgMaxHeatSetpointLimitPayload }
            export type IArgMaxPressurePayload = ValueType;
            export interface IArgMaxPressure extends IAttribute { value: IArgMaxPressurePayload }
            export type IArgMaxSpeedPayload = ValueType;
            export interface IArgMaxSpeed extends IAttribute { value: IArgMaxSpeedPayload }
            export type IArgMinCompPressurePayload = ValueType;
            export interface IArgMinCompPressure extends IAttribute { value: IArgMinCompPressurePayload }
            export type IArgMinConstFlowPayload = ValueType;
            export interface IArgMinConstFlow extends IAttribute { value: IArgMinConstFlowPayload }
            export type IArgMinConstPressurePayload = ValueType;
            export interface IArgMinConstPressure extends IAttribute { value: IArgMinConstPressurePayload }
            export type IArgMinConstSpeedPayload = ValueType;
            export interface IArgMinConstSpeed extends IAttribute { value: IArgMinConstSpeedPayload }
            export type IArgMinConstTempPayload = ValueType;
            export interface IArgMinConstTemp extends IAttribute { value: IArgMinConstTempPayload }
            export type IArgMinCoolSetpointLimitPayload = ValueType;
            export interface IArgMinCoolSetpointLimit extends IAttribute { value: IArgMinCoolSetpointLimitPayload }
            export type IArgMinHeatSetpointLimitPayload = ValueType;
            export interface IArgMinHeatSetpointLimit extends IAttribute { value: IArgMinHeatSetpointLimitPayload }
            export type IArgMinSetpointDeadBandPayload = ValueType;
            export interface IArgMinSetpointDeadBand extends IAttribute { value: IArgMinSetpointDeadBandPayload }
            export type IArgNumberOfDailyTransitionsPayload = ValueType;
            export interface IArgNumberOfDailyTransitions extends IAttribute { value: IArgNumberOfDailyTransitionsPayload }
            export type IArgNumberOfWeeklyTransitionsPayload = ValueType;
            export interface IArgNumberOfWeeklyTransitions extends IAttribute { value: IArgNumberOfWeeklyTransitionsPayload }
            export type IArgOccupancyPayload = ValueType;
            export interface IArgOccupancy extends IAttribute { value: IArgOccupancyPayload }
            export type IArgOccupiedCoolingSetpointPayload = ValueType;
            export interface IArgOccupiedCoolingSetpoint extends IAttribute { value: IArgOccupiedCoolingSetpointPayload }
            export type IArgOccupiedHeatingSetpointPayload = ValueType;
            export interface IArgOccupiedHeatingSetpoint extends IAttribute { value: IArgOccupiedHeatingSetpointPayload }
            export type IArgOccupiedSetbackPayload = ValueType;
            export interface IArgOccupiedSetback extends IAttribute { value: IArgOccupiedSetbackPayload }
            export type IArgOccupiedSetbackMaxPayload = ValueType;
            export interface IArgOccupiedSetbackMax extends IAttribute { value: IArgOccupiedSetbackMaxPayload }
            export type IArgOccupiedSetbackMinPayload = ValueType;
            export interface IArgOccupiedSetbackMin extends IAttribute { value: IArgOccupiedSetbackMinPayload }
            export type IArgOperationModePayload = ValueType;
            export interface IArgOperationMode extends IAttribute { value: IArgOperationModePayload }
            export type IArgOutdoorTemperaturePayload = ValueType;
            export interface IArgOutdoorTemperature extends IAttribute { value: IArgOutdoorTemperaturePayload }
            export type IArgPiCoolingDemandPayload = ValueType;
            export interface IArgPiCoolingDemand extends IAttribute { value: IArgPiCoolingDemandPayload }
            export type IArgPiHeatingDemandPayload = ValueType;
            export interface IArgPiHeatingDemand extends IAttribute { value: IArgPiHeatingDemandPayload }
            export type IArgPowerPayload = ValueType;
            export interface IArgPower extends IAttribute { value: IArgPowerPayload }
            export type IArgPumpAlarmMaskPayload = ValueType;
            export interface IArgPumpAlarmMask extends IAttribute { value: IArgPumpAlarmMaskPayload }
            export type IArgPumpStatusPayload = ValueType;
            export interface IArgPumpStatus extends IAttribute { value: IArgPumpStatusPayload }
            export type IArgRhDehumidificationSetpointPayload = ValueType;
            export interface IArgRhDehumidificationSetpoint extends IAttribute { value: IArgRhDehumidificationSetpointPayload }
            export type IArgRelativeHumidityPayload = ValueType;
            export interface IArgRelativeHumidity extends IAttribute { value: IArgRelativeHumidityPayload }
            export type IArgRelativeHumidityDisplayPayload = ValueType;
            export interface IArgRelativeHumidityDisplay extends IAttribute { value: IArgRelativeHumidityDisplayPayload }
            export type IArgRelativeHumidityModePayload = ValueType;
            export interface IArgRelativeHumidityMode extends IAttribute { value: IArgRelativeHumidityModePayload }
            export type IArgRelayStatusPayload = ValueType;
            export interface IArgRelayStatus extends IArgument { value: IArgRelayStatusPayload }
            export type IArgRelayStatusHumidityPayload = ValueType;
            export interface IArgRelayStatusHumidity extends IArgument { value: IArgRelayStatusHumidityPayload }
            export type IArgRelayStatusLocalTemperaturePayload = ValueType;
            export interface IArgRelayStatusLocalTemperature extends IArgument { value: IArgRelayStatusLocalTemperaturePayload }
            export type IArgRelayStatusLogTimeOfDayPayload = ValueType;
            export interface IArgRelayStatusLogTimeOfDay extends IArgument { value: IArgRelayStatusLogTimeOfDayPayload }
            export type IArgRelayStatusSetpointPayload = ValueType;
            export interface IArgRelayStatusSetpoint extends IArgument { value: IArgRelayStatusSetpointPayload }
            export type IArgRelayStatusUnreadEntriesPayload = ValueType;
            export interface IArgRelayStatusUnreadEntries extends IArgument { value: IArgRelayStatusUnreadEntriesPayload }
            export type IArgRemoteSensingPayload = ValueType;
            export interface IArgRemoteSensing extends IAttribute { value: IArgRemoteSensingPayload }
            export type IArgScheduleProgrammingVisibilityPayload = ValueType;
            export interface IArgScheduleProgrammingVisibility extends IAttribute { value: IArgScheduleProgrammingVisibilityPayload }
            export type IArgSetWeeklyCoolSetpoint1Payload = ValueType;
            export interface IArgSetWeeklyCoolSetpoint1 extends IArgument { value: IArgSetWeeklyCoolSetpoint1Payload }
            export type IArgSetWeeklyCoolSetpoint10Payload = ValueType;
            export interface IArgSetWeeklyCoolSetpoint10 extends IArgument { value: IArgSetWeeklyCoolSetpoint10Payload }
            export type IArgSetWeeklyDayOfWeekPayload = ValueType;
            export interface IArgSetWeeklyDayOfWeek extends IArgument { value: IArgSetWeeklyDayOfWeekPayload }
            export type IArgSetWeeklyHeatSetpoint1Payload = ValueType;
            export interface IArgSetWeeklyHeatSetpoint1 extends IArgument { value: IArgSetWeeklyHeatSetpoint1Payload }
            export type IArgSetWeeklyHeatSetpoint10Payload = ValueType;
            export interface IArgSetWeeklyHeatSetpoint10 extends IArgument { value: IArgSetWeeklyHeatSetpoint10Payload }
            export type IArgSetWeeklyModePayload = ValueType;
            export interface IArgSetWeeklyMode extends IArgument { value: IArgSetWeeklyModePayload }
            export type IArgSetWeeklyNumberOfTransitionsPayload = ValueType;
            export interface IArgSetWeeklyNumberOfTransitions extends IArgument { value: IArgSetWeeklyNumberOfTransitionsPayload }
            export type IArgSetWeeklyTransitionTime1Payload = ValueType;
            export interface IArgSetWeeklyTransitionTime1 extends IArgument { value: IArgSetWeeklyTransitionTime1Payload }
            export type IArgSetWeeklyTransitionTime10Payload = ValueType;
            export interface IArgSetWeeklyTransitionTime10 extends IArgument { value: IArgSetWeeklyTransitionTime10Payload }
            export type IArgSetpointAmountPayload = ValueType;
            export interface IArgSetpointAmount extends IArgument { value: IArgSetpointAmountPayload }
            export type IArgSetpointChangeAmountPayload = ValueType;
            export interface IArgSetpointChangeAmount extends IAttribute { value: IArgSetpointChangeAmountPayload }
            export type IArgSetpointChangeSourcePayload = ValueType;
            export interface IArgSetpointChangeSource extends IAttribute { value: IArgSetpointChangeSourcePayload }
            export type IArgSetpointChangeSourceTimestampPayload = ValueType;
            export interface IArgSetpointChangeSourceTimestamp extends IAttribute { value: IArgSetpointChangeSourceTimestampPayload }
            export type IArgSetpointModePayload = ValueType;
            export interface IArgSetpointMode extends IArgument { value: IArgSetpointModePayload }
            export type IArgSpeedPayload = ValueType;
            export interface IArgSpeed extends IAttribute { value: IArgSpeedPayload }
            export type IArgStartOfWeekPayload = ValueType;
            export interface IArgStartOfWeek extends IAttribute { value: IArgStartOfWeekPayload }
            export type IArgSystemModePayload = ValueType;
            export interface IArgSystemMode extends IAttribute { value: IArgSystemModePayload }
            export type IArgTemperatureDisplayModePayload = ValueType;
            export interface IArgTemperatureDisplayMode extends IAttribute { value: IArgTemperatureDisplayModePayload }
            export type IArgTemperatureSetpointHoldPayload = ValueType;
            export interface IArgTemperatureSetpointHold extends IAttribute { value: IArgTemperatureSetpointHoldPayload }
            export type IArgTemperatureSetpointHoldDurationPayload = ValueType;
            export interface IArgTemperatureSetpointHoldDuration extends IAttribute { value: IArgTemperatureSetpointHoldDurationPayload }
            export type IArgThermostatAlarmMaskPayload = ValueType;
            export interface IArgThermostatAlarmMask extends IAttribute { value: IArgThermostatAlarmMaskPayload }
            export type IArgThermostatProgrammingOperationModePayload = ValueType;
            export interface IArgThermostatProgrammingOperationMode extends IAttribute { value: IArgThermostatProgrammingOperationModePayload }
            export type IArgThermostatRunningModePayload = ValueType;
            export interface IArgThermostatRunningMode extends IAttribute { value: IArgThermostatRunningModePayload }
            export type IArgThermostatRunningStatePayload = ValueType;
            export interface IArgThermostatRunningState extends IAttribute { value: IArgThermostatRunningStatePayload }
            export type IArgUnoccupiedCoolingSetpointPayload = ValueType;
            export interface IArgUnoccupiedCoolingSetpoint extends IAttribute { value: IArgUnoccupiedCoolingSetpointPayload }
            export type IArgUnoccupiedHeatingSetpointPayload = ValueType;
            export interface IArgUnoccupiedHeatingSetpoint extends IAttribute { value: IArgUnoccupiedHeatingSetpointPayload }
            export type IArgUnoccupiedSetbackPayload = ValueType;
            export interface IArgUnoccupiedSetback extends IAttribute { value: IArgUnoccupiedSetbackPayload }
            export type IArgUnoccupiedSetbackMaxPayload = ValueType;
            export interface IArgUnoccupiedSetbackMax extends IAttribute { value: IArgUnoccupiedSetbackMaxPayload }
            export type IArgUnoccupiedSetbackMinPayload = ValueType;
            export interface IArgUnoccupiedSetbackMin extends IAttribute { value: IArgUnoccupiedSetbackMinPayload }    }

    export namespace ISecurityAndSafety {
        // noinspection ES6UnusedImports
        import IArgument = ZigBee.IArgument;
        // noinspection ES6UnusedImports
        import IAttribute = ZigBee.IAttribute;
        // noinspection ES6UnusedImports
        import ValueType = ZigBee.ValueType;
        
        export namespace IasZone {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
            export type ICmdZoneEnrollResponsePayload = { EnrollResponseCode?: IArgEnrollResponseCodePayload, ZoneId?: IArgZoneIdPayload, }
            export interface ICmdZoneEnrollResponse extends ICommand { value: ICmdZoneEnrollResponsePayload }
            export type ICmdInitiateNormalOperationModePayload = { }
            export interface ICmdInitiateNormalOperationMode extends ICommand { value: ICmdInitiateNormalOperationModePayload }
            export type ICmdInitiateTestModePayload = { TestModeDuration?: IArgTestModeDurationPayload, CurrentZoneSensitivityLevel?: IArgCurrentZoneSensitivityLevelPayload, }
            export interface ICmdInitiateTestMode extends ICommand { value: ICmdInitiateTestModePayload }
            export type ICmdZoneStateChangeNotificationPayload = { ZoneStatus?: IArgZoneStatusPayload, ExtendedStatus?: IArgExtendedStatusPayload, ZoneId?: IArgZoneIdPayload, Delay?: IArgDelayPayload, }
            export interface ICmdZoneStateChangeNotification extends ICommand { value: ICmdZoneStateChangeNotificationPayload }
            export type ICmdZoneEnrollRequestPayload = { ZoneType?: IArgZoneTypePayload, ManufacturerCode?: IArgManufacturerCodePayload, }
            export interface ICmdZoneEnrollRequest extends ICommand { value: ICmdZoneEnrollRequestPayload }
        }

            export type IArgCurrentZoneSensitivityLevelPayload = ValueType;
            export interface IArgCurrentZoneSensitivityLevel extends IAttribute { value: IArgCurrentZoneSensitivityLevelPayload }
            export type IArgDelayPayload = ValueType;
            export interface IArgDelay extends IArgument { value: IArgDelayPayload }
            export type IArgEnrollResponseCodePayload = ValueType;
            export interface IArgEnrollResponseCode extends IArgument { value: IArgEnrollResponseCodePayload }
            export type IArgExtendedStatusPayload = ValueType;
            export interface IArgExtendedStatus extends IArgument { value: IArgExtendedStatusPayload }
            export type IArgIasCieAddressMacPayload = ValueType;
            export interface IArgIasCieAddressMac extends IAttribute { value: IArgIasCieAddressMacPayload }
            export type IArgManufacturerCodePayload = ValueType;
            export interface IArgManufacturerCode extends IArgument { value: IArgManufacturerCodePayload }
            export type IArgNumberOfZoneSensitivityLevelsSupportedPayload = ValueType;
            export interface IArgNumberOfZoneSensitivityLevelsSupported extends IAttribute { value: IArgNumberOfZoneSensitivityLevelsSupportedPayload }
            export type IArgTestModeDurationPayload = ValueType;
            export interface IArgTestModeDuration extends IArgument { value: IArgTestModeDurationPayload }
            export type IArgZoneIdPayload = ValueType;
            export interface IArgZoneId extends IAttribute { value: IArgZoneIdPayload }
            export type IArgZoneStatePayload = ValueType;
            export interface IArgZoneState extends IAttribute { value: IArgZoneStatePayload }
            export type IArgZoneStatusPayload = ValueType;
            export interface IArgZoneStatus extends IAttribute { value: IArgZoneStatusPayload }
            export type IArgZoneTypePayload = ValueType;
            export interface IArgZoneType extends IAttribute { value: IArgZoneTypePayload }    }

    export namespace IIkea {
        // noinspection ES6UnusedImports
        import IArgument = ZigBee.IArgument;
        // noinspection ES6UnusedImports
        import IAttribute = ZigBee.IAttribute;
        // noinspection ES6UnusedImports
        import ValueType = ZigBee.ValueType;
        
        export namespace IkeaAirQuality {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

            export type IArgVocIndexPayload = ValueType;
            export interface IArgVocIndex extends IAttribute { value: IArgVocIndexPayload }    }

    export namespace ILighting {
        // noinspection ES6UnusedImports
        import IArgument = ZigBee.IArgument;
        // noinspection ES6UnusedImports
        import IAttribute = ZigBee.IAttribute;
        // noinspection ES6UnusedImports
        import ValueType = ZigBee.ValueType;
        
        export namespace ColorControl {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
            export type ICmdMoveToHuePayload = { Hue?: IArgHuePayload, MoveDirection?: IArgMoveDirectionPayload, TransitionTime?: IArgTransitionTimePayload, }
            export interface ICmdMoveToHue extends ICommand { value: ICmdMoveToHuePayload }
            export type ICmdMoveHuePayload = { MoveMode?: IArgMoveModePayload, Rate?: IArgRatePayload, }
            export interface ICmdMoveHue extends ICommand { value: ICmdMoveHuePayload }
            export type ICmdStepHuePayload = { StepMode?: IArgStepModePayload, StepSize?: IArgStepSizePayload, TransitionTime?: IArgTransitionTimePayload, }
            export interface ICmdStepHue extends ICommand { value: ICmdStepHuePayload }
            export type ICmdMoveToSaturationPayload = { Saturation?: IArgSaturationPayload, TransitionTime?: IArgTransitionTimePayload, }
            export interface ICmdMoveToSaturation extends ICommand { value: ICmdMoveToSaturationPayload }
            export type ICmdMoveSaturationPayload = { MoveMode?: IArgMoveModePayload, Rate?: IArgRatePayload, }
            export interface ICmdMoveSaturation extends ICommand { value: ICmdMoveSaturationPayload }
            export type ICmdStepSaturationPayload = { StepMode?: IArgStepModePayload, StepSize?: IArgStepSizePayload, TransitionTime?: IArgTransitionTimePayload, }
            export interface ICmdStepSaturation extends ICommand { value: ICmdStepSaturationPayload }
            export type ICmdMoveToHueAndSaturationPayload = { Hue?: IArgHuePayload, Saturation?: IArgSaturationPayload, TransitionTime?: IArgTransitionTimePayload, }
            export interface ICmdMoveToHueAndSaturation extends ICommand { value: ICmdMoveToHueAndSaturationPayload }
            export type ICmdMoveToColorPayload = { ColorX?: IArgColorXPayload, ColorY?: IArgColorYPayload, TransitionTime?: IArgTransitionTimePayload, }
            export interface ICmdMoveToColor extends ICommand { value: ICmdMoveToColorPayload }
            export type ICmdMoveColorPayload = { RateX?: IArgRateXPayload, RateY?: IArgRateYPayload, }
            export interface ICmdMoveColor extends ICommand { value: ICmdMoveColorPayload }
            export type ICmdStepColorPayload = { StepX?: IArgStepXPayload, StepY?: IArgStepYPayload, TransitionTime?: IArgTransitionTimePayload, }
            export interface ICmdStepColor extends ICommand { value: ICmdStepColorPayload }
            export type ICmdMoveToColorTemperaturePayload = { ColorTemperature?: IArgColorTemperaturePayload, TransitionTime?: IArgTransitionTimePayload, }
            export interface ICmdMoveToColorTemperature extends ICommand { value: ICmdMoveToColorTemperaturePayload }
            export type ICmdEnhancedMoveToHuePayload = { EnhancedHue?: IArgEnhancedHuePayload, MoveDirection?: IArgMoveDirectionPayload, TransitionTime?: IArgTransitionTimePayload, }
            export interface ICmdEnhancedMoveToHue extends ICommand { value: ICmdEnhancedMoveToHuePayload }
            export type ICmdEnhancedMoveHuePayload = { MoveMode?: IArgMoveModePayload, Rate?: IArgRatePayload, }
            export interface ICmdEnhancedMoveHue extends ICommand { value: ICmdEnhancedMoveHuePayload }
            export type ICmdEnhancedStepHuePayload = { StepMode?: IArgStepModePayload, StepSize?: IArgStepSizePayload, TransitionTime?: IArgTransitionTimePayload, }
            export interface ICmdEnhancedStepHue extends ICommand { value: ICmdEnhancedStepHuePayload }
            export type ICmdEnhancedMoveToHueAndSaturationPayload = { EnhancedHue?: IArgEnhancedHuePayload, Saturation?: IArgSaturationPayload, TransitionTime?: IArgTransitionTimePayload, }
            export interface ICmdEnhancedMoveToHueAndSaturation extends ICommand { value: ICmdEnhancedMoveToHueAndSaturationPayload }
            export type ICmdColorLoopSetPayload = { UpdateFlags?: IArgUpdateFlagsPayload, Action?: IArgActionPayload, HueDirection?: IArgHueDirectionPayload, Time?: IArgTimePayload, EnhancedHue?: IArgEnhancedHuePayload, }
            export interface ICmdColorLoopSet extends ICommand { value: ICmdColorLoopSetPayload }
            export type ICmdStopMoveStepPayload = { }
            export interface ICmdStopMoveStep extends ICommand { value: ICmdStopMoveStepPayload }
            export type ICmdMoveColorTemperaturePayload = { MoveMode?: IArgMoveModePayload, Rate?: IArgRatePayload, ColorTemperatureMin?: IArgColorTemperatureMinPayload, ColorTemperatureMax?: IArgColorTemperatureMaxPayload, }
            export interface ICmdMoveColorTemperature extends ICommand { value: ICmdMoveColorTemperaturePayload }
            export type ICmdStepColorTemperaturePayload = { StepMode?: IArgStepModePayload, StepSize?: IArgStepSizePayload, TransitionTime?: IArgTransitionTimePayload, ColorTemperatureMinMireds?: IArgColorTemperatureMinMiredsPayload, ColorTemperatureMaxMireds?: IArgColorTemperatureMaxMiredsPayload, }
            export interface ICmdStepColorTemperature extends ICommand { value: ICmdStepColorTemperaturePayload }
        }

        export namespace BallastConfiguration {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

            export type IArgActionPayload = ValueType;
            export interface IArgAction extends IArgument { value: IArgActionPayload }
            export type IArgBallastFactorAdjustmentPayload = ValueType;
            export interface IArgBallastFactorAdjustment extends IAttribute { value: IArgBallastFactorAdjustmentPayload }
            export type IArgBallastStatusPayload = ValueType;
            export interface IArgBallastStatus extends IAttribute { value: IArgBallastStatusPayload }
            export type IArgColorModePayload = ValueType;
            export interface IArgColorMode extends IAttribute { value: IArgColorModePayload }
            export type IArgColorPointBlueXPayload = ValueType;
            export interface IArgColorPointBlueX extends IAttribute { value: IArgColorPointBlueXPayload }
            export type IArgColorPointBlueYPayload = ValueType;
            export interface IArgColorPointBlueY extends IAttribute { value: IArgColorPointBlueYPayload }
            export type IArgColorPointBlueIntensityPayload = ValueType;
            export interface IArgColorPointBlueIntensity extends IAttribute { value: IArgColorPointBlueIntensityPayload }
            export type IArgColorPointGreenXPayload = ValueType;
            export interface IArgColorPointGreenX extends IAttribute { value: IArgColorPointGreenXPayload }
            export type IArgColorPointGreenYPayload = ValueType;
            export interface IArgColorPointGreenY extends IAttribute { value: IArgColorPointGreenYPayload }
            export type IArgColorPointGreenIntensityPayload = ValueType;
            export interface IArgColorPointGreenIntensity extends IAttribute { value: IArgColorPointGreenIntensityPayload }
            export type IArgColorPointRedXPayload = ValueType;
            export interface IArgColorPointRedX extends IAttribute { value: IArgColorPointRedXPayload }
            export type IArgColorPointRedYPayload = ValueType;
            export interface IArgColorPointRedY extends IAttribute { value: IArgColorPointRedYPayload }
            export type IArgColorPointRedIntensityPayload = ValueType;
            export interface IArgColorPointRedIntensity extends IAttribute { value: IArgColorPointRedIntensityPayload }
            export type IArgColorTemperaturePayload = ValueType;
            export interface IArgColorTemperature extends IArgument { value: IArgColorTemperaturePayload }
            export type IArgColorTemperatureMaxPayload = ValueType;
            export interface IArgColorTemperatureMax extends IArgument { value: IArgColorTemperatureMaxPayload }
            export type IArgColorTemperatureMaxMiredsPayload = ValueType;
            export interface IArgColorTemperatureMaxMireds extends IArgument { value: IArgColorTemperatureMaxMiredsPayload }
            export type IArgColorTemperatureMinPayload = ValueType;
            export interface IArgColorTemperatureMin extends IArgument { value: IArgColorTemperatureMinPayload }
            export type IArgColorTemperatureMinMiredsPayload = ValueType;
            export interface IArgColorTemperatureMinMireds extends IArgument { value: IArgColorTemperatureMinMiredsPayload }
            export type IArgColorXPayload = ValueType;
            export interface IArgColorX extends IArgument { value: IArgColorXPayload }
            export type IArgColorYPayload = ValueType;
            export interface IArgColorY extends IArgument { value: IArgColorYPayload }
            export type IArgColorCapabilitiesPayload = ValueType;
            export interface IArgColorCapabilities extends IAttribute { value: IArgColorCapabilitiesPayload }
            export type IArgColorControlOptionsPayload = ValueType;
            export interface IArgColorControlOptions extends IAttribute { value: IArgColorControlOptionsPayload }
            export type IArgColorLoopActivePayload = ValueType;
            export interface IArgColorLoopActive extends IAttribute { value: IArgColorLoopActivePayload }
            export type IArgColorLoopDirectionPayload = ValueType;
            export interface IArgColorLoopDirection extends IAttribute { value: IArgColorLoopDirectionPayload }
            export type IArgColorLoopStartEnhancedHuePayload = ValueType;
            export interface IArgColorLoopStartEnhancedHue extends IAttribute { value: IArgColorLoopStartEnhancedHuePayload }
            export type IArgColorLoopStoredEnhancedHuePayload = ValueType;
            export interface IArgColorLoopStoredEnhancedHue extends IAttribute { value: IArgColorLoopStoredEnhancedHuePayload }
            export type IArgColorLoopTimePayload = ValueType;
            export interface IArgColorLoopTime extends IAttribute { value: IArgColorLoopTimePayload }
            export type IArgColorTemperatureMiredsPayload = ValueType;
            export interface IArgColorTemperatureMireds extends IAttribute { value: IArgColorTemperatureMiredsPayload }
            export type IArgColorTemperaturePhysicalMaxMiredsPayload = ValueType;
            export interface IArgColorTemperaturePhysicalMaxMireds extends IAttribute { value: IArgColorTemperaturePhysicalMaxMiredsPayload }
            export type IArgColorTemperaturePhysicalMinMiredsPayload = ValueType;
            export interface IArgColorTemperaturePhysicalMinMireds extends IAttribute { value: IArgColorTemperaturePhysicalMinMiredsPayload }
            export type IArgCompensationTextPayload = ValueType;
            export interface IArgCompensationText extends IAttribute { value: IArgCompensationTextPayload }
            export type IArgCoupleColorTempToLevelMinMiredsPayload = ValueType;
            export interface IArgCoupleColorTempToLevelMinMireds extends IAttribute { value: IArgCoupleColorTempToLevelMinMiredsPayload }
            export type IArgCurrentXPayload = ValueType;
            export interface IArgCurrentX extends IAttribute { value: IArgCurrentXPayload }
            export type IArgCurrentYPayload = ValueType;
            export interface IArgCurrentY extends IAttribute { value: IArgCurrentYPayload }
            export type IArgCurrentHuePayload = ValueType;
            export interface IArgCurrentHue extends IAttribute { value: IArgCurrentHuePayload }
            export type IArgCurrentSaturationPayload = ValueType;
            export interface IArgCurrentSaturation extends IAttribute { value: IArgCurrentSaturationPayload }
            export type IArgDriftCompensationPayload = ValueType;
            export interface IArgDriftCompensation extends IAttribute { value: IArgDriftCompensationPayload }
            export type IArgEnhancedHuePayload = ValueType;
            export interface IArgEnhancedHue extends IArgument { value: IArgEnhancedHuePayload }
            export type IArgEnhancedColorModePayload = ValueType;
            export interface IArgEnhancedColorMode extends IAttribute { value: IArgEnhancedColorModePayload }
            export type IArgEnhancedCurrentHuePayload = ValueType;
            export interface IArgEnhancedCurrentHue extends IAttribute { value: IArgEnhancedCurrentHuePayload }
            export type IArgHuePayload = ValueType;
            export interface IArgHue extends IArgument { value: IArgHuePayload }
            export type IArgHueDirectionPayload = ValueType;
            export interface IArgHueDirection extends IArgument { value: IArgHueDirectionPayload }
            export type IArgIntrinsicBallastFactorPayload = ValueType;
            export interface IArgIntrinsicBallastFactor extends IAttribute { value: IArgIntrinsicBallastFactorPayload }
            export type IArgLampAlarmModePayload = ValueType;
            export interface IArgLampAlarmMode extends IAttribute { value: IArgLampAlarmModePayload }
            export type IArgLampBurnHoursPayload = ValueType;
            export interface IArgLampBurnHours extends IAttribute { value: IArgLampBurnHoursPayload }
            export type IArgLampBurnHoursTripPointPayload = ValueType;
            export interface IArgLampBurnHoursTripPoint extends IAttribute { value: IArgLampBurnHoursTripPointPayload }
            export type IArgLampManufacturerPayload = ValueType;
            export interface IArgLampManufacturer extends IAttribute { value: IArgLampManufacturerPayload }
            export type IArgLampQuantityPayload = ValueType;
            export interface IArgLampQuantity extends IAttribute { value: IArgLampQuantityPayload }
            export type IArgLampRatedHoursPayload = ValueType;
            export interface IArgLampRatedHours extends IAttribute { value: IArgLampRatedHoursPayload }
            export type IArgLampTypePayload = ValueType;
            export interface IArgLampType extends IAttribute { value: IArgLampTypePayload }
            export type IArgMaxLevelPayload = ValueType;
            export interface IArgMaxLevel extends IAttribute { value: IArgMaxLevelPayload }
            export type IArgMinLevelPayload = ValueType;
            export interface IArgMinLevel extends IAttribute { value: IArgMinLevelPayload }
            export type IArgMoveDirectionPayload = ValueType;
            export interface IArgMoveDirection extends IArgument { value: IArgMoveDirectionPayload }
            export type IArgMoveModePayload = ValueType;
            export interface IArgMoveMode extends IArgument { value: IArgMoveModePayload }
            export type IArgNumberOfPrimariesPayload = ValueType;
            export interface IArgNumberOfPrimaries extends IAttribute { value: IArgNumberOfPrimariesPayload }
            export type IArgPhysicalMaxLevelPayload = ValueType;
            export interface IArgPhysicalMaxLevel extends IAttribute { value: IArgPhysicalMaxLevelPayload }
            export type IArgPhysicalMinLevelPayload = ValueType;
            export interface IArgPhysicalMinLevel extends IAttribute { value: IArgPhysicalMinLevelPayload }
            export type IArgPowerOnLevelPayload = ValueType;
            export interface IArgPowerOnLevel extends IAttribute { value: IArgPowerOnLevelPayload }
            export type IArgPowerOnTimePayload = ValueType;
            export interface IArgPowerOnTime extends IAttribute { value: IArgPowerOnTimePayload }
            export type IArgPowerOnColorTemperaturePayload = ValueType;
            export interface IArgPowerOnColorTemperature extends IAttribute { value: IArgPowerOnColorTemperaturePayload }
            export type IArgPrimary1XPayload = ValueType;
            export interface IArgPrimary1X extends IAttribute { value: IArgPrimary1XPayload }
            export type IArgPrimary1YPayload = ValueType;
            export interface IArgPrimary1Y extends IAttribute { value: IArgPrimary1YPayload }
            export type IArgPrimary1IntensityPayload = ValueType;
            export interface IArgPrimary1Intensity extends IAttribute { value: IArgPrimary1IntensityPayload }
            export type IArgPrimary2XPayload = ValueType;
            export interface IArgPrimary2X extends IAttribute { value: IArgPrimary2XPayload }
            export type IArgPrimary2YPayload = ValueType;
            export interface IArgPrimary2Y extends IAttribute { value: IArgPrimary2YPayload }
            export type IArgPrimary2IntensityPayload = ValueType;
            export interface IArgPrimary2Intensity extends IAttribute { value: IArgPrimary2IntensityPayload }
            export type IArgPrimary3XPayload = ValueType;
            export interface IArgPrimary3X extends IAttribute { value: IArgPrimary3XPayload }
            export type IArgPrimary3YPayload = ValueType;
            export interface IArgPrimary3Y extends IAttribute { value: IArgPrimary3YPayload }
            export type IArgPrimary3IntensityPayload = ValueType;
            export interface IArgPrimary3Intensity extends IAttribute { value: IArgPrimary3IntensityPayload }
            export type IArgPrimary4XPayload = ValueType;
            export interface IArgPrimary4X extends IAttribute { value: IArgPrimary4XPayload }
            export type IArgPrimary4YPayload = ValueType;
            export interface IArgPrimary4Y extends IAttribute { value: IArgPrimary4YPayload }
            export type IArgPrimary4IntensityPayload = ValueType;
            export interface IArgPrimary4Intensity extends IAttribute { value: IArgPrimary4IntensityPayload }
            export type IArgPrimary5XPayload = ValueType;
            export interface IArgPrimary5X extends IAttribute { value: IArgPrimary5XPayload }
            export type IArgPrimary5YPayload = ValueType;
            export interface IArgPrimary5Y extends IAttribute { value: IArgPrimary5YPayload }
            export type IArgPrimary5IntensityPayload = ValueType;
            export interface IArgPrimary5Intensity extends IAttribute { value: IArgPrimary5IntensityPayload }
            export type IArgPrimary6XPayload = ValueType;
            export interface IArgPrimary6X extends IAttribute { value: IArgPrimary6XPayload }
            export type IArgPrimary6YPayload = ValueType;
            export interface IArgPrimary6Y extends IAttribute { value: IArgPrimary6YPayload }
            export type IArgPrimary6IntensityPayload = ValueType;
            export interface IArgPrimary6Intensity extends IAttribute { value: IArgPrimary6IntensityPayload }
            export type IArgRatePayload = ValueType;
            export interface IArgRate extends IArgument { value: IArgRatePayload }
            export type IArgRateXPayload = ValueType;
            export interface IArgRateX extends IArgument { value: IArgRateXPayload }
            export type IArgRateYPayload = ValueType;
            export interface IArgRateY extends IArgument { value: IArgRateYPayload }
            export type IArgRemainingTimePayload = ValueType;
            export interface IArgRemainingTime extends IAttribute { value: IArgRemainingTimePayload }
            export type IArgSaturationPayload = ValueType;
            export interface IArgSaturation extends IArgument { value: IArgSaturationPayload }
            export type IArgStepXPayload = ValueType;
            export interface IArgStepX extends IArgument { value: IArgStepXPayload }
            export type IArgStepYPayload = ValueType;
            export interface IArgStepY extends IArgument { value: IArgStepYPayload }
            export type IArgStepModePayload = ValueType;
            export interface IArgStepMode extends IArgument { value: IArgStepModePayload }
            export type IArgStepSizePayload = ValueType;
            export interface IArgStepSize extends IArgument { value: IArgStepSizePayload }
            export type IArgTimePayload = ValueType;
            export interface IArgTime extends IArgument { value: IArgTimePayload }
            export type IArgTransitionTimePayload = ValueType;
            export interface IArgTransitionTime extends IArgument { value: IArgTransitionTimePayload }
            export type IArgUpdateFlagsPayload = ValueType;
            export interface IArgUpdateFlags extends IArgument { value: IArgUpdateFlagsPayload }
            export type IArgWhitePointXPayload = ValueType;
            export interface IArgWhitePointX extends IAttribute { value: IArgWhitePointXPayload }
            export type IArgWhitePointYPayload = ValueType;
            export interface IArgWhitePointY extends IAttribute { value: IArgWhitePointYPayload }    }

    export namespace IMeasurementAndSensing {
        // noinspection ES6UnusedImports
        import IArgument = ZigBee.IArgument;
        // noinspection ES6UnusedImports
        import IAttribute = ZigBee.IAttribute;
        // noinspection ES6UnusedImports
        import ValueType = ZigBee.ValueType;
        
        export namespace IlluminanceMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace IlluminanceLevelSensing {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace TemperatureMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace PressureMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace FlowMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace RelativeHumidityMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace OccupancySensing {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace Pm25Measurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace ElectricalMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
            export type ICmdGetProfileInfoPayload = { }
            export interface ICmdGetProfileInfo extends ICommand { value: ICmdGetProfileInfoPayload }
            export type ICmdGetMeasurementProfilePayload = { AttributeId?: IArgAttributeIdPayload, StartTime?: IArgStartTimePayload, NumberOfIntervals?: IArgNumberOfIntervalsPayload, }
            export interface ICmdGetMeasurementProfile extends ICommand { value: ICmdGetMeasurementProfilePayload }
            export type ICmdGetProfileInfoResponsePayload = { ProfileCount?: IArgProfileCountPayload, ProfileIntervalPeriod?: IArgProfileIntervalPeriodPayload, MaxNumberOfIntervals?: IArgMaxNumberOfIntervalsPayload, ListOfAttributes?: IArgListOfAttributesPayload, }
            export interface ICmdGetProfileInfoResponse extends ICommand { value: ICmdGetProfileInfoResponsePayload }
            export type ICmdGetMeasurementProfileResponsePayload = { StartTime?: IArgStartTimePayload, MeasurementResponseStatus?: IArgMeasurementResponseStatusPayload, ProfileIntervalPeriod?: IArgProfileIntervalPeriodPayload, NumberOfIntervals?: IArgNumberOfIntervalsPayload, AttributeId?: IArgAttributeIdPayload, Intervals?: IArgIntervalsPayload, }
            export interface ICmdGetMeasurementProfileResponse extends ICommand { value: ICmdGetMeasurementProfileResponsePayload }
        }

        export namespace CarbonMonoxideMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace CarbonDioxideMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace EthyleneMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace EthyleneOxideMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace HydrogenMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace HydrogenSulfideMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace NitricOxideMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace NitrogenDioxideMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace OxygenMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace OzoneMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace SulfurDioxideMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace DissolvedOxygenMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace BromateMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace ChloraminesMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace ChlorineMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace FecalColiformEColiMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace FluorideMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace HaloaceticAcidsMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace TotalTrihalomethanesMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace TotalColiformBacteriaMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace TurbidityMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace CopperMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace LeadMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace ManganeseMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace SulfateMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace BromodichloromethaneMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace BromoformMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace ChlorodibromomethaneMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace ChloroformMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace SodiumMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace FormaldehydeMeasurement {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

            export type IArgAcActivePowerOverloadPayload = ValueType;
            export interface IArgAcActivePowerOverload extends IAttribute { value: IArgAcActivePowerOverloadPayload }
            export type IArgAcAlarmsMaskPayload = ValueType;
            export interface IArgAcAlarmsMask extends IAttribute { value: IArgAcAlarmsMaskPayload }
            export type IArgAcCurrentDivisorPayload = ValueType;
            export interface IArgAcCurrentDivisor extends IAttribute { value: IArgAcCurrentDivisorPayload }
            export type IArgAcCurrentMultiplierPayload = ValueType;
            export interface IArgAcCurrentMultiplier extends IAttribute { value: IArgAcCurrentMultiplierPayload }
            export type IArgAcCurrentOverloadPayload = ValueType;
            export interface IArgAcCurrentOverload extends IAttribute { value: IArgAcCurrentOverloadPayload }
            export type IArgAcFrequencyPayload = ValueType;
            export interface IArgAcFrequency extends IAttribute { value: IArgAcFrequencyPayload }
            export type IArgAcFrequencyDivisorPayload = ValueType;
            export interface IArgAcFrequencyDivisor extends IAttribute { value: IArgAcFrequencyDivisorPayload }
            export type IArgAcFrequencyMaxPayload = ValueType;
            export interface IArgAcFrequencyMax extends IAttribute { value: IArgAcFrequencyMaxPayload }
            export type IArgAcFrequencyMinPayload = ValueType;
            export interface IArgAcFrequencyMin extends IAttribute { value: IArgAcFrequencyMinPayload }
            export type IArgAcFrequencyMultiplierPayload = ValueType;
            export interface IArgAcFrequencyMultiplier extends IAttribute { value: IArgAcFrequencyMultiplierPayload }
            export type IArgAcPowerDivisorPayload = ValueType;
            export interface IArgAcPowerDivisor extends IAttribute { value: IArgAcPowerDivisorPayload }
            export type IArgAcPowerMultiplierPayload = ValueType;
            export interface IArgAcPowerMultiplier extends IAttribute { value: IArgAcPowerMultiplierPayload }
            export type IArgAcReactivePowerOverloadPayload = ValueType;
            export interface IArgAcReactivePowerOverload extends IAttribute { value: IArgAcReactivePowerOverloadPayload }
            export type IArgAcVoltageDivisorPayload = ValueType;
            export interface IArgAcVoltageDivisor extends IAttribute { value: IArgAcVoltageDivisorPayload }
            export type IArgAcVoltageMultiplierPayload = ValueType;
            export interface IArgAcVoltageMultiplier extends IAttribute { value: IArgAcVoltageMultiplierPayload }
            export type IArgAcVoltageOverloadPayload = ValueType;
            export interface IArgAcVoltageOverload extends IAttribute { value: IArgAcVoltageOverloadPayload }
            export type IArgActiveCurrentPayload = ValueType;
            export interface IArgActiveCurrent extends IAttribute { value: IArgActiveCurrentPayload }
            export type IArgActiveCurrentPhBPayload = ValueType;
            export interface IArgActiveCurrentPhB extends IAttribute { value: IArgActiveCurrentPhBPayload }
            export type IArgActiveCurrentPhCPayload = ValueType;
            export interface IArgActiveCurrentPhC extends IAttribute { value: IArgActiveCurrentPhCPayload }
            export type IArgActivePowerPayload = ValueType;
            export interface IArgActivePower extends IAttribute { value: IArgActivePowerPayload }
            export type IArgActivePowerMaxPayload = ValueType;
            export interface IArgActivePowerMax extends IAttribute { value: IArgActivePowerMaxPayload }
            export type IArgActivePowerMaxPhBPayload = ValueType;
            export interface IArgActivePowerMaxPhB extends IAttribute { value: IArgActivePowerMaxPhBPayload }
            export type IArgActivePowerMaxPhCPayload = ValueType;
            export interface IArgActivePowerMaxPhC extends IAttribute { value: IArgActivePowerMaxPhCPayload }
            export type IArgActivePowerMinPayload = ValueType;
            export interface IArgActivePowerMin extends IAttribute { value: IArgActivePowerMinPayload }
            export type IArgActivePowerMinPhBPayload = ValueType;
            export interface IArgActivePowerMinPhB extends IAttribute { value: IArgActivePowerMinPhBPayload }
            export type IArgActivePowerMinPhCPayload = ValueType;
            export interface IArgActivePowerMinPhC extends IAttribute { value: IArgActivePowerMinPhCPayload }
            export type IArgActivePowerPhBPayload = ValueType;
            export interface IArgActivePowerPhB extends IAttribute { value: IArgActivePowerPhBPayload }
            export type IArgActivePowerPhCPayload = ValueType;
            export interface IArgActivePowerPhC extends IAttribute { value: IArgActivePowerPhCPayload }
            export type IArgApparentPowerPayload = ValueType;
            export interface IArgApparentPower extends IAttribute { value: IArgApparentPowerPayload }
            export type IArgApparentPowerPhBPayload = ValueType;
            export interface IArgApparentPowerPhB extends IAttribute { value: IArgApparentPowerPhBPayload }
            export type IArgApparentPowerPhCPayload = ValueType;
            export interface IArgApparentPowerPhC extends IAttribute { value: IArgApparentPowerPhCPayload }
            export type IArgAttributeIdPayload = ValueType;
            export interface IArgAttributeId extends IArgument { value: IArgAttributeIdPayload }
            export type IArgAverageRmsOverVoltageCounterPayload = ValueType;
            export interface IArgAverageRmsOverVoltageCounter extends IAttribute { value: IArgAverageRmsOverVoltageCounterPayload }
            export type IArgAverageRmsOverVoltageCounterPhBPayload = ValueType;
            export interface IArgAverageRmsOverVoltageCounterPhB extends IAttribute { value: IArgAverageRmsOverVoltageCounterPhBPayload }
            export type IArgAverageRmsOverVoltageCounterPhCPayload = ValueType;
            export interface IArgAverageRmsOverVoltageCounterPhC extends IAttribute { value: IArgAverageRmsOverVoltageCounterPhCPayload }
            export type IArgAverageRmsUnderVoltageCounterPayload = ValueType;
            export interface IArgAverageRmsUnderVoltageCounter extends IAttribute { value: IArgAverageRmsUnderVoltageCounterPayload }
            export type IArgAverageRmsUnderVoltageCounterPhBPayload = ValueType;
            export interface IArgAverageRmsUnderVoltageCounterPhB extends IAttribute { value: IArgAverageRmsUnderVoltageCounterPhBPayload }
            export type IArgAverageRmsUnderVoltageCounterPhCPayload = ValueType;
            export interface IArgAverageRmsUnderVoltageCounterPhC extends IAttribute { value: IArgAverageRmsUnderVoltageCounterPhCPayload }
            export type IArgAverageRmsVoltageMeasurementPeriodPayload = ValueType;
            export interface IArgAverageRmsVoltageMeasurementPeriod extends IAttribute { value: IArgAverageRmsVoltageMeasurementPeriodPayload }
            export type IArgAverageRmsVoltageMeasurementPeriodPhBPayload = ValueType;
            export interface IArgAverageRmsVoltageMeasurementPeriodPhB extends IAttribute { value: IArgAverageRmsVoltageMeasurementPeriodPhBPayload }
            export type IArgAverageRmsVoltageMeasurementPeriodPhCPayload = ValueType;
            export interface IArgAverageRmsVoltageMeasurementPeriodPhC extends IAttribute { value: IArgAverageRmsVoltageMeasurementPeriodPhCPayload }
            export type IArgAverageRmsOverVoltagePayload = ValueType;
            export interface IArgAverageRmsOverVoltage extends IAttribute { value: IArgAverageRmsOverVoltagePayload }
            export type IArgAverageRmsUnderVoltagePayload = ValueType;
            export interface IArgAverageRmsUnderVoltage extends IAttribute { value: IArgAverageRmsUnderVoltagePayload }
            export type IArgConcentrationTolerancePayload = ValueType;
            export interface IArgConcentrationTolerance extends IAttribute { value: IArgConcentrationTolerancePayload }
            export type IArgDcCurrentPayload = ValueType;
            export interface IArgDcCurrent extends IAttribute { value: IArgDcCurrentPayload }
            export type IArgDcCurrentDivisorPayload = ValueType;
            export interface IArgDcCurrentDivisor extends IAttribute { value: IArgDcCurrentDivisorPayload }
            export type IArgDcCurrentMaxPayload = ValueType;
            export interface IArgDcCurrentMax extends IAttribute { value: IArgDcCurrentMaxPayload }
            export type IArgDcCurrentMinPayload = ValueType;
            export interface IArgDcCurrentMin extends IAttribute { value: IArgDcCurrentMinPayload }
            export type IArgDcCurrentMultiplierPayload = ValueType;
            export interface IArgDcCurrentMultiplier extends IAttribute { value: IArgDcCurrentMultiplierPayload }
            export type IArgDcCurrentOverloadPayload = ValueType;
            export interface IArgDcCurrentOverload extends IAttribute { value: IArgDcCurrentOverloadPayload }
            export type IArgDcOverloadAlarmsMaskPayload = ValueType;
            export interface IArgDcOverloadAlarmsMask extends IAttribute { value: IArgDcOverloadAlarmsMaskPayload }
            export type IArgDcPowerPayload = ValueType;
            export interface IArgDcPower extends IAttribute { value: IArgDcPowerPayload }
            export type IArgDcPowerDivisorPayload = ValueType;
            export interface IArgDcPowerDivisor extends IAttribute { value: IArgDcPowerDivisorPayload }
            export type IArgDcPowerMaxPayload = ValueType;
            export interface IArgDcPowerMax extends IAttribute { value: IArgDcPowerMaxPayload }
            export type IArgDcPowerMinPayload = ValueType;
            export interface IArgDcPowerMin extends IAttribute { value: IArgDcPowerMinPayload }
            export type IArgDcPowerMultiplierPayload = ValueType;
            export interface IArgDcPowerMultiplier extends IAttribute { value: IArgDcPowerMultiplierPayload }
            export type IArgDcVoltagePayload = ValueType;
            export interface IArgDcVoltage extends IAttribute { value: IArgDcVoltagePayload }
            export type IArgDcVoltageDivisorPayload = ValueType;
            export interface IArgDcVoltageDivisor extends IAttribute { value: IArgDcVoltageDivisorPayload }
            export type IArgDcVoltageMaxPayload = ValueType;
            export interface IArgDcVoltageMax extends IAttribute { value: IArgDcVoltageMaxPayload }
            export type IArgDcVoltageMinPayload = ValueType;
            export interface IArgDcVoltageMin extends IAttribute { value: IArgDcVoltageMinPayload }
            export type IArgDcVoltageMultiplierPayload = ValueType;
            export interface IArgDcVoltageMultiplier extends IAttribute { value: IArgDcVoltageMultiplierPayload }
            export type IArgDcVoltageOverloadPayload = ValueType;
            export interface IArgDcVoltageOverload extends IAttribute { value: IArgDcVoltageOverloadPayload }
            export type IArgElectricalMeasurementTypePayload = ValueType;
            export interface IArgElectricalMeasurementType extends IAttribute { value: IArgElectricalMeasurementTypePayload }
            export type IArgHarmonicCurrentMultiplierPayload = ValueType;
            export interface IArgHarmonicCurrentMultiplier extends IAttribute { value: IArgHarmonicCurrentMultiplierPayload }
            export type IArgIlluminanceLightSensorTypePayload = ValueType;
            export interface IArgIlluminanceLightSensorType extends IAttribute { value: IArgIlluminanceLightSensorTypePayload }
            export type IArgIlluminanceTargetLevelPayload = ValueType;
            export interface IArgIlluminanceTargetLevel extends IAttribute { value: IArgIlluminanceTargetLevelPayload }
            export type IArgIntervalsPayload = ValueType;
            export interface IArgIntervals extends IArgument { value: IArgIntervalsPayload }
            export type IArgLevelStatusPayload = ValueType;
            export interface IArgLevelStatus extends IAttribute { value: IArgLevelStatusPayload }
            export type IArgLightSensorTypePayload = ValueType;
            export interface IArgLightSensorType extends IAttribute { value: IArgLightSensorTypePayload }
            export type IArgLineCurrentPayload = ValueType;
            export interface IArgLineCurrent extends IAttribute { value: IArgLineCurrentPayload }
            export type IArgLineCurrentPhBPayload = ValueType;
            export interface IArgLineCurrentPhB extends IAttribute { value: IArgLineCurrentPhBPayload }
            export type IArgLineCurrentPhCPayload = ValueType;
            export interface IArgLineCurrentPhC extends IAttribute { value: IArgLineCurrentPhCPayload }
            export type IArgListOfAttributesPayload = ValueType;
            export interface IArgListOfAttributes extends IArgument { value: IArgListOfAttributesPayload }
            export type IArgMaxMeasuredConcentrationPayload = ValueType;
            export interface IArgMaxMeasuredConcentration extends IAttribute { value: IArgMaxMeasuredConcentrationPayload }
            export type IArgMaxMeasuredFlowPayload = ValueType;
            export interface IArgMaxMeasuredFlow extends IAttribute { value: IArgMaxMeasuredFlowPayload }
            export type IArgMaxMeasuredIlluminancePayload = ValueType;
            export interface IArgMaxMeasuredIlluminance extends IAttribute { value: IArgMaxMeasuredIlluminancePayload }
            export type IArgMaxMeasuredPm25Payload = ValueType;
            export interface IArgMaxMeasuredPm25 extends IAttribute { value: IArgMaxMeasuredPm25Payload }
            export type IArgMaxMeasuredPressurePayload = ValueType;
            export interface IArgMaxMeasuredPressure extends IAttribute { value: IArgMaxMeasuredPressurePayload }
            export type IArgMaxMeasuredRelativeHumidityPayload = ValueType;
            export interface IArgMaxMeasuredRelativeHumidity extends IAttribute { value: IArgMaxMeasuredRelativeHumidityPayload }
            export type IArgMaxMeasuredTemperaturePayload = ValueType;
            export interface IArgMaxMeasuredTemperature extends IAttribute { value: IArgMaxMeasuredTemperaturePayload }
            export type IArgMaxNumberOfIntervalsPayload = ValueType;
            export interface IArgMaxNumberOfIntervals extends IArgument { value: IArgMaxNumberOfIntervalsPayload }
            export type IArgMeasured11ThHarmonicCurrentPayload = ValueType;
            export interface IArgMeasured11ThHarmonicCurrent extends IAttribute { value: IArgMeasured11ThHarmonicCurrentPayload }
            export type IArgMeasured1StHarmonicCurrentPayload = ValueType;
            export interface IArgMeasured1StHarmonicCurrent extends IAttribute { value: IArgMeasured1StHarmonicCurrentPayload }
            export type IArgMeasured3RdHarmonicCurrentPayload = ValueType;
            export interface IArgMeasured3RdHarmonicCurrent extends IAttribute { value: IArgMeasured3RdHarmonicCurrentPayload }
            export type IArgMeasured5ThHarmonicCurrentPayload = ValueType;
            export interface IArgMeasured5ThHarmonicCurrent extends IAttribute { value: IArgMeasured5ThHarmonicCurrentPayload }
            export type IArgMeasured7ThHarmonicCurrentPayload = ValueType;
            export interface IArgMeasured7ThHarmonicCurrent extends IAttribute { value: IArgMeasured7ThHarmonicCurrentPayload }
            export type IArgMeasured9ThHarmonicCurrentPayload = ValueType;
            export interface IArgMeasured9ThHarmonicCurrent extends IAttribute { value: IArgMeasured9ThHarmonicCurrentPayload }
            export type IArgMeasuredConcentrationPayload = ValueType;
            export interface IArgMeasuredConcentration extends IAttribute { value: IArgMeasuredConcentrationPayload }
            export type IArgMeasuredFlowPayload = ValueType;
            export interface IArgMeasuredFlow extends IAttribute { value: IArgMeasuredFlowPayload }
            export type IArgMeasuredIlluminancePayload = ValueType;
            export interface IArgMeasuredIlluminance extends IAttribute { value: IArgMeasuredIlluminancePayload }
            export type IArgMeasuredPm25Payload = ValueType;
            export interface IArgMeasuredPm25 extends IAttribute { value: IArgMeasuredPm25Payload }
            export type IArgMeasuredPhase11ThHarmonicCurrentPayload = ValueType;
            export interface IArgMeasuredPhase11ThHarmonicCurrent extends IAttribute { value: IArgMeasuredPhase11ThHarmonicCurrentPayload }
            export type IArgMeasuredPhase1StHarmonicCurrentPayload = ValueType;
            export interface IArgMeasuredPhase1StHarmonicCurrent extends IAttribute { value: IArgMeasuredPhase1StHarmonicCurrentPayload }
            export type IArgMeasuredPhase3RdHarmonicCurrentPayload = ValueType;
            export interface IArgMeasuredPhase3RdHarmonicCurrent extends IAttribute { value: IArgMeasuredPhase3RdHarmonicCurrentPayload }
            export type IArgMeasuredPhase5ThHarmonicCurrentPayload = ValueType;
            export interface IArgMeasuredPhase5ThHarmonicCurrent extends IAttribute { value: IArgMeasuredPhase5ThHarmonicCurrentPayload }
            export type IArgMeasuredPhase7ThHarmonicCurrentPayload = ValueType;
            export interface IArgMeasuredPhase7ThHarmonicCurrent extends IAttribute { value: IArgMeasuredPhase7ThHarmonicCurrentPayload }
            export type IArgMeasuredPhase9ThHarmonicCurrentPayload = ValueType;
            export interface IArgMeasuredPhase9ThHarmonicCurrent extends IAttribute { value: IArgMeasuredPhase9ThHarmonicCurrentPayload }
            export type IArgMeasuredPressurePayload = ValueType;
            export interface IArgMeasuredPressure extends IAttribute { value: IArgMeasuredPressurePayload }
            export type IArgMeasuredRelativeHumidityPayload = ValueType;
            export interface IArgMeasuredRelativeHumidity extends IAttribute { value: IArgMeasuredRelativeHumidityPayload }
            export type IArgMeasuredTemperaturePayload = ValueType;
            export interface IArgMeasuredTemperature extends IAttribute { value: IArgMeasuredTemperaturePayload }
            export type IArgMeasurementResponseStatusPayload = ValueType;
            export interface IArgMeasurementResponseStatus extends IArgument { value: IArgMeasurementResponseStatusPayload }
            export type IArgMinMeasuredConcentrationPayload = ValueType;
            export interface IArgMinMeasuredConcentration extends IAttribute { value: IArgMinMeasuredConcentrationPayload }
            export type IArgMinMeasuredFlowPayload = ValueType;
            export interface IArgMinMeasuredFlow extends IAttribute { value: IArgMinMeasuredFlowPayload }
            export type IArgMinMeasuredIlluminancePayload = ValueType;
            export interface IArgMinMeasuredIlluminance extends IAttribute { value: IArgMinMeasuredIlluminancePayload }
            export type IArgMinMeasuredPm25Payload = ValueType;
            export interface IArgMinMeasuredPm25 extends IAttribute { value: IArgMinMeasuredPm25Payload }
            export type IArgMinMeasuredPressurePayload = ValueType;
            export interface IArgMinMeasuredPressure extends IAttribute { value: IArgMinMeasuredPressurePayload }
            export type IArgMinMeasuredRelativeHumidityPayload = ValueType;
            export interface IArgMinMeasuredRelativeHumidity extends IAttribute { value: IArgMinMeasuredRelativeHumidityPayload }
            export type IArgMinMeasuredTemperaturePayload = ValueType;
            export interface IArgMinMeasuredTemperature extends IAttribute { value: IArgMinMeasuredTemperaturePayload }
            export type IArgNeutralCurrentPayload = ValueType;
            export interface IArgNeutralCurrent extends IAttribute { value: IArgNeutralCurrentPayload }
            export type IArgNumberOfIntervalsPayload = ValueType;
            export interface IArgNumberOfIntervals extends IArgument { value: IArgNumberOfIntervalsPayload }
            export type IArgOccupancyPayload = ValueType;
            export interface IArgOccupancy extends IAttribute { value: IArgOccupancyPayload }
            export type IArgOccupancyTypePayload = ValueType;
            export interface IArgOccupancyType extends IAttribute { value: IArgOccupancyTypePayload }
            export type IArgPirOccupiedToUnoccupiedDelayPayload = ValueType;
            export interface IArgPirOccupiedToUnoccupiedDelay extends IAttribute { value: IArgPirOccupiedToUnoccupiedDelayPayload }
            export type IArgPirUnoccupiedToOccupiedDelayPayload = ValueType;
            export interface IArgPirUnoccupiedToOccupiedDelay extends IAttribute { value: IArgPirUnoccupiedToOccupiedDelayPayload }
            export type IArgPirUnoccupiedToOccupiedThresholdPayload = ValueType;
            export interface IArgPirUnoccupiedToOccupiedThreshold extends IAttribute { value: IArgPirUnoccupiedToOccupiedThresholdPayload }
            export type IArgPm25TolerancePayload = ValueType;
            export interface IArgPm25Tolerance extends IAttribute { value: IArgPm25TolerancePayload }
            export type IArgPhaseHarmonicCurrentMultiplierPayload = ValueType;
            export interface IArgPhaseHarmonicCurrentMultiplier extends IAttribute { value: IArgPhaseHarmonicCurrentMultiplierPayload }
            export type IArgPowerDivisorPayload = ValueType;
            export interface IArgPowerDivisor extends IAttribute { value: IArgPowerDivisorPayload }
            export type IArgPowerFactorPayload = ValueType;
            export interface IArgPowerFactor extends IAttribute { value: IArgPowerFactorPayload }
            export type IArgPowerFactorPhBPayload = ValueType;
            export interface IArgPowerFactorPhB extends IAttribute { value: IArgPowerFactorPhBPayload }
            export type IArgPowerFactorPhCPayload = ValueType;
            export interface IArgPowerFactorPhC extends IAttribute { value: IArgPowerFactorPhCPayload }
            export type IArgPowerMultiplierPayload = ValueType;
            export interface IArgPowerMultiplier extends IAttribute { value: IArgPowerMultiplierPayload }
            export type IArgProfileCountPayload = ValueType;
            export interface IArgProfileCount extends IArgument { value: IArgProfileCountPayload }
            export type IArgProfileIntervalPeriodPayload = ValueType;
            export interface IArgProfileIntervalPeriod extends IArgument { value: IArgProfileIntervalPeriodPayload }
            export type IArgRmsCurrentPayload = ValueType;
            export interface IArgRmsCurrent extends IAttribute { value: IArgRmsCurrentPayload }
            export type IArgRmsCurrentMaxPayload = ValueType;
            export interface IArgRmsCurrentMax extends IAttribute { value: IArgRmsCurrentMaxPayload }
            export type IArgRmsCurrentMaxPhBPayload = ValueType;
            export interface IArgRmsCurrentMaxPhB extends IAttribute { value: IArgRmsCurrentMaxPhBPayload }
            export type IArgRmsCurrentMaxPhCPayload = ValueType;
            export interface IArgRmsCurrentMaxPhC extends IAttribute { value: IArgRmsCurrentMaxPhCPayload }
            export type IArgRmsCurrentMinPayload = ValueType;
            export interface IArgRmsCurrentMin extends IAttribute { value: IArgRmsCurrentMinPayload }
            export type IArgRmsCurrentMinPhBPayload = ValueType;
            export interface IArgRmsCurrentMinPhB extends IAttribute { value: IArgRmsCurrentMinPhBPayload }
            export type IArgRmsCurrentMinPhCPayload = ValueType;
            export interface IArgRmsCurrentMinPhC extends IAttribute { value: IArgRmsCurrentMinPhCPayload }
            export type IArgRmsCurrentPhBPayload = ValueType;
            export interface IArgRmsCurrentPhB extends IAttribute { value: IArgRmsCurrentPhBPayload }
            export type IArgRmsCurrentPhCPayload = ValueType;
            export interface IArgRmsCurrentPhC extends IAttribute { value: IArgRmsCurrentPhCPayload }
            export type IArgRmsExtremeOverVoltagePeriodPayload = ValueType;
            export interface IArgRmsExtremeOverVoltagePeriod extends IAttribute { value: IArgRmsExtremeOverVoltagePeriodPayload }
            export type IArgRmsExtremeOverVoltagePeriodPhBPayload = ValueType;
            export interface IArgRmsExtremeOverVoltagePeriodPhB extends IAttribute { value: IArgRmsExtremeOverVoltagePeriodPhBPayload }
            export type IArgRmsExtremeOverVoltagePeriodPhCPayload = ValueType;
            export interface IArgRmsExtremeOverVoltagePeriodPhC extends IAttribute { value: IArgRmsExtremeOverVoltagePeriodPhCPayload }
            export type IArgRmsExtremeUnderVoltagePeriodPayload = ValueType;
            export interface IArgRmsExtremeUnderVoltagePeriod extends IAttribute { value: IArgRmsExtremeUnderVoltagePeriodPayload }
            export type IArgRmsExtremeUnderVoltagePeriodPhBPayload = ValueType;
            export interface IArgRmsExtremeUnderVoltagePeriodPhB extends IAttribute { value: IArgRmsExtremeUnderVoltagePeriodPhBPayload }
            export type IArgRmsExtremeUnderVoltagePeriodPhCPayload = ValueType;
            export interface IArgRmsExtremeUnderVoltagePeriodPhC extends IAttribute { value: IArgRmsExtremeUnderVoltagePeriodPhCPayload }
            export type IArgRmsVoltagePayload = ValueType;
            export interface IArgRmsVoltage extends IAttribute { value: IArgRmsVoltagePayload }
            export type IArgRmsVoltageMaxPayload = ValueType;
            export interface IArgRmsVoltageMax extends IAttribute { value: IArgRmsVoltageMaxPayload }
            export type IArgRmsVoltageMaxPhBPayload = ValueType;
            export interface IArgRmsVoltageMaxPhB extends IAttribute { value: IArgRmsVoltageMaxPhBPayload }
            export type IArgRmsVoltageMaxPhCPayload = ValueType;
            export interface IArgRmsVoltageMaxPhC extends IAttribute { value: IArgRmsVoltageMaxPhCPayload }
            export type IArgRmsVoltageMinPayload = ValueType;
            export interface IArgRmsVoltageMin extends IAttribute { value: IArgRmsVoltageMinPayload }
            export type IArgRmsVoltageMinPhBPayload = ValueType;
            export interface IArgRmsVoltageMinPhB extends IAttribute { value: IArgRmsVoltageMinPhBPayload }
            export type IArgRmsVoltageMinPhCPayload = ValueType;
            export interface IArgRmsVoltageMinPhC extends IAttribute { value: IArgRmsVoltageMinPhCPayload }
            export type IArgRmsVoltagePhBPayload = ValueType;
            export interface IArgRmsVoltagePhB extends IAttribute { value: IArgRmsVoltagePhBPayload }
            export type IArgRmsVoltagePhCPayload = ValueType;
            export interface IArgRmsVoltagePhC extends IAttribute { value: IArgRmsVoltagePhCPayload }
            export type IArgRmsVoltageSagPayload = ValueType;
            export interface IArgRmsVoltageSag extends IAttribute { value: IArgRmsVoltageSagPayload }
            export type IArgRmsVoltageSagPeriodPayload = ValueType;
            export interface IArgRmsVoltageSagPeriod extends IAttribute { value: IArgRmsVoltageSagPeriodPayload }
            export type IArgRmsVoltageSagPeriodPhBPayload = ValueType;
            export interface IArgRmsVoltageSagPeriodPhB extends IAttribute { value: IArgRmsVoltageSagPeriodPhBPayload }
            export type IArgRmsVoltageSagPeriodPhCPayload = ValueType;
            export interface IArgRmsVoltageSagPeriodPhC extends IAttribute { value: IArgRmsVoltageSagPeriodPhCPayload }
            export type IArgRmsVoltageSwellPayload = ValueType;
            export interface IArgRmsVoltageSwell extends IAttribute { value: IArgRmsVoltageSwellPayload }
            export type IArgRmsVoltageSwellPeriodPayload = ValueType;
            export interface IArgRmsVoltageSwellPeriod extends IAttribute { value: IArgRmsVoltageSwellPeriodPayload }
            export type IArgRmsVoltageSwellPeriodPhBPayload = ValueType;
            export interface IArgRmsVoltageSwellPeriodPhB extends IAttribute { value: IArgRmsVoltageSwellPeriodPhBPayload }
            export type IArgRmsVoltageSwellPeriodPhCPayload = ValueType;
            export interface IArgRmsVoltageSwellPeriodPhC extends IAttribute { value: IArgRmsVoltageSwellPeriodPhCPayload }
            export type IArgRmsExtremeOverVoltagePayload = ValueType;
            export interface IArgRmsExtremeOverVoltage extends IAttribute { value: IArgRmsExtremeOverVoltagePayload }
            export type IArgRmsExtremeUnderVoltagePayload = ValueType;
            export interface IArgRmsExtremeUnderVoltage extends IAttribute { value: IArgRmsExtremeUnderVoltagePayload }
            export type IArgReactiveCurrentPayload = ValueType;
            export interface IArgReactiveCurrent extends IAttribute { value: IArgReactiveCurrentPayload }
            export type IArgReactiveCurrentPhBPayload = ValueType;
            export interface IArgReactiveCurrentPhB extends IAttribute { value: IArgReactiveCurrentPhBPayload }
            export type IArgReactiveCurrentPhCPayload = ValueType;
            export interface IArgReactiveCurrentPhC extends IAttribute { value: IArgReactiveCurrentPhCPayload }
            export type IArgReactivePowerPayload = ValueType;
            export interface IArgReactivePower extends IAttribute { value: IArgReactivePowerPayload }
            export type IArgReactivePowerPhBPayload = ValueType;
            export interface IArgReactivePowerPhB extends IAttribute { value: IArgReactivePowerPhBPayload }
            export type IArgReactivePowerPhCPayload = ValueType;
            export interface IArgReactivePowerPhC extends IAttribute { value: IArgReactivePowerPhCPayload }
            export type IArgScalePayload = ValueType;
            export interface IArgScale extends IAttribute { value: IArgScalePayload }
            export type IArgScaledMaxPressurePayload = ValueType;
            export interface IArgScaledMaxPressure extends IAttribute { value: IArgScaledMaxPressurePayload }
            export type IArgScaledMinPressurePayload = ValueType;
            export interface IArgScaledMinPressure extends IAttribute { value: IArgScaledMinPressurePayload }
            export type IArgScaledPressurePayload = ValueType;
            export interface IArgScaledPressure extends IAttribute { value: IArgScaledPressurePayload }
            export type IArgScaledTolerancePayload = ValueType;
            export interface IArgScaledTolerance extends IAttribute { value: IArgScaledTolerancePayload }
            export type IArgStartTimePayload = ValueType;
            export interface IArgStartTime extends IArgument { value: IArgStartTimePayload }
            export type IArgTolerancePayload = ValueType;
            export interface IArgTolerance extends IAttribute { value: IArgTolerancePayload }
            export type IArgTotalActivePowerPayload = ValueType;
            export interface IArgTotalActivePower extends IAttribute { value: IArgTotalActivePowerPayload }
            export type IArgTotalApparentPowerPayload = ValueType;
            export interface IArgTotalApparentPower extends IAttribute { value: IArgTotalApparentPowerPayload }
            export type IArgTotalReactivePowerPayload = ValueType;
            export interface IArgTotalReactivePower extends IAttribute { value: IArgTotalReactivePowerPayload }
            export type IArgUltrasonicOccupiedToUnoccupiedDelayPayload = ValueType;
            export interface IArgUltrasonicOccupiedToUnoccupiedDelay extends IAttribute { value: IArgUltrasonicOccupiedToUnoccupiedDelayPayload }
            export type IArgUltrasonicUnoccupiedToOccupiedDelayPayload = ValueType;
            export interface IArgUltrasonicUnoccupiedToOccupiedDelay extends IAttribute { value: IArgUltrasonicUnoccupiedToOccupiedDelayPayload }
            export type IArgUltrasonicUnoccupiedToOccupiedThresholdPayload = ValueType;
            export interface IArgUltrasonicUnoccupiedToOccupiedThreshold extends IAttribute { value: IArgUltrasonicUnoccupiedToOccupiedThresholdPayload }    }

    export namespace IOtau {
        // noinspection ES6UnusedImports
        import IArgument = ZigBee.IArgument;
        // noinspection ES6UnusedImports
        import IAttribute = ZigBee.IAttribute;
        // noinspection ES6UnusedImports
        import ValueType = ZigBee.ValueType;
        
        export namespace Otau {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
            export type ICmdQueryNextImagePayload = { HardwareVersionPresent?: IArgHardwareVersionPresentPayload, ManufacturerCode?: IArgManufacturerCodePayload, ImageType?: IArgImageTypePayload, FileVersion?: IArgFileVersionPayload, HardwareVersion?: IArgHardwareVersionPayload, }
            export interface ICmdQueryNextImage extends ICommand { value: ICmdQueryNextImagePayload }
            export type ICmdImageBlockRequestPayload = { BlockRequestOptions?: IArgBlockRequestOptionsPayload, ManufacturerCode?: IArgManufacturerCodePayload, ImageType?: IArgImageTypePayload, FileVersion?: IArgFileVersionPayload, FileOffset?: IArgFileOffsetPayload, MaxDataSize?: IArgDataSizePayload, }
            export interface ICmdImageBlockRequest extends ICommand { value: ICmdImageBlockRequestPayload }
            export type ICmdUpgradeEndRequestPayload = { Status?: IArgStatusPayload, ManufacturerCode?: IArgManufacturerCodePayload, ImageType?: IArgImageTypePayload, FileVersion?: IArgFileVersionPayload, }
            export interface ICmdUpgradeEndRequest extends ICommand { value: ICmdUpgradeEndRequestPayload }
            export type ICmdImageNotifyPayload = { ImageNotifyPayloadType?: IArgImageNotifyPayloadTypePayload, QueryJitter?: IArgQueryJitterPayload, ManufacturerCode?: IArgManufacturerCodePayload, ImageType?: IArgImageTypePayload, FileVersion?: IArgFileVersionPayload, }
            export interface ICmdImageNotify extends ICommand { value: ICmdImageNotifyPayload }
            export type ICmdQueryNextImageResponsePayload = { Status?: IArgStatusPayload, ManufacturerCode?: IArgManufacturerCodePayload, ImageType?: IArgImageTypePayload, FileVersion?: IArgFileVersionPayload, ImageSize?: IArgImageSizePayload, }
            export interface ICmdQueryNextImageResponse extends ICommand { value: ICmdQueryNextImageResponsePayload }
            export type ICmdImageBlockResponsePayload = { Status?: IArgStatusPayload, ManufacturerCode?: IArgManufacturerCodePayload, ImageType?: IArgImageTypePayload, FileVersion?: IArgFileVersionPayload, FileOffset?: IArgFileOffsetPayload, Payload?: IArgPayloadPayload, CurrentTime?: IArgCurrentTimePayload, RequestTime?: IArgRequestTimePayload, }
            export interface ICmdImageBlockResponse extends ICommand { value: ICmdImageBlockResponsePayload }
            export type ICmdUpgradeEndResponsePayload = { ManufacturerCode?: IArgManufacturerCodePayload, ImageType?: IArgImageTypePayload, FileVersion?: IArgFileVersionPayload, CurrentTime?: IArgCurrentTimePayload, RequestTime?: IArgRequestTimePayload, }
            export interface ICmdUpgradeEndResponse extends ICommand { value: ICmdUpgradeEndResponsePayload }
        }

            export type IArgBlockRequestOptionsPayload = ValueType;
            export interface IArgBlockRequestOptions extends IArgument { value: IArgBlockRequestOptionsPayload }
            export type IArgCurrentFileVersionPayload = ValueType;
            export interface IArgCurrentFileVersion extends IAttribute { value: IArgCurrentFileVersionPayload }
            export type IArgCurrentStackVersionPayload = ValueType;
            export interface IArgCurrentStackVersion extends IAttribute { value: IArgCurrentStackVersionPayload }
            export type IArgCurrentTimePayload = ValueType;
            export interface IArgCurrentTime extends IArgument { value: IArgCurrentTimePayload }
            export type IArgDataSizePayload = ValueType;
            export interface IArgDataSize extends IArgument { value: IArgDataSizePayload }
            export type IArgDownloadedFileVersionPayload = ValueType;
            export interface IArgDownloadedFileVersion extends IAttribute { value: IArgDownloadedFileVersionPayload }
            export type IArgDownloadedStackVersionPayload = ValueType;
            export interface IArgDownloadedStackVersion extends IAttribute { value: IArgDownloadedStackVersionPayload }
            export type IArgFileOffsetPayload = ValueType;
            export interface IArgFileOffset extends IAttribute { value: IArgFileOffsetPayload }
            export type IArgFileVersionPayload = ValueType;
            export interface IArgFileVersion extends IArgument { value: IArgFileVersionPayload }
            export type IArgHardwareVersionPayload = ValueType;
            export interface IArgHardwareVersion extends IArgument { value: IArgHardwareVersionPayload }
            export type IArgHardwareVersionPresentPayload = ValueType;
            export interface IArgHardwareVersionPresent extends IArgument { value: IArgHardwareVersionPresentPayload }
            export type IArgImageNotifyPayloadTypePayload = ValueType;
            export interface IArgImageNotifyPayloadType extends IArgument { value: IArgImageNotifyPayloadTypePayload }
            export type IArgImageSizePayload = ValueType;
            export interface IArgImageSize extends IArgument { value: IArgImageSizePayload }
            export type IArgImageTypePayload = ValueType;
            export interface IArgImageType extends IArgument { value: IArgImageTypePayload }
            export type IArgImageUpgradeStatusPayload = ValueType;
            export interface IArgImageUpgradeStatus extends IAttribute { value: IArgImageUpgradeStatusPayload }
            export type IArgManufacturerCodePayload = ValueType;
            export interface IArgManufacturerCode extends IArgument { value: IArgManufacturerCodePayload }
            export type IArgMinBlockRequestDelayPayload = ValueType;
            export interface IArgMinBlockRequestDelay extends IAttribute { value: IArgMinBlockRequestDelayPayload }
            export type IArgNextImageStatusPayload = ValueType;
            export interface IArgNextImageStatus extends IArgument { value: IArgNextImageStatusPayload }
            export type IArgPayloadPayload = ValueType;
            export interface IArgPayload extends IArgument { value: IArgPayloadPayload }
            export type IArgQueryJitterPayload = ValueType;
            export interface IArgQueryJitter extends IArgument { value: IArgQueryJitterPayload }
            export type IArgRequestTimePayload = ValueType;
            export interface IArgRequestTime extends IArgument { value: IArgRequestTimePayload }
            export type IArgStatusPayload = ValueType;
            export interface IArgStatus extends IArgument { value: IArgStatusPayload }
            export type IArgUpgradeServerPayload = ValueType;
            export interface IArgUpgradeServer extends IAttribute { value: IArgUpgradeServerPayload }    }

    export namespace ISmartEnergy {
        // noinspection ES6UnusedImports
        import IArgument = ZigBee.IArgument;
        // noinspection ES6UnusedImports
        import IAttribute = ZigBee.IAttribute;
        // noinspection ES6UnusedImports
        import ValueType = ZigBee.ValueType;
        
        export namespace Metering {
            // noinspection ES6UnusedImports
            import ICommand = ZigBee.ICommand;
            // noinspection ES6UnusedImports
            import ValueType = ZigBee.ValueType;

        
        }

            export type IArgControlTemperaturePayload = ValueType;
            export interface IArgControlTemperature extends IAttribute { value: IArgControlTemperaturePayload }
            export type IArgCurrentBlockPayload = ValueType;
            export interface IArgCurrentBlock extends IAttribute { value: IArgCurrentBlockPayload }
            export type IArgCurrentBlockPeriodConsumptionDeliveredPayload = ValueType;
            export interface IArgCurrentBlockPeriodConsumptionDelivered extends IAttribute { value: IArgCurrentBlockPeriodConsumptionDeliveredPayload }
            export type IArgCurrentInletEnergyCarrierDemandPayload = ValueType;
            export interface IArgCurrentInletEnergyCarrierDemand extends IAttribute { value: IArgCurrentInletEnergyCarrierDemandPayload }
            export type IArgCurrentInletEnergyCarrierSummationPayload = ValueType;
            export interface IArgCurrentInletEnergyCarrierSummation extends IAttribute { value: IArgCurrentInletEnergyCarrierSummationPayload }
            export type IArgCurrentMaxDemandDeliveredPayload = ValueType;
            export interface IArgCurrentMaxDemandDelivered extends IAttribute { value: IArgCurrentMaxDemandDeliveredPayload }
            export type IArgCurrentMaxDemandDeliveredTimePayload = ValueType;
            export interface IArgCurrentMaxDemandDeliveredTime extends IAttribute { value: IArgCurrentMaxDemandDeliveredTimePayload }
            export type IArgCurrentMaxDemandReceivedPayload = ValueType;
            export interface IArgCurrentMaxDemandReceived extends IAttribute { value: IArgCurrentMaxDemandReceivedPayload }
            export type IArgCurrentMaxDemandReceivedTimePayload = ValueType;
            export interface IArgCurrentMaxDemandReceivedTime extends IAttribute { value: IArgCurrentMaxDemandReceivedTimePayload }
            export type IArgCurrentOutletEnergyCarrierDemandPayload = ValueType;
            export interface IArgCurrentOutletEnergyCarrierDemand extends IAttribute { value: IArgCurrentOutletEnergyCarrierDemandPayload }
            export type IArgCurrentOutletEnergyCarrierSummationPayload = ValueType;
            export interface IArgCurrentOutletEnergyCarrierSummation extends IAttribute { value: IArgCurrentOutletEnergyCarrierSummationPayload }
            export type IArgCurrentSummationDeliveredPayload = ValueType;
            export interface IArgCurrentSummationDelivered extends IAttribute { value: IArgCurrentSummationDeliveredPayload }
            export type IArgCurrentSummationReceivedPayload = ValueType;
            export interface IArgCurrentSummationReceived extends IAttribute { value: IArgCurrentSummationReceivedPayload }
            export type IArgDftSummationPayload = ValueType;
            export interface IArgDftSummation extends IAttribute { value: IArgDftSummationPayload }
            export type IArgDailyConsumptionTargetPayload = ValueType;
            export interface IArgDailyConsumptionTarget extends IAttribute { value: IArgDailyConsumptionTargetPayload }
            export type IArgDailyFreezeTimePayload = ValueType;
            export interface IArgDailyFreezeTime extends IAttribute { value: IArgDailyFreezeTimePayload }
            export type IArgDefaultUpdatePeriodPayload = ValueType;
            export interface IArgDefaultUpdatePeriod extends IAttribute { value: IArgDefaultUpdatePeriodPayload }
            export type IArgFastPollUpdatePeriodPayload = ValueType;
            export interface IArgFastPollUpdatePeriod extends IAttribute { value: IArgFastPollUpdatePeriodPayload }
            export type IArgFlowRestrictionPayload = ValueType;
            export interface IArgFlowRestriction extends IAttribute { value: IArgFlowRestrictionPayload }
            export type IArgInletTemperaturePayload = ValueType;
            export interface IArgInletTemperature extends IAttribute { value: IArgInletTemperaturePayload }
            export type IArgIntervalReadReportingPeriodPayload = ValueType;
            export interface IArgIntervalReadReportingPeriod extends IAttribute { value: IArgIntervalReadReportingPeriodPayload }
            export type IArgOutletTemperaturePayload = ValueType;
            export interface IArgOutletTemperature extends IAttribute { value: IArgOutletTemperaturePayload }
            export type IArgPowerFactorPayload = ValueType;
            export interface IArgPowerFactor extends IAttribute { value: IArgPowerFactorPayload }
            export type IArgPresetReadingTimePayload = ValueType;
            export interface IArgPresetReadingTime extends IAttribute { value: IArgPresetReadingTimePayload }
            export type IArgPreviousBlockPeriodConsumptionDeliveredPayload = ValueType;
            export interface IArgPreviousBlockPeriodConsumptionDelivered extends IAttribute { value: IArgPreviousBlockPeriodConsumptionDeliveredPayload }
            export type IArgProfileIntervalPeriodPayload = ValueType;
            export interface IArgProfileIntervalPeriod extends IAttribute { value: IArgProfileIntervalPeriodPayload }
            export type IArgReadingSnapShotTimePayload = ValueType;
            export interface IArgReadingSnapShotTime extends IAttribute { value: IArgReadingSnapShotTimePayload }
            export type IArgSupplyStatusPayload = ValueType;
            export interface IArgSupplyStatus extends IAttribute { value: IArgSupplyStatusPayload }
            export type IArgVolumePerReportPayload = ValueType;
            export interface IArgVolumePerReport extends IAttribute { value: IArgVolumePerReportPayload }    }



    export const units = {
        SquareMeters: { unit: "m²", format: (v) => `${v}m²` },
        SquareFeet: { unit: "ft²", format: (v) => `${v}ft²` },
        Milliamperes: { unit: "mA", format: (v) => `${v}mA` },
        Amperes: { unit: "A", format: (v) => `${v}A` },
        Ohms: { unit: "Ω", format: (v) => `${v} Ω` },
        Volts: { unit: "V", format: (v) => `${v}V` },
        KiloVolts: { unit: "kV", format: (v) => `${v}kV` },
        MegaVolts: { unit: "MV", format: (v) => `${v}MV` },
        VoltAmperes: { unit: "VA", format: (v) => `${v}VA` },
        KiloVoltAmperes: { unit: "kVA", format: (v) => `${v}kVA` },
        MegaVoltAmperes: { unit: "MVA", format: (v) => `${v}MVA` },
        VoltAmperesReactive: { unit: "var", format: (v) => `${v} var` },
        KiloVoltAmperesReactive: { unit: "kvar", format: (v) => `${v} kvar` },
        MegaVoltAmperesReactive: { unit: "Mvar", format: (v) => `${v} Mvar` },
        DegreesPhase: { unit: "°", format: (v) => `${v}°` },
        PowerFactor: { unit: "PFC", format: (v) => `${v} PFC` },
        Joules: { unit: "J", format: (v) => `${v}J` },
        Kilojoules: { unit: "kJ", format: (v) => `${v}kJ` },
        WattHours: { unit: "Wh", format: (v) => `${v}Wh` },
        KilowattHours: { unit: "kWh", format: (v) => `${v}kWh` },
        BTUs: { unit: "BTU", format: (v) => `${v} BTU` },
        Therms: { unit: "thm", format: (v) => `${v} thm` }, // == 100000 BTU
        TonHours: { unit: "Th", format: (v) => `${v} Th` },
        JoulesPerKilogramDryAir: { unit: "J/kg dry air", format: (v) => `${v} J/kg dry air` },
        BTUsPerPoundDryAir: { unit: "BTU/lbs dry air", format: (v) => `${v} BTU/lbs dry air` },
        CyclesPerHour: { unit: "cycles/hr", format: (v) => `${v} cycles/hr` },
        CyclesPerMinute: { unit: "cycles/min", format: (v) => `${v} cycles/min` },
        Hertz: { unit: "Hz", format: (v) => `${v} Hz` },
        GramsOfWaterPerKilogramDryAir: { unit: "g water/kg dry air", format: (v) => `${v}g water/kg dry air` },
        PercentRelativeHumidity: { unit: "% RH", format: (v) => `${v}% RH` },
        Millimeters: { unit: "mm", format: (v) => `${v}mm` },
        Meters: { unit: "m", format: (v) => `${v}m` },
        Inches: { unit: "in", format: (v) => `${v}in` },
        Feet: { unit: "ft", format: (v) => `${v}ft` },
        WattsPerSquareFoot: { unit: "W/ft²", format: (v) => `${v}W/ft²` },
        WattsPerSquareMeter: { unit: "W/m²", format: (v) => `${v}W/m²` },
        Lumens: { unit: "lm", format: (v) => `${v}lm` },
        Luxes: { unit: "lux", format: (v) => `${v}lux` },
        FootCandles: { unit: "lm/ft²", format: (v) => `${v}lm/ft²` },
        Kilograms: { unit: "kg", format: (v) => `${v}kg` },
        PoundsMass: { unit: "lbm", format: (v) => `${v}lbm` },
        Tons: { unit: "ton", format: (v) => `${v} ton` },
        KilogramsPerSecond: { unit: "kg/s", format: (v) => `${v}kg/s` },
        KilogramsPerMinute: { unit: "kg/min", format: (v) => `${v}kg/min` },
        KilogramsPerHour: { unit: "kg/h", format: (v) => `${v}kg/h` },
        PoundsMassPerMinute: { unit: "lbm/min", format: (v) => `${v}lbm/min` },
        PoundsMassPerHour: { unit: "lbm/h", format: (v) => `${v}lbm/h` },
        Watts: { unit: "W", format: (v) => `${v}W` },
        Kilowatts: { unit: "kW", format: (v) => `${v}kW` },
        Megawatts: { unit: "MW", format: (v) => `${v}MW` },
        BTUsPerHour: { unit: "BTU/h", format: (v) => `${v} BTU/h` },
        Horsepower: { unit: "hp", format: (v) => `${v}hp` },
        TonsRefrigeration: { unit: "TR", format: (v) => `${v} TR` },
        Pascals: { unit: "Pa", format: (v) => `${v}Pa` },
        Kilopascals: { unit: "kPa", format: (v) => `${v}kPa` },
        Bars: { unit: "bar", format: (v) => `${v} bar` },
        PoundsForcePerSquareInch: { unit: "lbf/in²", format: (v) => `${v}lbf/in²` },
        CentimetersOfWater: { unit: "cm h2o", format: (v) => `${v}cm h2o` },
        InchesOfWater: { unit: "in h2o", format: (v) => `${v}in h2o` },
        MillimetersOfMercury: { unit: "mmHg", format: (v) => `${v}mmHg` },
        CentimetersOfMercury: { unit: "cmHg", format: (v) => `${v}cmHg` },
        InchesOfMercury: { unit: "inHg", format: (v) => `${v}inHg` },
        DegreesCelsius: { unit: "°C", format: (v) => `${v}°C` },
        DegreesKelvin: { unit: "K", format: (v) => `${v}K` },
        DegreesFahrenheit: { unit: "°F", format: (v) => `${v}°F` },
        DegreeDaysCelsius: { unit: "°C DD", format: (v) => `${v}°C DD` },
        DegreeDaysFahrenheit: { unit: "°F DD", format: (v) => `${v}°F DD` },
        Years: { unit: "years", format: (v) => `${v} year${v!=1?'s':''}` },
        Months: { unit: "months", format: (v) => `${v} month${v!=1?'s':''}` },
        Weeks: { unit: "weeks", format: (v) => `${v} week${v!=1?'s':''}` },
        Days: { unit: "days", format: (v) => `${v} day${v!=1?'s':''}` },
        Hours: { unit: "hours", format: (v) => `${v} hour${v!=1?'s':''}` },
        Minutes: { unit: "minutes", format: (v) => `${v} minute${v!=1?'s':''}` },
        Seconds: { unit: "seconds", format: (v) => `${v} second${v!=1?'s':''}` },
        MetersPerSecond: { unit: "m/s", format: (v) => `${v}m/s` },
        KilometersPerHour: { unit: "km/h", format: (v) => `${v}km/h` },
        FeetPerSecond: { unit: "ft/s", format: (v) => `${v}ft/s` },
        FeetPerMinute: { unit: "ft/min", format: (v) => `${v}ft/min` },
        MilesPerHour: { unit: "mph", format: (v) => `${v}mph` },
        CubicFeet: { unit: "ft³", format: (v) => `${v}ft³` },
        CubicMeters: { unit: "m³", format: (v) => `${v}m³` },
        ImperialGallons: { unit: "gal (Imp.)", format: (v) => `${v}gal (Imp.)` },
        Liters: { unit: "L", format: (v) => `${v}L` },
        UsGallons: { unit: "gal (US)", format: (v) => `${v}gal (US)` },
        CubicFeetPerMinute: { unit: "ft³/min", format: (v) => `${v}ft³/min` },
        CubicMetersPerSecond: { unit: "m³/s", format: (v) => `${v}m³/s` },
        ImperialGallonsPerMinute: { unit: "gal/min (Imp.)", format: (v) => `${v}gal/min (Imp.)` },
        LitersPerSecond: { unit: "L/s", format: (v) => `${v}L/s` },
        LitersPerMinute: { unit: "L/min", format: (v) => `${v}L/min` },
        UsGallonsPerMinute: { unit: "gal/min (US)", format: (v) => `${v}gal/min (US)` },
        DegreesAngular: { unit: "°", format: (v) => `${v}°` },
        DegreesCelsiusPerHour: { unit: "°C/h", format: (v) => `${v}°C/h` },
        DegreesCelsiusPerMinute: { unit: "°C/min", format: (v) => `${v}°C/min` },
        DegreesFahrenheitPerHour: { unit: "°F/h", format: (v) => `${v}°F/h` },
        DegreesFahrenheitPerMinute: { unit: "°F/min", format: (v) => `${v}°F/min` },
        NoUnits: { unit: "", format: (v) => `${v}` },
        PartsPerMillion: { unit: "ppm", format: (v) => `${v}ppm` },
        PartsPerBillion: { unit: "ppb", format: (v) => `${v}ppb` },
        Percent: { unit: "%", format: (v) => `${Math.floor(v)}%` },
        PercentPerSecond: { unit: "%/s", format: (v) => `${Math.floor(v)}%/s` },
        PerMinute: { unit: "/min", format: (v) => `${v}/min` },
        PerSecond: { unit: "/s", format: (v) => `${v}/s` },
        PsiPerDegreeFahrenheit: { unit: "psi/°F", format: (v) => `${v}psi/°F` },
        Radians: { unit: "rad", format: (v) => `${v}rad` },
        RevolutionsPerMinute: { unit: "RPM", format: (v) => `${v} RPM` },
        Currency1: { unit: "Currency1", format: (v) => `${v}` },
        Currency2: { unit: "Currency2", format: (v) => `${v}` },
        Currency3: { unit: "Currency3", format: (v) => `${v}` },
        Currency4: { unit: "Currency4", format: (v) => `${v}` },
        Currency5: { unit: "Currency5", format: (v) => `${v}` },
        Currency6: { unit: "Currency6", format: (v) => `${v}` },
        Currency7: { unit: "Currency7", format: (v) => `${v}` },
        Currency8: { unit: "Currency8", format: (v) => `${v}` },
        Currency9: { unit: "Currency9", format: (v) => `${v}` },
        Currency10: { unit: "Currency10", format: (v) => `${v}` },
        SquareInches: { unit: "in²", format: (v) => `${v}in²` },
        SquareCentimeters: { unit: "cm²", format: (v) => `${v}cm²` },
        BTUsPerPound: { unit: "BTU/lbs", format: (v) => `${v} BTU/lbs` },
        Centimeters: { unit: "cm", format: (v) => `${v}cm` },
        PoundsMassPerSecond: { unit: "lbm/s", format: (v) => `${v}lbm/s` },
        DeltaDegreesFahrenheit: { unit: "delta °F", format: (v) => `${v} delta °F` },
        DeltaDegreesKelvin: { unit: "delta-K", format: (v) => `${v} delta-K` },
        Kilohms: { unit: "kΩ", format: (v) => `${v}kΩ` },
        Megohms: { unit: "MΩ", format: (v) => `${v}MΩ` },
        Millivolts: { unit: "mV", format: (v) => `${v}mV` },
        KilojoulesPerKilogram: { unit: "kJ/kg", format: (v) => `${v}kJ/kg` },
        Megajoules: { unit: "MJ", format: (v) => `${v}MJ` },
        JoulesPerDegreeKelvin: { unit: "J/°K", format: (v) => `${v}J/°K` },
        JoulesPerKilogramDegreeKelvin: { unit: "J/kg-K", format: (v) => `${v}J/kg-K` },
        Kilohertz: { unit: "kHz", format: (v) => `${v}kHz` },
        Megahertz: { unit: "MHz", format: (v) => `${v}MHz` },
        PerHour: { unit: "/h", format: (v) => `${v}/h` },
        Milliwatts: { unit: "mW", format: (v) => `${v}mW` },
        Hectopascals: { unit: "hpa", format: (v) => `${v}hpa` },
        Millibars: { unit: "mbar", format: (v) => `${v}mbar` },
        CubicMetersPerHour: { unit: "m³/h", format: (v) => `${v}m³/h` },
        LitersPerHour: { unit: "L/h", format: (v) => `${v}L/h` },
        KilowattHoursPerSquareMeter: { unit: "kWh/m²", format: (v) => `${v}kWh/m²` },
        KilowattHoursPerSquareFoot: { unit: "kWh/ft²", format: (v) => `${v}kWh/ft²` },
        MegajoulesPerSquareMeter: { unit: "MJ/m²", format: (v) => `${v}MJ/m²` },
        MegajoulesPerSquareFoot: { unit: "MJ/ft²", format: (v) => `${v}MJ/ft²` },
        WattsPerSquareMeterDegreeKelvin: { unit: "W/m²-K", format: (v) => `${v}W/m²-K` },
        CubicFeetPerSecond: { unit: "ft³/s", format: (v) => `${v}ft³/s` },
        PercentObscurationPerFoot: { unit: "% obscuration/ft", format: (v) => `${v}% obscuration/ft` },
        PercentObscurationPerMeter: { unit: "% obscuration/m", format: (v) => `${v}% obscuration/m` },
        Milliohms: { unit: "mΩ", format: (v) => `${v}mΩ` },
        MegawattHours: { unit: "MWh", format: (v) => `${v}MWh` },
        KiloBTUs: { unit: "kBTU", format: (v) => `${v} kBTU` },
        MegaBTUs: { unit: "MBTU", format: (v) => `${v} MBTU` },
        KilojoulesPerKilogramDryAir: { unit: "KJ/kg dry air", format: (v) => `${v} KJ/kg dry air` },
        MegajoulesPerKilogramDryAir: { unit: "MJ/kg dry air", format: (v) => `${v} MJ/kg dry air` },
        KilojoulesPerDegreeKelvin: { unit: "KJ/deg-K", format: (v) => `${v} KJ/deg-K` },
        MegajoulesPerDegreeKelvin: { unit: "MJ/deg-K", format: (v) => `${v} MJ/deg-K` },
        Newton: { unit: "N", format: (v) => `${v}N` },
        GramsPerSecond: { unit: "g/s", format: (v) => `${v}g/s` },
        GramsPerMinute: { unit: "g/min", format: (v) => `${v}g/min` },
        TonsPerHour: { unit: "tons/h", format: (v) => `${v} tons/h` },
        KiloBTUsPerHour: { unit: "kBTU/h", format: (v) => `${v} kBTU/h` },
        HundredthsSeconds: { unit: "Centiseconds", format: (v) => `${v} centisecond` },
        Milliseconds: { unit: "ms", format: (v) => `${v}ms` },
        NewtonMeters: { unit: "Nm", format: (v) => `${v}Nm` },
        MillimetersPerSecond: { unit: "mm/s", format: (v) => `${v}mm/s` },
        MillimetersPerMinute: { unit: "mm/min", format: (v) => `${v}mm/min` },
        MetersPerMinute: { unit: "m/min", format: (v) => `${v}m/min` },
        MetersPerHour: { unit: "m/h", format: (v) => `${v}m/h` },
        CubicMetersPerMinute: { unit: "m³/min", format: (v) => `${v}m³/min` },
        MetersPerSecondPerSecond: { unit: "m/s²", format: (v) => `${v}m/s²` },
        AmperesPerMeter: { unit: "A/m", format: (v) => `${v}A/m` },
        AmperesPerSquareMeter: { unit: "A/m²", format: (v) => `${v}A/m²` },
        AmpereSquareMeters: { unit: "Am²", format: (v) => `${v}Am²` },
        Farads: { unit: "Farads", format: (v) => `${v} Farads` },
        Henrys: { unit: "Henrys", format: (v) => `${v} Henrys` },
        OhmMeters: { unit: "Ωm", format: (v) => `${v}Ωm` },
        Siemens: { unit: "Siemens", format: (v) => `${v} Siemens` },
        SiemensPerMeter: { unit: "Siemens/m", format: (v) => `${v} Siemens/m` },
        Teslas: { unit: "Teslas", format: (v) => `${v} Teslas` },
        VoltsPerDegreeKelvin: { unit: "V/deg-K", format: (v) => `${v}V/deg-K` },
        VoltsPerMeter: { unit: "V/m", format: (v) => `${v}V/m` },
        Webers: { unit: "Webers", format: (v) => `${v} Webers` },
        Candelas: { unit: "Candelas", format: (v) => `${v} Candelas` },
        CandelasPerSquareMeter: { unit: "Candelas/m²", format: (v) => `${v} Candelas/m²` },
        KelvinsPerHour: { unit: "K/h", format: (v) => `${v}K/h` },
        KelvinsPerMinute: { unit: "K/min", format: (v) => `${v}K/min` },
        JouleSeconds: { unit: "Js", format: (v) => `${v}Js` },
        SquareMetersPerNewton: { unit: "m²/N", format: (v) => `${v}m²/N` },
        KilogramPerCubicMeter: { unit: "kg/m³", format: (v) => `${v}kg/m³` },
        NewtonSeconds: { unit: "Ns", format: (v) => `${v}Ns` },
        NewtonsPerMeter: { unit: "N/m", format: (v) => `${v}N/m` },
        WattsPerMeterPerDegreeKelvin: { unit: "W/m/deg-K", format: (v) => `${v}W/m/deg-K` },
        Other: { unit: "other", format: (v) => `${v}` },
        DecibelMilliWatts: { unit: "dBm", format: (v) => `${v}` },
        MilliAmpereHours: { unit: "mAh", format: (v) => `${v}mAh` },
        Mired: { unit: "mired", format: (v) => `${v} mired` },
        MicrogramPerCubicMeter: { unit: "µg/m³", format: (v) => `${v}µg/m³` },
        Bytes: { unit: "bytes",
            format: (v) => {
                let si = ['Ki', 'Mi', 'Gi', 'Ti', 'Pi'];
                let suf = '';
                while (v >= 2048 && si.length > 0) {
                    suf = si.shift();
                    v /= 1024;
                }
                return `${v}${suf}B`
            }},
        Concentration: { unit: "concentration",
            format: (v) => {
                if (v >= 0.001) return `${v*100}%`;
                if (v >= 0.000001) return `${v*1000000} ppm`;
                return `${v*1000000000} ppb`;
            }},
    };
    export const typeID = {
        8: base.dat8,
        9: base.dat16,
        10: base.dat24,
        11: base.dat32,
        12: base.dat40,
        13: base.dat48,
        14: base.dat56,
        15: base.dat64,
        16: base.bool,
        24: base.bmp8,
        25: base.bmp16,
        26: base.bmp24,
        27: base.bmp32,
        28: base.bmp40,
        29: base.bmp48,
        30: base.bmp56,
        31: base.bmp64,
        32: base.u8,
        33: base.u16,
        34: base.u24,
        35: base.u32,
        36: base.u40,
        37: base.u48,
        38: base.u56,
        39: base.u64,
        40: base.s8,
        41: base.s16,
        42: base.s24,
        43: base.s32,
        44: base.s40,
        45: base.s48,
        46: base.s56,
        47: base.s64,
        48: base.enum8,
        49: base.enum16,
        56: base.semi,
        57: base.float,
        58: base.double,
        65: base.ostring,
        66: base.cstring,
        67: base.lostring,
        68: base.lcstring,
        72: base.array,
        76: base.struct,
        80: base.set,
        81: base.bag,
        224: base.time,
        225: base.date,
        226: base.utc,
        232: base.cid,
        233: base.aid,
        234: base.oid,
        240: base.uid,
        241: base.seckey,
    };
    export const unitID = {
        0x00: units.SquareMeters,  //  Area
        0x01: units.SquareFeet,  //  Area
        0x02: units.Milliamperes,  //  Electrical
        0x03: units.Amperes,  //  Electrical
        0x04: units.Ohms,  //  Electrical
        0x05: units.Volts,  //  Electrical
        0x06: units.KiloVolts,  //  Electrical
        0x07: units.MegaVolts,  //  Electrical
        0x08: units.VoltAmperes,  //  Electrical
        0x09: units.KiloVoltAmperes,  //  Electrical
        0x0A: units.MegaVoltAmperes,  //  Electrical
        0x0B: units.VoltAmperesReactive,  //  Electrical
        0x0C: units.KiloVoltAmperesReactive,  //  Electrical
        0x0D: units.MegaVoltAmperesReactive,  //  Electrical
        0x0E: units.DegreesPhase,  //  Electrical
        0x0F: units.PowerFactor,  //  Electrical
        0x10: units.Joules,  //  Energy
        0x11: units.Kilojoules,  //  Energy
        0x12: units.WattHours,  //  Energy
        0x13: units.KilowattHours,  //  Energy
        0x14: units.BTUs,  //  Energy
        0x15: units.Therms,  //  Energy
        0x16: units.TonHours,  //  Energy
        0x17: units.JoulesPerKilogramDryAir,  //  Enthalpy
        0x18: units.BTUsPerPoundDryAir,  //  Enthalpy
        0x19: units.CyclesPerHour,  //  Frequency
        0x1A: units.CyclesPerMinute,  //  Frequency
        0x1B: units.Hertz,  //  Frequency
        0x1C: units.GramsOfWaterPerKilogramDryAir,  //  Humidity
        0x1D: units.PercentRelativeHumidity,  //  Humidity
        0x1E: units.Millimeters,  //  Length
        0x1F: units.Meters,  //  Length
        0x20: units.Inches,  //  Length
        0x21: units.Feet,  //  Length
        0x22: units.WattsPerSquareFoot,  //  Light
        0x23: units.WattsPerSquareMeter,  //  Light
        0x24: units.Lumens,  //  Light
        0x25: units.Luxes,  //  Light
        0x26: units.FootCandles,  //  Light
        0x27: units.Kilograms,  //  Mass
        0x28: units.PoundsMass,  //  Mass
        0x29: units.Tons,  //  Mass
        0x2A: units.KilogramsPerSecond,  //  Mass Flow
        0x2B: units.KilogramsPerMinute,  //  Mass Flow
        0x2C: units.KilogramsPerHour,  //  Mass Flow
        0x2D: units.PoundsMassPerMinute,  //  Mass Flow
        0x2E: units.PoundsMassPerHour,  //  Mass Flow
        0x2F: units.Watts,  //  Power
        0x30: units.Kilowatts,  //  Power
        0x31: units.Megawatts,  //  Power
        0x32: units.BTUsPerHour,  //  Power
        0x33: units.Horsepower,  //  Power
        0x34: units.TonsRefrigeration,  //  Power
        0x35: units.Pascals,  //  Pressure
        0x36: units.Kilopascals,  //  Pressure
        0x37: units.Bars,  //  Pressure
        0x38: units.PoundsForcePerSquareInch,  //  Pressure
        0x39: units.CentimetersOfWater,  //  Pressure
        0x3A: units.InchesOfWater,  //  Pressure
        0x3B: units.MillimetersOfMercury,  //  Pressure
        0x3C: units.CentimetersOfMercury,  //  Pressure
        0x3D: units.InchesOfMercury,  //  Pressure
        0x3E: units.DegreesCelsius,  //  Temperature
        0x3F: units.DegreesKelvin,  //  Temperature
        0x40: units.DegreesFahrenheit,  //  Temperature
        0x41: units.DegreeDaysCelsius,  //  Temperature
        0x42: units.DegreeDaysFahrenheit,  //  Temperature
        0x43: units.Years,  //  Time
        0x44: units.Months,  //  Time
        0x45: units.Weeks,  //  Time
        0x46: units.Days,  //  Time
        0x47: units.Hours,  //  Time
        0x48: units.Minutes,  //  Time
        0x49: units.Seconds,  //  Time
        0x4A: units.MetersPerSecond,  //  Velocity
        0x4B: units.KilometersPerHour,  //  Velocity
        0x4C: units.FeetPerSecond,  //  Velocity
        0x4D: units.FeetPerMinute,  //  Velocity
        0x4E: units.MilesPerHour,  //  Velocity
        0x4F: units.CubicFeet,  //  Volume
        0x50: units.CubicMeters,  //  Volume
        0x51: units.ImperialGallons,  //  Volume
        0x52: units.Liters,  //  Volume
        0x53: units.UsGallons,  //  Volume
        0x54: units.CubicFeetPerMinute,  //  Volumetric Flow
        0x55: units.CubicMetersPerSecond,  //  Volumetric Flow
        0x56: units.ImperialGallonsPerMinute,  //  Volumetric Flow
        0x57: units.LitersPerSecond,  //  Volumetric Flow
        0x58: units.LitersPerMinute,  //  Volumetric Flow
        0x59: units.UsGallonsPerMinute,  //  Volumetric Flow
        0x5A: units.DegreesAngular,  //  Other
        0x5B: units.DegreesCelsiusPerHour,  //  Other
        0x5C: units.DegreesCelsiusPerMinute,  //  Other
        0x5D: units.DegreesFahrenheitPerHour,  //  Other
        0x5E: units.DegreesFahrenheitPerMinute,  //  Other
        0x5F: units.NoUnits,  //  Other
        0x60: units.PartsPerMillion,  //  Other
        0x61: units.PartsPerBillion,  //  Other
        0x62: units.Percent,  //  Other
        0x63: units.PercentPerSecond,  //  Other
        0x64: units.PerMinute,  //  Other
        0x65: units.PerSecond,  //  Other
        0x66: units.PsiPerDegreeFahrenheit,  //  Other
        0x67: units.Radians,  //  Other
        0x68: units.RevolutionsPerMinute,  //  Other
        0x69: units.Currency1,  //  Currency
        0x6A: units.Currency2,  //  Currency
        0x6B: units.Currency3,  //  Currency
        0x6C: units.Currency4,  //  Currency
        0x6D: units.Currency5,  //  Currency
        0x6E: units.Currency6,  //  Currency
        0x6F: units.Currency7,  //  Currency
        0x70: units.Currency8,  //  Currency
        0x71: units.Currency9,  //  Currency
        0x72: units.Currency10,  //  Currency
        0x73: units.SquareInches,  //  Area
        0x74: units.SquareCentimeters,  //  Area
        0x75: units.BTUsPerPound,  //  Enthalpy
        0x76: units.Centimeters,  //  Length
        0x77: units.PoundsMassPerSecond,  //  Mass Flow
        0x78: units.DeltaDegreesFahrenheit,  //  Temperature
        0x79: units.DeltaDegreesKelvin,  //  Temperature
        0x7A: units.Kilohms,  //  Electrical
        0x7B: units.Megohms,  //  Electrical
        0x7C: units.Millivolts,  //  Electrical
        0x7D: units.KilojoulesPerKilogram,  //  Energy
        0x7E: units.Megajoules,  //  Energy
        0x7F: units.JoulesPerDegreeKelvin,  //  Entrophy
        0x80: units.JoulesPerKilogramDegreeKelvin,  //  Entrophy
        0x81: units.Kilohertz,  //  Frequency
        0x82: units.Megahertz,  //  Frequency
        0x83: units.PerHour,  //  Frequency
        0x84: units.Milliwatts,  //  Power
        0x85: units.Hectopascals,  //  Pressure
        0x86: units.Millibars,  //  Pressure
        0x87: units.CubicMetersPerHour,  //  Volumetric Flow
        0x88: units.LitersPerHour,  //  Volumetric Flow
        0x89: units.KilowattHoursPerSquareMeter,  //  Other
        0x8A: units.KilowattHoursPerSquareFoot,  //  Other
        0x8B: units.MegajoulesPerSquareMeter,  //  Other
        0x8C: units.MegajoulesPerSquareFoot,  //  Other
        0x8D: units.WattsPerSquareMeterDegreeKelvin,  //  Other
        0x8E: units.CubicFeetPerSecond,  //  Volumetric Flow
        0x8F: units.PercentObscurationPerFoot,  //  Other
        0x90: units.PercentObscurationPerMeter,  //  Other
        0x91: units.Milliohms,  //  Electrical
        0x92: units.MegawattHours,  //  Energy
        0x93: units.KiloBTUs,  //  Energy
        0x94: units.MegaBTUs,  //  Energy
        0x95: units.KilojoulesPerKilogramDryAir,  //  Enthalpy
        0x96: units.MegajoulesPerKilogramDryAir,  //  Enthalpy
        0x97: units.KilojoulesPerDegreeKelvin,  //  Entrophy
        0x98: units.MegajoulesPerDegreeKelvin,  //  Entrophy
        0x99: units.Newton,  //  Force
        0x9A: units.GramsPerSecond,  //  Mass Flow
        0x9B: units.GramsPerMinute,  //  Mass Flow
        0x9C: units.TonsPerHour,  //  Mass Flow
        0x9D: units.KiloBTUsPerHour,  //  Power
        0x9E: units.HundredthsSeconds,  //  Time
        0x9F: units.Milliseconds,  //  Time
        0xA0: units.NewtonMeters,  //  Torque
        0xA1: units.MillimetersPerSecond,  //  Velocity
        0xA2: units.MillimetersPerMinute,  //  Velocity
        0xA3: units.MetersPerMinute,  //  Velocity
        0xA4: units.MetersPerHour,  //  Velocity
        0xA5: units.CubicMetersPerMinute,  //  Volumetric Flow
        0xA6: units.MetersPerSecondPerSecond,  //  Acceleration
        0xA7: units.AmperesPerMeter,  //  Electrical
        0xA8: units.AmperesPerSquareMeter,  //  Electrical
        0xA9: units.AmpereSquareMeters,  //  Electrical
        0xAA: units.Farads,  //  Electrical
        0xAB: units.Henrys,  //  Electrical
        0xAC: units.OhmMeters,  //  Electrical
        0xAD: units.Siemens,  //  Electrical
        0xAE: units.SiemensPerMeter,  //  Electrical
        0xAF: units.Teslas,  //  Electrical
        0xB0: units.VoltsPerDegreeKelvin,  //  Electrical
        0xB1: units.VoltsPerMeter,  //  Electrical
        0xB2: units.Webers,  //  Electrical
        0xB3: units.Candelas,  //  Light
        0xB4: units.CandelasPerSquareMeter,  //  Light
        0xB5: units.KelvinsPerHour,  //  Temperature
        0xB6: units.KelvinsPerMinute,  //  Temperature
        0xB7: units.JouleSeconds,  //  Other
        0xB8: units.SquareMetersPerNewton,  //  Other
        0xB9: units.KilogramPerCubicMeter,  //  Other
        0xBA: units.NewtonSeconds,  //  Other
        0xBB: units.NewtonsPerMeter,  //  Other
        0xBC: units.WattsPerMeterPerDegreeKelvin,  //  Other
        0xFF: units.Other,
    };

}

