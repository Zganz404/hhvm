// Autogenerated by Thrift for thrift/annotation/rust.thrift
//
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
//  @generated

package rust

import (
    "fmt"
    "strings"

    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction)
var _ = fmt.Printf
var _ = strings.Split
var _ = thrift.ZERO


type Name struct {
    Name string `thrift:"name,1" json:"name" db:"name"`
}
// Compile time interface enforcer
var _ thrift.Struct = (*Name)(nil)

func NewName() *Name {
    return (&Name{}).
        SetNameNonCompat("")
}

func (x *Name) GetName() string {
    return x.Name
}

func (x *Name) SetNameNonCompat(value string) *Name {
    x.Name = value
    return x
}

func (x *Name) SetName(value string) *Name {
    x.Name = value
    return x
}

func (x *Name) writeField1(p thrift.Format) error {  // Name
    if err := p.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Name
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Name) readField1(p thrift.Format) error {  // Name
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.Name = result
    return nil
}

func (x *Name) toString1() string {  // Name
    return fmt.Sprintf("%v", x.Name)
}



func (x *Name) Write(p thrift.Format) error {
    if err := p.WriteStructBegin("Name"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *Name) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        case (id == 1 && wireType == thrift.Type(thrift.STRING)):  // name
            if err := x.readField1(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *Name) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("Name({")
    sb.WriteString(fmt.Sprintf("Name:%s", x.toString1()))
    sb.WriteString("})")

    return sb.String()
}

type Copy struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*Copy)(nil)

func NewCopy() *Copy {
    return (&Copy{})
}



func (x *Copy) Write(p thrift.Format) error {
    if err := p.WriteStructBegin("Copy"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *Copy) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *Copy) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("Copy({")
    sb.WriteString("})")

    return sb.String()
}

type RequestContext struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*RequestContext)(nil)

func NewRequestContext() *RequestContext {
    return (&RequestContext{})
}



func (x *RequestContext) Write(p thrift.Format) error {
    if err := p.WriteStructBegin("RequestContext"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *RequestContext) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *RequestContext) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("RequestContext({")
    sb.WriteString("})")

    return sb.String()
}

type Arc struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*Arc)(nil)

func NewArc() *Arc {
    return (&Arc{})
}



func (x *Arc) Write(p thrift.Format) error {
    if err := p.WriteStructBegin("Arc"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *Arc) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *Arc) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("Arc({")
    sb.WriteString("})")

    return sb.String()
}

type Box struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*Box)(nil)

func NewBox() *Box {
    return (&Box{})
}



func (x *Box) Write(p thrift.Format) error {
    if err := p.WriteStructBegin("Box"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *Box) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *Box) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("Box({")
    sb.WriteString("})")

    return sb.String()
}

type Exhaustive struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*Exhaustive)(nil)

func NewExhaustive() *Exhaustive {
    return (&Exhaustive{})
}



func (x *Exhaustive) Write(p thrift.Format) error {
    if err := p.WriteStructBegin("Exhaustive"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *Exhaustive) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *Exhaustive) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("Exhaustive({")
    sb.WriteString("})")

    return sb.String()
}

type Ord struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*Ord)(nil)

func NewOrd() *Ord {
    return (&Ord{})
}



func (x *Ord) Write(p thrift.Format) error {
    if err := p.WriteStructBegin("Ord"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *Ord) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *Ord) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("Ord({")
    sb.WriteString("})")

    return sb.String()
}

type NewType_ struct {
}
// Compile time interface enforcer
var _ thrift.Struct = (*NewType_)(nil)

func NewNewType_() *NewType_ {
    return (&NewType_{})
}



func (x *NewType_) Write(p thrift.Format) error {
    if err := p.WriteStructBegin("NewType"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *NewType_) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *NewType_) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("NewType_({")
    sb.WriteString("})")

    return sb.String()
}

type Type struct {
    Name string `thrift:"name,1" json:"name" db:"name"`
}
// Compile time interface enforcer
var _ thrift.Struct = (*Type)(nil)

