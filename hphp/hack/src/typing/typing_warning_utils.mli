(*
 * Copyright (c) Meta Platforms, Inc. and affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the "hack" directory of this source tree.
 *
 *)

val code_is_enabled : GlobalOptions.t -> Error_codes.Warning.t -> bool

val add : Typing_env_types.env -> 'a Typing_warning.t -> unit

val add_for_migration :
  GlobalOptions.t ->
  as_lint:Tast.check_status option option ->
  Typing_warning.migrated Typing_warning.t ->
  unit
