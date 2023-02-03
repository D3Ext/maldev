package system

import (
  "os/user"
)

func GetGroups() ([]string, error) {
  var groups_to_return []string // This variable handle group names (e.g. sys, root, admin...)

  current_user, err := user.Current() // Get atributes from current session
  if err != nil {
    return groups_to_return, err
  }

  general_groups, err := current_user.GroupIds() // Get list of group IDs
  if err != nil {
    return groups_to_return, err
  }

  for _, id := range general_groups { // Iterate over each ID
    group_info, err := user.LookupGroupId(id) // Find ID associated group
    if err != nil {
      return groups_to_return, err
    }
    groups_to_return = append(groups_to_return, group_info.Name)
  }
  
  return groups_to_return, nil
}