func NewType() *Type {
    return (&Type{}).
        SetNameNonCompat("")
}

func (x *Type) GetName() string {
    return x.Name
}

func (x *Type) SetNameNonCompat(value string) *Type {
    x.Name = value
    return x
}

func (x *Type) SetName(value string) *Type {
    x.Name = value
    return x
}

func (x *Type) writeField1(p thrift.Format) error {  // Name
    if err := p.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Name
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Type) readField1(p thrift.Format) error {  // Name
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.Name = result
    return nil
}

func (x *Type) toString1() string {  // Name
    return fmt.Sprintf("%v", x.Name)
}



func (x *Type) Write(p thrift.Format) error {
    if err := p.WriteStructBegin("Type"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *Type) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        case (id == 1 && wireType == thrift.Type(thrift.STRING)):  // name
            if err := x.readField1(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *Type) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("Type({")
    sb.WriteString(fmt.Sprintf("Name:%s", x.toString1()))
    sb.WriteString("})")

    return sb.String()
}

type Adapter struct {
    Name string `thrift:"name,1" json:"name" db:"name"`
}
// Compile time interface enforcer
var _ thrift.Struct = (*Adapter)(nil)

func NewAdapter() *Adapter {
    return (&Adapter{}).
        SetNameNonCompat("")
}

func (x *Adapter) GetName() string {
    return x.Name
}

func (x *Adapter) SetNameNonCompat(value string) *Adapter {
    x.Name = value
    return x
}

func (x *Adapter) SetName(value string) *Adapter {
    x.Name = value
    return x
}

func (x *Adapter) writeField1(p thrift.Format) error {  // Name
    if err := p.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Name
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Adapter) readField1(p thrift.Format) error {  // Name
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.Name = result
    return nil
}

func (x *Adapter) toString1() string {  // Name
    return fmt.Sprintf("%v", x.Name)
}



func (x *Adapter) Write(p thrift.Format) error {
    if err := p.WriteStructBegin("Adapter"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *Adapter) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        case (id == 1 && wireType == thrift.Type(thrift.STRING)):  // name
            if err := x.readField1(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *Adapter) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("Adapter({")
    sb.WriteString(fmt.Sprintf("Name:%s", x.toString1()))
    sb.WriteString("})")

    return sb.String()
}

type Derive struct {
    Derives []string `thrift:"derives,1" json:"derives" db:"derives"`
}
// Compile time interface enforcer
var _ thrift.Struct = (*Derive)(nil)

func NewDerive() *Derive {
    return (&Derive{}).
        SetDerivesNonCompat(make([]string, 0))
}

func (x *Derive) GetDerives() []string {
    if !x.IsSetDerives() {
        return make([]string, 0)
    }

    return x.Derives
}

func (x *Derive) SetDerivesNonCompat(value []string) *Derive {
    x.Derives = value
    return x
}

func (x *Derive) SetDerives(value []string) *Derive {
    x.Derives = value
    return x
}

func (x *Derive) IsSetDerives() bool {
    return x != nil && x.Derives != nil
}

func (x *Derive) writeField1(p thrift.Format) error {  // Derives
    if err := p.WriteFieldBegin("derives", thrift.LIST, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Derives
    if err := p.WriteListBegin(thrift.STRING, len(item)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
}
for _, v := range item {
    {
        item := v
        if err := p.WriteString(item); err != nil {
    return err
}
    }
}
if err := p.WriteListEnd(); err != nil {
    return thrift.PrependError("error writing list end: ", err)
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Derive) readField1(p thrift.Format) error {  // Derives
    _ /* elemType */, size, err := p.ReadListBegin()
if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
}

listResult := make([]string, 0, size)
for i := 0; i < size; i++ {
    var elem string
    {
        result, err := p.ReadString()
if err != nil {
    return err
}
        elem = result
    }
    listResult = append(listResult, elem)
}

if err := p.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
}
result := listResult

    x.Derives = result
    return nil
}

func (x *Derive) toString1() string {  // Derives
    return fmt.Sprintf("%v", x.Derives)
}



