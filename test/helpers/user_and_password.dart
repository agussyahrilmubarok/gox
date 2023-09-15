import 'package:blog_shelf/users/dtos/user_dto.dart';

class UserAndPassword {
  final UserDto user;
  final String password;

  UserAndPassword({required this.user, required this.password});
}
