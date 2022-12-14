/*
 * eve - management toolkit for libvirt servers
 * Copyright (C) 2022-2023  BNS Services LLC

 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.

 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.

 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package sessions

import (
	"context"

	"github.com/BasedDevelopment/eve/internal/db"
	"github.com/BasedDevelopment/eve/internal/tokens"
)

// Delete removes a session from the database (logout)
func Delete(ctx context.Context, token tokens.Token) error {
	_, err := db.Pool.Exec(ctx, "DELETE FROM sessions WHERE token_public = $1", token.Public)

	return err
}
