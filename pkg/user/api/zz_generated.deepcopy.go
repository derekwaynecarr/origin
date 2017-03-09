// +build !ignore_autogenerated_openshift

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package api

import (
	reflect "reflect"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	pkg_api "k8s.io/kubernetes/pkg/api"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_Group, InType: reflect.TypeOf(&Group{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_GroupList, InType: reflect.TypeOf(&GroupList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_Identity, InType: reflect.TypeOf(&Identity{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_IdentityList, InType: reflect.TypeOf(&IdentityList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_User, InType: reflect.TypeOf(&User{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_UserIdentityMapping, InType: reflect.TypeOf(&UserIdentityMapping{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_UserList, InType: reflect.TypeOf(&UserList{})},
	)
}

func DeepCopy_api_Group(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Group)
		out := out.(*Group)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if in.Users != nil {
			in, out := &in.Users, &out.Users
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.Users = nil
		}
		return nil
	}
}

func DeepCopy_api_GroupList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GroupList)
		out := out.(*GroupList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Group, len(*in))
			for i := range *in {
				if err := DeepCopy_api_Group(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_api_Identity(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Identity)
		out := out.(*Identity)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		out.ProviderName = in.ProviderName
		out.ProviderUserName = in.ProviderUserName
		out.User = in.User
		if in.Extra != nil {
			in, out := &in.Extra, &out.Extra
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.Extra = nil
		}
		return nil
	}
}

func DeepCopy_api_IdentityList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IdentityList)
		out := out.(*IdentityList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Identity, len(*in))
			for i := range *in {
				if err := DeepCopy_api_Identity(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_api_User(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*User)
		out := out.(*User)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		out.FullName = in.FullName
		if in.Identities != nil {
			in, out := &in.Identities, &out.Identities
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.Identities = nil
		}
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.Groups = nil
		}
		return nil
	}
}

func DeepCopy_api_UserIdentityMapping(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UserIdentityMapping)
		out := out.(*UserIdentityMapping)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		out.Identity = in.Identity
		out.User = in.User
		return nil
	}
}

func DeepCopy_api_UserList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UserList)
		out := out.(*UserList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]User, len(*in))
			for i := range *in {
				if err := DeepCopy_api_User(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}