func (x *Derive) Write(p thrift.Format) error {
    if err := p.WriteStructBegin("Derive"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *Derive) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        case (id == 1 && wireType == thrift.Type(thrift.LIST)):  // derives
            if err := x.readField1(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *Derive) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("Derive({")
    sb.WriteString(fmt.Sprintf("Derives:%s", x.toString1()))
    sb.WriteString("})")

    return sb.String()
}

type ServiceExn struct {
    AnyhowToApplicationExn bool `thrift:"anyhow_to_application_exn,1" json:"anyhow_to_application_exn" db:"anyhow_to_application_exn"`
}
// Compile time interface enforcer
var _ thrift.Struct = (*ServiceExn)(nil)

func NewServiceExn() *ServiceExn {
    return (&ServiceExn{}).
        SetAnyhowToApplicationExnNonCompat(false)
}

func (x *ServiceExn) GetAnyhowToApplicationExn() bool {
    return x.AnyhowToApplicationExn
}

func (x *ServiceExn) SetAnyhowToApplicationExnNonCompat(value bool) *ServiceExn {
    x.AnyhowToApplicationExn = value
    return x
}

func (x *ServiceExn) SetAnyhowToApplicationExn(value bool) *ServiceExn {
    x.AnyhowToApplicationExn = value
    return x
}

func (x *ServiceExn) writeField1(p thrift.Format) error {  // AnyhowToApplicationExn
    if err := p.WriteFieldBegin("anyhow_to_application_exn", thrift.BOOL, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.AnyhowToApplicationExn
    if err := p.WriteBool(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *ServiceExn) readField1(p thrift.Format) error {  // AnyhowToApplicationExn
    result, err := p.ReadBool()
if err != nil {
    return err
}

    x.AnyhowToApplicationExn = result
    return nil
}

func (x *ServiceExn) toString1() string {  // AnyhowToApplicationExn
    return fmt.Sprintf("%v", x.AnyhowToApplicationExn)
}



func (x *ServiceExn) Write(p thrift.Format) error {
    if err := p.WriteStructBegin("ServiceExn"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *ServiceExn) Read(p thrift.Format) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        case (id == 1 && wireType == thrift.Type(thrift.BOOL)):  // anyhow_to_application_exn
            if err := x.readField1(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *ServiceExn) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("ServiceExn({")
    sb.WriteString(fmt.Sprintf("AnyhowToApplicationExn:%s", x.toString1()))
    sb.WriteString("})")

    return sb.String()
}

// RegisterTypes registers types found in this file that have a thrift_uri with the passed in registry.
func RegisterTypes(registry interface {
  RegisterType(name string, initializer func() any)
}) {
    registry.RegisterType("facebook.com/thrift/annotation/rust/Name", func() any { return NewName() })
    registry.RegisterType("facebook.com/thrift/annotation/rust/Copy", func() any { return NewCopy() })
    registry.RegisterType("facebook.com/thrift/annotation/rust/RequestContext", func() any { return NewRequestContext() })
    registry.RegisterType("facebook.com/thrift/annotation/rust/Arc", func() any { return NewArc() })
    registry.RegisterType("facebook.com/thrift/annotation/rust/Box", func() any { return NewBox() })
    registry.RegisterType("facebook.com/thrift/annotation/rust/Exhaustive", func() any { return NewExhaustive() })
    registry.RegisterType("facebook.com/thrift/annotation/rust/Ord", func() any { return NewOrd() })
    registry.RegisterType("facebook.com/thrift/annotation/rust/NewType", func() any { return NewNewType_() })
    registry.RegisterType("facebook.com/thrift/annotation/rust/Type", func() any { return NewType() })
    registry.RegisterType("facebook.com/thrift/annotation/rust/Adapter", func() any { return NewAdapter() })
    registry.RegisterType("facebook.com/thrift/annotation/rust/Derive", func() any { return NewDerive() })
    registry.RegisterType("facebook.com/thrift/annotation/rust/ServiceExn", func() any { return NewServiceExn() })

}
