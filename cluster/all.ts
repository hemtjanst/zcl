
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

                toString() {
                    let val = o.value;
                    let nv;
                    switch (typeof val) {
                        case 'undefined': return null;
                        case 'string': nv = parseFloat(val); break;
                        case 'number': nv = val; break;
                        case 'object':
                            if (Array.isArray(val) && val.length === 1 && typeof val[0] === 'number') {
                                nv = val[0];
                            }
                            break;
                    }
                    if (nv && typeof o.values === 'object' && typeof o.values[nv] !== 'undefined') {
                        return o.values[nv];
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

    const unitFmt = (a: IVal) => {

    };

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
                let yr = ((v & 0xFF000000) >> 24) + 1900,
                    mo = ((v & 0x00FF0000) >> 16) - 1,
                    day = (v & 0x0000FF00) >> 8,
                    wd = (v & 0x000000FF);
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

    export namespace Type {
        class Array {
            constructor(private name: string, private size) {
            }
        }
        class Struct {
            constructor(private name: string, private size) {
            }
        }
        class Bitmap {
            constructor(public name: string, public size) {
            }
        }
        class Byte {}
        class Uint {}
        class Int {}
        class Float {}
        class String {}
    }

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
                0x100B: `Philips`, 
                0x1021: `Legrand`, 
                0x1037: `NXP`, 
                0x10F2: `Ubisys`, 
                0x115F: `Xiaomi`, 
                0x1166: `innr`, 
                0x117C: `Ikea`,  },
                
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
                report: false,
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
                report: false,
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
        0x00: ZigBee.SecurityAndSafety.IasZone.ZoneStateChangeNotification,
        0x01: ZigBee.SecurityAndSafety.IasZone.ZoneEnrollRequest,
    };
    ZigBee.SecurityAndSafety.IasZone.Client.Command = { 
        0x00: ZigBee.SecurityAndSafety.IasZone.ZoneEnrollResponse,
        0x01: ZigBee.SecurityAndSafety.IasZone.InitiateNormalOperationMode,
        0x02: ZigBee.SecurityAndSafety.IasZone.InitiateTestMode,
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

    export const cluster: {[id: number]: ICluster} = { 
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
        0x0500: ZigBee.SecurityAndSafety.IasZone,
        0x0300: ZigBee.Lighting.ColorControl,
        0x0301: ZigBee.Lighting.BallastConfiguration,
        0x0400: ZigBee.MeasurementAndSensing.IlluminanceMeasurement,
        0x0401: ZigBee.MeasurementAndSensing.IlluminanceLevelSensing,
        0x0402: ZigBee.MeasurementAndSensing.TemperatureMeasurement,
        0x0403: ZigBee.MeasurementAndSensing.PressureMeasurement,
        0x0404: ZigBee.MeasurementAndSensing.FlowMeasurement,
        0x0405: ZigBee.MeasurementAndSensing.RelativeHumidityMeasurement,
        0x0406: ZigBee.MeasurementAndSensing.OccupancySensing,
        0x0019: ZigBee.Otau.Otau,
    };

    export namespace IZDP {
        import ICommand = ZigBee.ICommand;
        import IArgument = ZigBee.IArgument;
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
        import IArgument = ZigBee.IArgument;
        import IAttribute = ZigBee.IAttribute;
        import ValueType = ZigBee.ValueType;
        
        export namespace WindowCovering {
            import ICommand = ZigBee.ICommand;
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
        import IArgument = ZigBee.IArgument;
        import IAttribute = ZigBee.IAttribute;
        import ValueType = ZigBee.ValueType;
        
        export namespace Basic {
            import ICommand = ZigBee.ICommand;
            import ValueType = ZigBee.ValueType;

        
            export type ICmdResetToFactoryDefaultsPayload = { }
            export interface ICmdResetToFactoryDefaults extends ICommand { value: ICmdResetToFactoryDefaultsPayload }
        }

        export namespace PowerConfiguration {
            import ICommand = ZigBee.ICommand;
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace DeviceTemperatureConfiguration {
            import ICommand = ZigBee.ICommand;
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace Identify {
            import ICommand = ZigBee.ICommand;
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
            import ICommand = ZigBee.ICommand;
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
            import ICommand = ZigBee.ICommand;
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
            import ICommand = ZigBee.ICommand;
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
            import ICommand = ZigBee.ICommand;
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace LevelControl {
            import ICommand = ZigBee.ICommand;
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
            import ICommand = ZigBee.ICommand;
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
            import ICommand = ZigBee.ICommand;
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace Location {
            import ICommand = ZigBee.ICommand;
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
            import ICommand = ZigBee.ICommand;
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace AnalogOutput {
            import ICommand = ZigBee.ICommand;
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace AnalogValue {
            import ICommand = ZigBee.ICommand;
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace BinaryInput {
            import ICommand = ZigBee.ICommand;
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace BinaryOutput {
            import ICommand = ZigBee.ICommand;
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace BinaryValue {
            import ICommand = ZigBee.ICommand;
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace MultistateInput {
            import ICommand = ZigBee.ICommand;
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace MultistateOutput {
            import ICommand = ZigBee.ICommand;
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace MultistateValue {
            import ICommand = ZigBee.ICommand;
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace PollControl {
            import ICommand = ZigBee.ICommand;
            import ValueType = ZigBee.ValueType;

        
            export type ICmdCheckInPayload = { }
            export interface ICmdCheckIn extends ICommand { value: ICmdCheckInPayload }
        }

        export namespace MeterIdentification {
            import ICommand = ZigBee.ICommand;
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace Diagnostics {
            import ICommand = ZigBee.ICommand;
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

    export namespace ISecurityAndSafety {
        import IArgument = ZigBee.IArgument;
        import IAttribute = ZigBee.IAttribute;
        import ValueType = ZigBee.ValueType;
        
        export namespace IasZone {
            import ICommand = ZigBee.ICommand;
            import ValueType = ZigBee.ValueType;

        
            export type ICmdZoneStateChangeNotificationPayload = { ZoneStatus?: IArgZoneStatusPayload, ExtendedStatus?: IArgExtendedStatusPayload, ZoneId?: IArgZoneIdPayload, Delay?: IArgDelayPayload, }
            export interface ICmdZoneStateChangeNotification extends ICommand { value: ICmdZoneStateChangeNotificationPayload }
            export type ICmdZoneEnrollRequestPayload = { ZoneType?: IArgZoneTypePayload, ManufacturerCode?: IArgManufacturerCodePayload, }
            export interface ICmdZoneEnrollRequest extends ICommand { value: ICmdZoneEnrollRequestPayload }
            export type ICmdZoneEnrollResponsePayload = { EnrollResponseCode?: IArgEnrollResponseCodePayload, ZoneId?: IArgZoneIdPayload, }
            export interface ICmdZoneEnrollResponse extends ICommand { value: ICmdZoneEnrollResponsePayload }
            export type ICmdInitiateNormalOperationModePayload = { }
            export interface ICmdInitiateNormalOperationMode extends ICommand { value: ICmdInitiateNormalOperationModePayload }
            export type ICmdInitiateTestModePayload = { TestModeDuration?: IArgTestModeDurationPayload, CurrentZoneSensitivityLevel?: IArgCurrentZoneSensitivityLevelPayload, }
            export interface ICmdInitiateTestMode extends ICommand { value: ICmdInitiateTestModePayload }
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

    export namespace ILighting {
        import IArgument = ZigBee.IArgument;
        import IAttribute = ZigBee.IAttribute;
        import ValueType = ZigBee.ValueType;
        
        export namespace ColorControl {
            import ICommand = ZigBee.ICommand;
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
            import ICommand = ZigBee.ICommand;
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
        import IArgument = ZigBee.IArgument;
        import IAttribute = ZigBee.IAttribute;
        import ValueType = ZigBee.ValueType;
        
        export namespace IlluminanceMeasurement {
            import ICommand = ZigBee.ICommand;
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace IlluminanceLevelSensing {
            import ICommand = ZigBee.ICommand;
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace TemperatureMeasurement {
            import ICommand = ZigBee.ICommand;
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace PressureMeasurement {
            import ICommand = ZigBee.ICommand;
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace FlowMeasurement {
            import ICommand = ZigBee.ICommand;
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace RelativeHumidityMeasurement {
            import ICommand = ZigBee.ICommand;
            import ValueType = ZigBee.ValueType;

        
        }

        export namespace OccupancySensing {
            import ICommand = ZigBee.ICommand;
            import ValueType = ZigBee.ValueType;

        
        }

            export type IArgIlluminanceLightSensorTypePayload = ValueType;
            export interface IArgIlluminanceLightSensorType extends IAttribute { value: IArgIlluminanceLightSensorTypePayload }
            export type IArgIlluminanceTargetLevelPayload = ValueType;
            export interface IArgIlluminanceTargetLevel extends IAttribute { value: IArgIlluminanceTargetLevelPayload }
            export type IArgLevelStatusPayload = ValueType;
            export interface IArgLevelStatus extends IAttribute { value: IArgLevelStatusPayload }
            export type IArgLightSensorTypePayload = ValueType;
            export interface IArgLightSensorType extends IAttribute { value: IArgLightSensorTypePayload }
            export type IArgMaxMeasuredFlowPayload = ValueType;
            export interface IArgMaxMeasuredFlow extends IAttribute { value: IArgMaxMeasuredFlowPayload }
            export type IArgMaxMeasuredIlluminancePayload = ValueType;
            export interface IArgMaxMeasuredIlluminance extends IAttribute { value: IArgMaxMeasuredIlluminancePayload }
            export type IArgMaxMeasuredPressurePayload = ValueType;
            export interface IArgMaxMeasuredPressure extends IAttribute { value: IArgMaxMeasuredPressurePayload }
            export type IArgMaxMeasuredRelativeHumidityPayload = ValueType;
            export interface IArgMaxMeasuredRelativeHumidity extends IAttribute { value: IArgMaxMeasuredRelativeHumidityPayload }
            export type IArgMaxMeasuredTemperaturePayload = ValueType;
            export interface IArgMaxMeasuredTemperature extends IAttribute { value: IArgMaxMeasuredTemperaturePayload }
            export type IArgMeasuredFlowPayload = ValueType;
            export interface IArgMeasuredFlow extends IAttribute { value: IArgMeasuredFlowPayload }
            export type IArgMeasuredIlluminancePayload = ValueType;
            export interface IArgMeasuredIlluminance extends IAttribute { value: IArgMeasuredIlluminancePayload }
            export type IArgMeasuredPressurePayload = ValueType;
            export interface IArgMeasuredPressure extends IAttribute { value: IArgMeasuredPressurePayload }
            export type IArgMeasuredRelativeHumidityPayload = ValueType;
            export interface IArgMeasuredRelativeHumidity extends IAttribute { value: IArgMeasuredRelativeHumidityPayload }
            export type IArgMeasuredTemperaturePayload = ValueType;
            export interface IArgMeasuredTemperature extends IAttribute { value: IArgMeasuredTemperaturePayload }
            export type IArgMinMeasuredFlowPayload = ValueType;
            export interface IArgMinMeasuredFlow extends IAttribute { value: IArgMinMeasuredFlowPayload }
            export type IArgMinMeasuredIlluminancePayload = ValueType;
            export interface IArgMinMeasuredIlluminance extends IAttribute { value: IArgMinMeasuredIlluminancePayload }
            export type IArgMinMeasuredPressurePayload = ValueType;
            export interface IArgMinMeasuredPressure extends IAttribute { value: IArgMinMeasuredPressurePayload }
            export type IArgMinMeasuredRelativeHumidityPayload = ValueType;
            export interface IArgMinMeasuredRelativeHumidity extends IAttribute { value: IArgMinMeasuredRelativeHumidityPayload }
            export type IArgMinMeasuredTemperaturePayload = ValueType;
            export interface IArgMinMeasuredTemperature extends IAttribute { value: IArgMinMeasuredTemperaturePayload }
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
            export type IArgTolerancePayload = ValueType;
            export interface IArgTolerance extends IAttribute { value: IArgTolerancePayload }
            export type IArgUltrasonicOccupiedToUnoccupiedDelayPayload = ValueType;
            export interface IArgUltrasonicOccupiedToUnoccupiedDelay extends IAttribute { value: IArgUltrasonicOccupiedToUnoccupiedDelayPayload }
            export type IArgUltrasonicUnoccupiedToOccupiedDelayPayload = ValueType;
            export interface IArgUltrasonicUnoccupiedToOccupiedDelay extends IAttribute { value: IArgUltrasonicUnoccupiedToOccupiedDelayPayload }
            export type IArgUltrasonicUnoccupiedToOccupiedThresholdPayload = ValueType;
            export interface IArgUltrasonicUnoccupiedToOccupiedThreshold extends IAttribute { value: IArgUltrasonicUnoccupiedToOccupiedThresholdPayload }    }

    export namespace IOtau {
        import IArgument = ZigBee.IArgument;
        import IAttribute = ZigBee.IAttribute;
        import ValueType = ZigBee.ValueType;
        
        export namespace Otau {
            import ICommand = ZigBee.ICommand;
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

